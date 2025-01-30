package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/iidexic/go-CA-experiments/core"
	"golang.org/x/image/font/gofont/goregular"
)

// globals and Structs (temporary)
// ==================================

// :)
var (
	PixWidth    int  = 1280
	PixHeight   int  = 720
	GameWidth   int  = 960 //16
	GameHeight  int  = 540 //9
	tick, frame uint = 0, 0
	layoutCount int  = 0
)

// ==================================

func main() {
	//^=====| INITIALIZATION IN MAIN |=====
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	_ = s
	if err != nil {
		log.Panicf("Font did not load %s", err)
	}
	//--- temp above here ---
	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowTitle("CA Experimentor")
	g := core.GameInit(GameWidth, GameHeight)
	//^====================================
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
