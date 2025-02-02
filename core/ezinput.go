package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/input"
)

func inputActions(g *GameSim) {
	//cursX,cursY:=ebiten.CursorPosition()
	_, wy := ebiten.Wheel()
	if wy > 0 {
		//mouseWheelUp
	} else if wy < 0 {
		//mouseWheelDown
	}
	for _, k := range input.GetJustPressedKeys() {
		g.callKey(k)
	}

}
func (g *GameSim) callKey(k ebiten.Key) {
	switch k {
	case ebiten.KeyG:
		g.maingrid.Draw = !g.maingrid.Draw
	case ebiten.KeyE:

	case ebiten.KeyR:
		g.maingrid.Pixels = gfx.Randpx(uint(g.maingrid.Area))
		g.maingrid.Img.WritePixels(g.maingrid.Pixels)
	case ebiten.KeyQ:

	case ebiten.KeyEnter:
		if g.RunSimulation > 0 {
			g.RunSimulation = 0
		} else {
			g.RunSimulation = 1
		}
	case ebiten.KeyArrowUp:
		g.modAdd++
		if g.modAdd < 16 {
			g.modAdd = g.modMult
		}
	case ebiten.KeyArrowDown:
		if g.modAdd > 0 {
			g.modAdd--
		}
	case ebiten.KeyArrowLeft:
		if g.modMult > 0 {
			g.modMult--
		}
	case ebiten.KeyArrowRight:
		if g.modMult < 16 { //somewhat arbitrary
			g.modMult++
		}
	}
}
