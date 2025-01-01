package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func debugMsg() string {
	tps := ebiten.ActualTPS()

	msg := fmt.Sprintf("tps: %f", tps)
	return msg
}

// ==================================
// ANCHOR ====[Ebiten game base]====
// ==================================
var fillcolor color.RGBA

// Game struct - ebiten
type Game struct {
}

// Update game
func (g *Game) Update() error {
	return nil
}
func initColor(screen *ebiten.Image) {

}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) {
	_, wy := ebiten.Wheel()
	if wy > 0 {
		fillcolor = randcolor()
	}

	screen.Fill(fillcolor)
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
