package main

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type msgGen func() string

type showDebugInfo struct {
	showDebug, tps, tick, screen, layouts, windowPX, frames bool
	len                                                     int
	output                                                  string
}

var dbg showDebugInfo = showDebugInfo{
	showDebug: true,
	tps:       false,
	tick:      false,
	frames:    false,
	screen:    false,
	layouts:   false,
	windowPX:  false,
	len:       0,
	output:    ""}

// debugMsgControl puts together debug messages. Uncertain of impact to performance
// TODO: make proper input handling.
func debugMsgControl() {
	if !dbg.showDebug {
		dbg.output = "debug should be off"
	}

	if ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			dbg.tps = !dbg.tps
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			dbg.tick = !dbg.tick
		} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
			dbg.screen = !dbg.screen
		} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
			dbg.layouts = !dbg.layouts
		} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
			dbg.windowPX = !dbg.windowPX
		} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
			dbg.frames = !dbg.frames
		}

	}
	dbgPack := []bool{dbg.tps, dbg.tick, dbg.screen, dbg.windowPX, dbg.frames, dbg.layouts}
	genPack := []msgGen{debugTPS, debugTick10, debugScreen, debugPX, debugFrames10, debugLayouts10}
	dbg.output = stringMerge(dbgPack, genPack)
}

// stringMerge merges strings. if bool[i] in shouldWrite is true, result of gen[i] is appended.
func stringMerge(shouldWrite []bool, gen []msgGen) string {
	var sb strings.Builder
	tf := 0
	dbg.len = 0
	for i, bgen := range shouldWrite {
		if bgen {
			tf, _ = sb.WriteString(gen[i]())
			dbg.len += tf
		}
	}
	return sb.String()
}

func debugTPS() string {
	return fmt.Sprintf("| tps: %f ", ebiten.ActualTPS())
}
func debugTick10() string {
	return fmt.Sprintf("| tick: %03d ", tick/10)
}
func debugFrames10() string {

	return fmt.Sprintf("| frames: %03d ", frame/10)
}

func debugScreen() string {
	return fmt.Sprintf("| game/screen: %dx%d ", gameWidth, gameHeight)
}

func debugPX() string {
	return fmt.Sprintf("| px: %dx%d ", pixWidth, pixHeight)
}
func debugLayouts10() string {
	return fmt.Sprintf("| layout: %d ", layoutCount/10)
}
func countFrames() {
	frame++
	if frame >= 2800 {
		frame = 0
	}
}
func countLayout() {
	layoutCount++
	if layoutCount >= 2800 {
		layoutCount = 0
	}
}
func countTicks() {
	tick++
	if tick >= 1800 {
		tick = 0
	}
}
