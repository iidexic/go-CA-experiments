package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/iidexic/go-CA-experiments/input"
)

type msgGen func() string
type msgScreenGen func(int, int) string
type showDebugInfo struct { // bools toggle what gets put into debug msg. len(Output)
	showDebug, tps, tick, screen, layouts, windowPX, frames, fps, keys, keysQty, khandlr bool
	len, pixW, pixH, gameW, gameH                                                        int
	Output                                                                               string
	keysAppend                                                                           []ebiten.Key
	keysDown                                                                             []ebiten.Key
	SelectDebug                                                                          []int
}

const (
	showTPS int = iota //0
	showTick
	showFrames
	showScreen
	showLayouts
	showWindowPX
	showFPS
	showKhandlr
)

// For now this is working, in the future maybe switch to iota+switch or select loop
var (
	frame, tick, layoutCount         int
	gWidth, gHeight, pWidth, pHeight int
	//Dbg is used to toggle desired debug output, receive screen info, and houses the final output
	Dbg showDebugInfo = showDebugInfo{
		showDebug:   true,
		tps:         true,
		tick:        false,
		frames:      false,
		screen:      false,
		layouts:     false,
		windowPX:    false,
		fps:         true,
		khandlr:     true,
		len:         0,
		gameW:       0,
		gameH:       0,
		pixW:        0,
		pixH:        0,
		keysAppend:  make([]ebiten.Key, 0, 12),
		Output:      "",
		SelectDebug: []int{showTPS, showFPS, showKhandlr},
	}
)

// SetValues currently sets screen values for debug display
func (d *showDebugInfo) SetValues(gameW, gameH, pixW, pixH int) {
	d.gameH = gameH
	d.gameW = gameW
	d.pixH = pixH
	d.pixW = pixW
}

// DebugBuildOutput is the NEW DEBUG MESSAGE GENERATOR AND PRINTER
func (d *showDebugInfo) DebugBuildOutput() {

	//^ Debug Writer/Function Store
	var sb strings.Builder
	var e error
	Dbg.Output = ""
	//outSlice := make([]string, len(d.SelectDebug))
	for _, v := range Dbg.SelectDebug {
		switch v { // can actually do full string assembly in here by using a strings.Builder...
		case showTPS:
			_, e = sb.WriteString(fmt.Sprintf("| tps: %f ", ebiten.ActualTPS()))
		case showTick:
			_, e = sb.WriteString(fmt.Sprintf("| tick: %03d ", tick/10))
		case showFrames:
			_, e = sb.WriteString(fmt.Sprintf("| frames: %03d ", frame/10))
		case showScreen:
			_, e = sb.WriteString(fmt.Sprintf("| game/screen: %dx%d ", d.gameW, d.gameH))
		case showLayouts:
			_, e = sb.WriteString(fmt.Sprintf("| layout: %d ", layoutCount/10))
		case showWindowPX:
			_, e = sb.WriteString(fmt.Sprintf("| px: %dx%d ", d.pixW, d.pixH))
		case showFPS:
			_, e = sb.WriteString(fmt.Sprintf("| fps: %f ", ebiten.ActualFPS()))
		case showKhandlr:
			kstr := ""
			keys := input.KeysOut()
			for _, k := range *keys {
				kstr += k.String()
			}
			_, e = sb.WriteString(fmt.Sprintf("\n| inKB[len %d]: %s", len(*keys), kstr))
		}
		if e != nil {
			log.Default()
		}

	}
	d.Output = sb.String()

	if !d.showDebug {
		d.Output = "[!!debug should be off!!]\n\n"
	}

}

// DebugMsgControl Builds the debug message
func DebugMsgControl(gameW, gameH, pixW, pixH int) {
	Dbg.gameW = gameW
	Dbg.gameH = gameH
	Dbg.pixW = pixW
	Dbg.pixH = pixH
	if !Dbg.showDebug {
		Dbg.Output = "debug should be off"
	}

	if ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			Dbg.tps = !Dbg.tps
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			Dbg.tick = !Dbg.tick
		} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
			Dbg.screen = !Dbg.screen
		} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
			Dbg.layouts = !Dbg.layouts
		} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
			Dbg.windowPX = !Dbg.windowPX
		} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
			Dbg.frames = !Dbg.frames
		} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
			Dbg.fps = !Dbg.fps
		}

	}
	dbgPack := []bool{
		Dbg.tps, Dbg.tick,
		Dbg.screen, Dbg.windowPX,
		Dbg.frames, Dbg.layouts,
		Dbg.fps, Dbg.khandlr,
		Dbg.keys, Dbg.keysQty}
	genPack := []msgGen{
		debugTPS, debugTick10,
		debugScreen, debugPX,
		debugFrames10, debugLayouts10,
		debugFPS, debugKeyHandler,
		debugKeys, debugKeysQty}
	Dbg.Output = stringMerge(dbgPack, genPack)
}

//*==Debug MsgGen Functions===============================================

// stringMerge merges strings. if bool[i] in shouldWrite is true, result of gen[i] is appended.
func stringMerge(shouldWrite []bool, gen []msgGen) string {
	var sb strings.Builder
	tf := 0
	Dbg.len = 0
	for i, bgen := range shouldWrite {
		if bgen {
			tf, _ = sb.WriteString(gen[i]())
			Dbg.len += tf
		}
	}
	return sb.String()
}

func debugTPS() string {
	return fmt.Sprintf("| tps: %f ", ebiten.ActualTPS())
}
func debugFPS() string {
	return fmt.Sprintf("| fps: %f ", ebiten.ActualFPS())
}
func debugTick10() string {
	return fmt.Sprintf("| tick: %03d ", tick/10)
}
func debugFrames10() string {
	return fmt.Sprintf("| frames: %03d ", frame/10)
}
func debugScreen() string {
	return fmt.Sprintf("| game/screen: %dx%d ", Dbg.gameW, Dbg.gameH)
}
func debugPX() string {
	return fmt.Sprintf("| px: %dx%d ", Dbg.pixW, Dbg.pixH)
}
func debugLayouts10() string {
	return fmt.Sprintf("| layout: %d ", layoutCount/10)
}
func debugKeyHandler() string {
	kqty := 0
	kstr := ""
	keys := input.KeysOut()
	for _, v := range *keys {
		kstr += v.String()
	}
	return fmt.Sprintf("\n| KBHandl[qty %d,len %d]: %s", kqty, len(*keys), kstr)
}

func debugKeys() string {
	var keysStr string = "\n"
	for _, v := range Dbg.keysDown {
		keysStr += v.String()
	}
	return keysStr
}
func debugKeysQty() string {
	return fmt.Sprintf(" | len key: %d", len(Dbg.keysDown))
}

// DbgCountFrames will run each draw call. Max at 2800 (arbitrary) then resets
func DbgCountFrames() {
	frame++
	if frame >= 2800 {
		frame = 0
	}
}

// DbgCountLayout will run each Layout call. Max at 2800 (arbitrary) then resets
func DbgCountLayout() {
	layoutCount++
	if layoutCount >= 2800 {
		layoutCount = 0
	}
}

// DbgCountTicks will run each Update call. Max at 1800 (arbitrary) then resets
func DbgCountTicks() {
	tick++
	if tick >= 1800 {
		tick = 0
	}
}

// DbgCaptureInput runs in Update to get pressed keys for debug string
func DbgCaptureInput() {
	Dbg.keysDown = inpututil.AppendPressedKeys(Dbg.keysAppend[:0])
}
