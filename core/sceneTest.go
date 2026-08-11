package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
)

type testScene struct {
	Display *gfx.WindowConfig
}

var ts *testScene = &testScene{}

func init() {
}

func (s *testScene) Update() error {

	return nil
}

func (s *testScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
}
