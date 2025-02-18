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
		g.maingrid.Px = gfx.Randpx(uint(g.maingrid.Area))
		g.maingrid.Img.WritePixels(g.maingrid.Px)
	case ebiten.KeyQ:

	case ebiten.KeyEnter:
		g.SimSpeed = -g.SimSpeed
		if g.SimSpeed == 0 {
			g.SimSpeed += 5
		}
	case ebiten.KeyArrowUp:
		if 0 < g.SimSpeed && g.SimSpeed < 12 {
			g.SimSpeed++
		} else if g.SimSpeed < 24 {
			g.SimSpeed += 2
		} else if g.SimSpeed < 36 {
			g.SimSpeed += 4
		} else if g.SimSpeed < 56 {
			g.SimSpeed += 8
		}
	case ebiten.KeyArrowDown:
		if 1 < g.SimSpeed && g.SimSpeed < 12 {
			g.SimSpeed--
		} else if g.SimSpeed < 24 {
			g.SimSpeed -= 2
		} else if g.SimSpeed < 36 {
			g.SimSpeed -= 4
		} else if g.SimSpeed < 56 {
			g.SimSpeed -= 8
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
