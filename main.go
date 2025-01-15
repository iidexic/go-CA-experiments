package main

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/util"
	"golang.org/x/image/font/gofont/goregular"
)

// globals and Structs (temporary)
// ==================================

// :)
var (
	PixWidth    int        = 1280
	PixHeight   int        = 720
	GameWidth   int        = 640
	GameHeight  int        = 360
	bgcolor     color.RGBA = color.RGBA{32, 29, 31, 255}
	latch       bool       = true
	tick, frame uint       = 0, 0
	layoutCount int        = 0
)

type testobj struct {
	img *ebiten.Image
	r   image.Rectangle
	op  *ebiten.DrawImageOptions
	set bool
}

// ==================================
// ==================================

func testMove(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(1, 1)
}
func inputActions(g *Game) {
	_, wy := ebiten.Wheel()
	if wy > 0 {

		latch = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {

		g.maingrid.Draw = true
	}

}

// ==================================

// Game struct - ebiten
type Game struct { //^-GAME STRUCT-
	maingrid entity.GridEntity
}

// Update game - game logic, assume locked at 60TPS.
func (g *Game) Update() error { //^UPDATE
	defer util.DebugMsgControl(GameWidth, GameHeight, PixWidth, PixHeight)
	util.DbgCountTicks()
	util.DbgCaptureInput()
	//the input stuff ain't working. Checking in Debug
	//input.UpdateKeys()
	inputActions(g)
	return nil
}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) { //^DRAW
	util.DbgCountFrames()

	screen.Fill(bgcolor)

	screen.DrawImage(g.maingrid.Img, &g.maingrid.Op)
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
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	_ = s
	if err != nil {
		log.Panicf("Font did not load %s", err)
	}
	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}
	g.maingrid = entity.MakeGridDefault(GameWidth, GameHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
