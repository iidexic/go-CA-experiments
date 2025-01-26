package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/input"
)

type msgGen func() string
type msgScreenGen func(int, int) string
type showDebugInfo struct { // bools toggle what gets put into debug msg. len(Output)
	showDebug                        bool
	len, pixW, pixH, gameW, gameH    int
	Output, UpdateDetail, DrawDetail string
	keysAppend                       []ebiten.Key
	keysDown                         []ebiten.Key
	SelectDebug                      []int
}

// Show consts determine which debug messages are included onscreen
const (
	showTPS int = iota //0
	showTick
	showFrames
	showScreen
	showLayouts
	showWindowPX
	showFPS
	showKhandlr
	showUpdateDetail
	showDrawDetail
	showMouseDetail
	nl
)

var (
	frame, tick, layoutCount int
)

// Dbg houses all required information/settings for debug messages
var Dbg showDebugInfo = showDebugInfo{
	showDebug:   true,
	len:         0,
	gameW:       0,
	gameH:       0,
	pixW:        0,
	pixH:        0,
	keysAppend:  make([]ebiten.Key, 0, 12),
	SelectDebug: []int{showTPS, showFPS, nl, showMouseDetail, showKhandlr, showUpdateDetail, showDrawDetail},
}

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
			_, e = sb.WriteString(fmt.Sprintf("| tps: %0.0f ", ebiten.ActualTPS()))
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
			_, e = sb.WriteString(fmt.Sprintf("| fps: %0.0f ", ebiten.ActualFPS()))
		case showKhandlr:
			kstr := ""
			keys := input.KeysOut()
			for _, k := range *keys {
				kstr += k.String()
			}
			_, e = sb.WriteString(fmt.Sprintf("| inKB[len %d]: %s", len(*keys), kstr))
		case showUpdateDetail:
			_, e = sb.WriteString(d.UpdateDetail)
		case showDrawDetail:
			_, e = sb.WriteString(d.DrawDetail)
		case showMouseDetail:
			mX, mY := ebiten.CursorPosition()
			_, e = sb.WriteString(fmt.Sprintf("| pos:(%3d,%3d), keys:%s", mX, mY, dbgGetMouse()))
		case nl:
			_, e = sb.WriteString("\n")
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

// *==Debug MsgGen Functions===============================================
func dbgGetMouse() string { //crusty mb get func, works tho
	btnstr := ""
	if ebiten.IsMouseButtonPressed(0) {
		btnstr += "lmb "
	}
	if ebiten.IsMouseButtonPressed(1) {
		btnstr += "rmb "
	}
	if ebiten.IsMouseButtonPressed(2) {
		btnstr += "mmb "
	}
	if ebiten.IsMouseButtonPressed(3) {
		btnstr += "mb4 "
	}
	if ebiten.IsMouseButtonPressed(4) {
		btnstr += "mb5 "
	}
	return btnstr
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
