package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/input"
	"github.com/iidexic/go-CA-experiments/util"
	"golang.org/x/image/font/gofont/goregular"
)

// globals and Structs (temporary)
// ==================================

// :)
var (
	PixWidth    int  = 1280
	PixHeight   int  = 720
	GameWidth   int  = 640
	GameHeight  int  = 360
	tick, frame uint = 0, 0
	layoutCount int  = 0
)

// ==================================
// ==================================

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
		g.simShift++
		if g.simShift > g.simBitmax {
			g.simShift = g.simBitmax
		}
	case ebiten.KeyArrowDown:
		if g.simShift > 1 {
			g.simShift--
		}
	case ebiten.KeyArrowLeft:
		if g.simBitmax > 0 {
			g.simBitmax--
		}
	case ebiten.KeyArrowRight:
		if g.simBitmax < 28 { //somewhat arbitrary
			g.simBitmax++
		}
	}

}

//==================================

// Game struct - ebiten
type Game struct { //^-GAME STRUCT-
	maingrid                           *entity.GridEntity
	pal                                []color.RGBA
	RunSimulation, simShift, simBitmax int
}

// Update game - game logic, assume locked at 60TPS.
func (g *Game) Update() error { //^UPDATE
	defer util.Dbg.DebugBuildOutput()
	util.DbgCountTicks()
	input.GetInKB()
	//util.DbgCaptureInput()
	//the input stuff ain't working. Checking in Debug
	if g.RunSimulation > 0 {
		g.maingrid.TestSimulate(g.simShift, g.simBitmax)
		g.maingrid.Img.WritePixels(g.maingrid.Pixels)
	}
	inputActions(g)
	return nil
}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) { //^DRAW
	util.DbgCountFrames()
	util.DbgCaptureInput()
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
	return GameWidth, GameHeight
}

func main() {
	//^=====| INITIALIZATION IN MAIN |=====
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	_ = s
	if err != nil {
		log.Panicf("Font did not load %s", err)
	}
	//--- temp above here ---
	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{
		RunSimulation: 0,
		simShift:      1,
		simBitmax:     4,
		maingrid:      entity.MakeGridDefault(GameWidth, GameHeight),
		pal:           gfx.Palette,
	}
	//^====================================
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
