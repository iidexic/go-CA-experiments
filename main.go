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

func (obj *testobj) genSquare() {

	//*image.Rect (x0,y0)(x1,y1) does not translate. only (x1-x0, y1-y0) is used
	obj.r = image.Rect(0, 0, 100, 100)
	options := ebiten.NewImageOptions{}
	obj.img = ebiten.NewImageWithOptions(obj.r, &options)
	obj.op = &ebiten.DrawImageOptions{}

	testMove(obj.op)
	obj.set = true
}

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

/* etxt package text rendering.
I am just going to try ebiten's builtin
	g.txt = etxt.NewRenderer()
g.txt.SetFont(lbrtmono.Font())
g.txt.Utils().SetCache8MiB()
*/
