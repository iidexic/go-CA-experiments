package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/input"
)

var assignedKeys = []ebiten.Key{ebiten.KeyE,
	ebiten.KeyG, ebiten.KeyR, ebiten.KeyQ,
	ebiten.KeyD, ebiten.KeyC, ebiten.KeySpace,
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
	m := input.Mouse()
	if m.CursOn(g.maingrid.Bounds) > 0 {

	}
	// if m.CursOn() == 1 {}
	g.presstime(assignedKeys)
}

func highlightcursor(m input.EZmouse) {
	if m.CursOn([]int{0, 0, 480, 240}) > 0 {

	}
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
	case ebiten.KeyC:
	case ebiten.KeyE:
		/*//[previous troubleshooting]
		cbt := gfx.GetQuickRNG(8)
		cbt.ROPcheck()
		bt := []byte{<-cbt.C, <-cbt.C, <-cbt.C, <-cbt.C, <-cbt.C, <-cbt.C, <-cbt.C}
		fmt.Printf("[%b.%b.%b.%b.%b.%b.%b]\n", bt[0], bt[1], bt[2], bt[3], bt[4], bt[5], bt[6])
		*/
	case ebiten.KeyR:
		g.maingrid.Px = gfx.Randpx(uint(g.maingrid.Area))
		g.maingrid.Img.WritePixels(g.maingrid.Px)
	case ebiten.KeyQ:
		g.maingrid.Debug = !g.maingrid.Debug
	case ebiten.KeyEnter:
		g.SimSpeed = -g.SimSpeed
		if g.SimSpeed == 0 {
			g.SimSpeed += 5
		}
	case ebiten.KeyArrowUp:
		if g.SimSpeed < 12 {
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
		} else if g.SimSpeed >= 56 {
			g.SimSpeed -= 8
		}

	case ebiten.KeyArrowLeft:
		entity.CutoffDown()
	case ebiten.KeyArrowRight:
		entity.CutoffUp()
	}
}
