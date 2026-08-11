package util

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type showDebugInfo struct { // bools toggle what gets put into debug msg. len(Output)
	showDebug                        bool
	len, pixW, pixH, gameW, gameH    int
	Output, UpdateDetail, DrawDetail string
	keysAppend                       []ebiten.Key
	keysDown                         []ebiten.Key
	SelectDebug                      []debugstat
	errorText                        string
}

type debugstat int

// Show consts determine which debug messages are included onscreen
const (
	ShowTPS debugstat = iota //0
	ShowTick
	ShowFrames
	ShowScreen
	ShowLayouts
	ShowWindowPX
	ShowFPS
	ShowKhandlr
	ShowUpdateDetail
	ShowDrawDetail
	ShowMouseDetail
	ShowErrorText
	newline
)

var (
	frame, tick, layoutCount int
)

//TODO: Rebuild so doesn't import other pacakges.

// Dbg houses all required information/settings for debug messages
var Dbg showDebugInfo = showDebugInfo{
	showDebug:  true,
	len:        0,
	gameW:      0,
	gameH:      0,
	pixW:       0,
	pixH:       0,
	keysAppend: make([]ebiten.Key, 0, 12),
	SelectDebug: []debugstat{ShowWindowPX, ShowScreen, ShowTPS, ShowFPS, ShowMouseDetail,
		newline, ShowUpdateDetail, ShowDrawDetail,
		newline, ShowErrorText},
}

func (d *showDebugInfo) WriteDraw(str string)   { d.DrawDetail = str }
func (d *showDebugInfo) WriteUpdate(str string) { d.UpdateDetail = str }
func (d *showDebugInfo) AddStat(stat debugstat) { d.SelectDebug = append(d.SelectDebug, stat) }
func (d *showDebugInfo) RemoveStat(stat debugstat) {
	for i, v := range d.SelectDebug {
		if v == stat {
			d.SelectDebug = slices.Delete(d.SelectDebug, i, i+1)
			break
		}
	}
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
		case ShowTPS:
			_, e = sb.WriteString(fmt.Sprintf("| tps: %0.0f ", ebiten.ActualTPS()))
		case ShowTick:
			_, e = sb.WriteString(fmt.Sprintf("| tick: %03d ", tick/10))
		case ShowFrames:
			_, e = sb.WriteString(fmt.Sprintf("| frames: %03d ", frame/10))
		case ShowScreen:
			_, e = sb.WriteString(fmt.Sprintf("| game/screen: %dx%d ", d.gameW, d.gameH))
		case ShowLayouts:
			_, e = sb.WriteString(fmt.Sprintf("| layout: %d ", layoutCount/10))
		case ShowWindowPX:
			_, e = sb.WriteString(fmt.Sprintf("| px: %dx%d ", d.pixW, d.pixH))
		case ShowFPS:
			_, e = sb.WriteString(fmt.Sprintf("| fps: %0.0f ", ebiten.ActualFPS()))
		case ShowKhandlr: // TODO: add new input handling system back in here
			// kstr := ""
			// keys := input.KeysOut()
			// for _, k := range *keys {
			// 	kstr += k.String()
			// }
			// _, e = sb.WriteString(fmt.Sprintf("| inKB[len %d]: %s", len(*keys), kstr))
		case ShowUpdateDetail:
			_, e = sb.WriteString(d.UpdateDetail)
		case ShowDrawDetail:
			_, e = sb.WriteString(d.DrawDetail)
		case ShowMouseDetail:
			mX, mY := ebiten.CursorPosition()
			_, e = sb.WriteString(fmt.Sprintf("| pos:(%3d,%3d), keys:%s", mX, mY, dbgGetMouse()))
		case ShowErrorText:
			_, e = sb.WriteString(d.errorText)
		case newline:
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

func (d *showDebugInfo) AddErrorF(ftext string, args ...any) {
	if ftext != "" {
		d.errorText += fmt.Sprintf("["+ftext+"]", args...)
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
