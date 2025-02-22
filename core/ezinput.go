package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
)

var assignedKeys = []ebiten.Key{ebiten.KeyE, ebiten.KeyG, ebiten.KeyR, ebiten.KeyQ,
	ebiten.KeyArrowDown, ebiten.KeyArrowUp,
	ebiten.KeyArrowLeft, ebiten.KeyArrowRight,
	ebiten.KeyEnter}

func inputActions(g *GameSim) {
	//cursX,cursY:=ebiten.CursorPosition()
	_, wy := ebiten.Wheel()
	if wy > 0 {
		//mouseWheelUp
	} else if wy < 0 {
		//mouseWheelDown
	}
	g.presstime(assignedKeys)
	/* replacing with presstime
	kbKeys := input.GetJustPressedKeys()
	var kbHold []ebiten.Key = make([]ebiten.Key, 0, 24)
	for _, k := range kbKeys {
		g.callKey(k)
	}
	*/
}
func (g *GameSim) presstime(kbKeys []ebiten.Key) {
	for _, key := range kbKeys {
		//= repeat behavior
		intime := inpututil.KeyPressDuration(key)
		if intime == 1 || (intime > 20 && intime%6 == 0) {
			g.callKey(key)
		}
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
		entity.CutoffDown()
	case ebiten.KeyArrowRight:
		entity.CutoffUp()
	}
}
