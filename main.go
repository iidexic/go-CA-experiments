package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/core"
)

// globals and Structs
// ==================================
var (
	PixWidth    int  = 1280
	PixHeight   int  = 720
	GameWidth   int  = 320 //16 (960)
	GameHeight  int  = 180 //9 (540)
	tick, frame uint = 0, 0
	layoutCount int  = 0
)

// ==================================
func main() {
	//s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("CA Experimentor")
	g := core.GameSimInit(GameWidth, GameHeight)
	//>----/ launch game loop /----<//
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
