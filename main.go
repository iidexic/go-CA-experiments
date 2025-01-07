package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// globals and Structs (temporary)
// ==================================
const (
	pixWidth, pixHeight   int = 1280, 720
	gameWidth, gameHeight int = 640, 360
)

var (
	fillcolor   color.RGBA
	latch       bool = true
	tick, frame uint = 0, 0
	layoutCount int  = 0
	sqri        *ebiten.Image
	sqrobj      *testobj
)

type testobj struct {
	img *ebiten.Image
	r   image.Rectangle
	op  *ebiten.DrawImageOptions
	set bool
}

// ==================================
// My Functions (temporary)
// ==================================

func (obj *testobj) genSquare() {

	//*image.Rect (x0,y0)(x1,y1) does not translate. only (x1-x0, y1-y0) is used
	obj.r = image.Rect(0, 0, 100, 100)
	options := ebiten.NewImageOptions{}
	obj.img = ebiten.NewImageWithOptions(obj.r, &options)
	obj.op = &ebiten.DrawImageOptions{}

	obj.img.WritePixels()

	testMove(obj.op)
	obj.set = true
}

func imagenoise(img *ebiten.Image) {
	area := uint(img.Bounds().Dx() * img.Bounds().Dy())
	img.WritePixels(Randpx(area))
}

func testMove(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(1, 1)
}
func inputActions() {
	_, wy := ebiten.Wheel()
	if wy > 0 {

		latch = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.maingrid.Defaults()
		g.maingrid.draw = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		imagenoise(sqrobj.img)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		testMove(sqrobj.op)
	}
}

// TODO - determine best  update/draw relation
// ==================================
// ANCHOR ====[Ebiten game base]====
// ==================================

// Game struct - ebiten
type Game struct {
	maingrid GridEntity
}

// Update game - game logic, assume locked at 60TPS.
func (g *Game) Update() error {
	defer debugMsgControl()
	countTicks()
	inputActions()
	return nil
}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) {
	countFrames()

	if !latch { //FIXME - replace this
		sqrobj.genSquare()
		latch = true
	}

	screen.Fill(fillcolor)
	if sqrobj.set {
		screen.DrawImage(sqrobj.img, sqrobj.op)
	}
	ebitenutil.DebugPrintAt(screen, dbg.output, 120, 0)

}

// Layout of game window (screen/game)
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	layoutCount++
	return gameWidth, gameHeight
}

func main() {
	//? best practice on initialization? Here, or in update, or in separate init function?

	sqrobj = &testobj{}
	sqrobj.img = ebiten.NewImage(1, 1)
	sqrobj.set = false

	ebiten.SetWindowSize(pixWidth, pixHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{maingrid: makeGridDefault()}); err != nil {
		log.Fatal(err)
	}
}
