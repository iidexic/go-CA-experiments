package core

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/input"
	"github.com/iidexic/go-CA-experiments/util"
)

func testMove(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(1, 1)
}
func inputActions(g *Game) {
	//cursX,cursY:=ebiten.CursorPosition()
	_, wy := ebiten.Wheel()
	if wy > 0 {
		//mouseWheelUp
	} else if wy < 0 {
		//mouseWheelDown
	}
	for _, k := range input.GetJustPressedKeys() {
		g.callKey(k)
	}

}
func (g *Game) callKey(k ebiten.Key) {
	switch k {
	case ebiten.KeyG:
		g.maingrid.Draw = !g.maingrid.Draw
	case ebiten.KeyE:

	case ebiten.KeyR:
		g.maingrid.Pixels = gfx.Randpx(uint(g.maingrid.Area))
		g.maingrid.Img.WritePixels(g.maingrid.Pixels)
	case ebiten.KeyQ:

	case ebiten.KeyEnter:
		if g.RunSimulation > 0 {
			g.RunSimulation = 0
		} else {
			g.RunSimulation = 1
		}
	case ebiten.KeyArrowUp:
		g.modAdd++
		if g.modAdd < 16 {
			g.modAdd = g.modMult
		}
	case ebiten.KeyArrowDown:
		if g.modAdd > 0 {
			g.modAdd--
		}
	case ebiten.KeyArrowLeft:
		if g.modMult > 0 {
			g.modMult--
		}
	case ebiten.KeyArrowRight:
		if g.modMult < 16 { //somewhat arbitrary
			g.modMult++
		}
	}

}

// GameInit returns Game pointer for main sim scene with default settings
func GameInit(gameWidth, gameHeight int) *Game {
	g := &Game{
		RunSimulation: 0,
		modAdd:        1,
		modMult:       4,
		gWidth:        gameWidth,
		gHeight:       gameHeight,
		maingrid:      entity.MakeGridDefault(gameWidth, gameHeight),
		pal:           gfx.Palette,
	}
	return g
}

//==================================

// Game struct - ebiten
type Game struct { //^-GAME STRUCT-
	maingrid                         *entity.GridEntity
	pal                              []color.RGBA
	gWidth, gHeight, pWidth, pHeight int
	RunSimulation, modAdd, modMult   int
}

// Update game - game logic, assume locked at 60TPS.
func (g *Game) Update() error { //^UPDATE
	defer util.Dbg.DebugBuildOutput()
	util.DbgCountTicks()
	util.Dbg.UpdateDetail = fmt.Sprintf("||U:add/shift=%d, multBitmax=%d", g.modAdd, g.modMult)
	input.GetInKB()
	//util.DbgCaptureInput()
	if g.RunSimulation > 0 {
		g.maingrid.SetMod(g.modAdd, g.modMult)
		g.maingrid.SimstepLVSD(true)
		//g.maingrid.TestSimulate(g.modAdd, g.modMult) //this is original simulate call
		g.maingrid.Img.WritePixels(g.maingrid.Pixels)
	}
	inputActions(g)
	return nil
}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) { //^DRAW
	util.DbgCountFrames()
	screen.Fill(g.pal[0])

	if g.maingrid.Draw {
		screen.DrawImage(g.maingrid.Img, &g.maingrid.Op)
	}
	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)

}

// Layout of game window (screen/game)
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	util.DbgCountLayout()
	//TODO: Write functionality for scaling.
	// i.e. : Maintain Ratio, relative screen position
	// and maintain size of game area!! (i.e. not adding additional pixels to sim  by resizing)
	return g.gWidth, g.gHeight
}
