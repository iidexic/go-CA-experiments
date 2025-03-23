package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/core"
)

// globals and Structs
// ==================================
var ( //16 by 9: 1920x1080, 960x540
	PixWidth    int  = 640
	PixHeight   int  = 1000
	GameWidth   int  = 320
	GameHeight  int  = 500
	tick, frame uint = 0, 0
	layoutCount int  = 0
)

// ==================================
func main() {
	//s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("CA Experimentor")
	ebiten.SetWindowPosition(0, 80)
	g := core.GameSimInit(GameWidth, GameHeight)
	//>----/ launch game loop /----<//
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
