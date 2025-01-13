package util

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type debugMessager interface {
	msgGen | msgScreenGen
}
type msgGen func() string
type msgScreenGen func(int, int) string
type showDebugInfo struct { // bools toggle what gets put into debug msg. len(Output)
	showDebug, tps, tick, screen, layouts, windowPX, frames, fps bool
	len, pixW, pixH, gameW, gameH                                int
	Output                                                       string
}

var (
	frame, tick, layoutCount         int
	gWidth, gHeight, pWidth, pHeight int
	//Dbg is used to toggle desired debug output, receive screen info, and houses the final output
	Dbg showDebugInfo = showDebugInfo{
		showDebug: true,
		tps:       true,
		tick:      true,
		frames:    false,
		screen:    false,
		layouts:   false,
		windowPX:  false,
		fps:       true,
		len:       0,
		gameW:     0,
		gameH:     0,
		pixW:      0,
		pixH:      0,
		Output:    ""}
)

// TODO: make proper input handling.

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
	dbgPack := []bool{Dbg.tps, Dbg.tick, Dbg.screen, Dbg.windowPX, Dbg.frames, Dbg.layouts, Dbg.fps}
	genPack := []msgGen{debugTPS, debugTick10, debugScreen, debugPX, debugFrames10, debugLayouts10, debugFPS}
	Dbg.Output = stringMerge(dbgPack, genPack)
}

// DrawDebug is run in the Draw loop to display debug stats at the top.
func DrawDebug() {

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

// CountFrames will run each draw call. Max at 2800 (arbitrary) then resets
func DbgCountFrames() {
	frame++
	if frame >= 2800 {
		frame = 0
	}
}

// CountLayout will run each Layout call. Max at 2800 (arbitrary) then resets
func DbgCountLayout() {
	layoutCount++
	if layoutCount >= 2800 {
		layoutCount = 0
	}
}

// CountTicks will run each Update call. Max at 1800 (arbitrary) then resets
func DbgCountTicks() {
	tick++
	if tick >= 1800 {
		tick = 0
	}
}
