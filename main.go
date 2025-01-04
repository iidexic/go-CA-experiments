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
var fillcolor color.RGBA

type testsquare struct {
	sizex, sizey int
	locx, locy   float64
	px           []byte
	i            *ebiten.Image
	clr          color.RGBA
	geo          ebiten.GeoM
	drawoption   *ebiten.DrawImageOptions
}

// ==================================
// My Functions (temporary)
// ==================================

func debugMsg() string {
	tps := ebiten.ActualTPS()

	msg := fmt.Sprintf("tps: %f", tps)
	return msg
}
func genTestSquare() *testsquare {
	sqr := testsquare{sizex: 10, sizey: 10, locx: 20.0, locy: 20.0}
	sqr.clr = color.RGBA{R: 0xf0, G: 0x0b, B: 0x1f, A: 0xff}
	sqr.i = ebiten.NewImage(sqr.sizex, sqr.sizey)
	sqr.i.Fill(sqr.clr)
	sqr.geo.Translate(sqr.locx, sqr.locy)
	sqr.drawoption.GeoM = sqr.geo
	return &sqr
}

func ezgenSquare() *ebiten.Image {
	r := image.Rect(10, 10, 20, 20)
	options := ebiten.NewImageOptions{}
	img := ebiten.NewImageWithOptions(r, &options)
	img.Fill(color.RGBA{245, 10, 20, 255})
	return img
}

// for WritePixels len(pix) must be 4*imageWidth*imageHeight

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
	}
	sqr := ezgenSquare()
	screen.Fill(fillcolor)
	screen.DrawImage(sqr, nil)
	ebitenutil.DebugPrint(screen, debugMsg())
}

// Layout of game window
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
