package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// globals and Structs (temporary)
// ==================================
var (
	fillcolor color.RGBA
	latch     bool = true
	tick      uint = 0
	sqri      *ebiten.Image
	sqrobj    *testobj
)

type testobj struct {
	img        *ebiten.Image
	r          image.Rectangle
	op         *ebiten.DrawImageOptions
	set, drawn bool
}

// ==================================
// My Functions (temporary)
// ==================================

func debugMsg() string {
	tps := ebiten.ActualTPS()
	tick++
	if tick >= 1000 {
		tick = 0
	}
	msg := fmt.Sprintf("tps: %f, tick: %03d", tps, tick)
	return msg
}

func ezgenSquare() testobj {
	var obj testobj

	obj.r = image.Rect(10, 10, 0, 0)
	options := ebiten.NewImageOptions{}
	obj.img = ebiten.NewImageWithOptions(obj.r, &options)
	obj.op = &ebiten.DrawImageOptions{}

	obj.img.Fill(color.RGBA{245, 10, 20, 255})
	movearound(obj.op)
	return obj
}
func (obj *testobj) genSquare() {

	obj.r = image.Rect(10, 10, 0, 0)
	options := ebiten.NewImageOptions{}
	obj.img = ebiten.NewImageWithOptions(obj.r, &options)
	obj.op = &ebiten.DrawImageOptions{}

	obj.img.Fill(color.RGBA{245, 10, 20, 255})
	movearound(obj.op)
	obj.set = true
}
func giveSquarePlease() *testobj {
	obj := testobj{
		img: ebiten.NewImage(1, 1),
		op:  &ebiten.DrawImageOptions{},
	}
	return &obj
}

/*
func imagenoise(img *ebiten.Image) {
	r := img.Bounds()
	inew := image.NewRGBA(r)
	for y := range r.Dy() {
		for x := range r.Dx() {
			img.WritePixels()
		}
	}
}*/

func movearound(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(1, 1)
}

// for WritePixels len(pix) must be 4*imageWidth*imageHeight
//Pix = []byte { R, G, B, A, R, G, B, A, R, G, B, A}
// Image interface has value [?] that aids in maneuvering pixel rows
//NOTE Good Practices
// - GeoM moves BEFORE fill/write pixel?
// ==================================
// ANCHOR ====[Ebiten game base]====
// ==================================

// Game struct - ebiten
type Game struct {
}

// Update game
func (g *Game) Update() error {
	return nil
}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) {

	_, wy := ebiten.Wheel()
	if wy > 0 {
		fillcolor = Randcolor()
		latch = false
	}

	//* preventing re-generation of ezgenSquare each frame. Possibly a better solution?
	if !latch {
		sqrobj.genSquare()
		latch = true
	}

	screen.Fill(fillcolor)
	if sqrobj.set {
		screen.DrawImage(sqrobj.img, sqrobj.op)
	}
	ebitenutil.DebugPrintAt(screen, debugMsg(), 120, 0)

}

// Layout of game window (Original width height 320x240, this is gameworld  size/scale)
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 384, 256
}

func main() {
	sqrobj = &testobj{}
	sqrobj.img = ebiten.NewImage(1, 1)
	sqrobj.set = false
	ebiten.SetWindowSize(768, 512)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// !=====================================================================
// !=====================================================================
// old
/*
func genTestSquare() *testsquare {
	sqr := testsquare{sizex: 10, sizey: 10, locx: 20.0, locy: 20.0}
	sqr.clr = color.RGBA{R: 0xf0, G: 0x0b, B: 0x1f, A: 0xff}
	sqr.i = ebiten.NewImage(sqr.sizex, sqr.sizey)
	sqr.i.Fill(sqr.clr)
	//sqr.geo.Translate(sqr.locx, sqr.locy)
	//sqr.drawoption.GeoM = sqr.geo
	return &sqr
}
*/
