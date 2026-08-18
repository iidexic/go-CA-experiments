package core

import (
	"image"

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
	ebiten.KeyEnter, ebiten.KeyEscape}

func inputActions(g *GameSim) {
	cx, cy := ebiten.CursorPosition()
	cursor := image.Pt(cx, cy)

	_, wy := ebiten.Wheel()
	if wy > 0 {
		g.aperture.Zoom(1, cursor)
	} else if wy < 0 {
		g.aperture.Zoom(-1, cursor)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.aperture.BeginPan(cursor)
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.aperture.UpdatePan(cursor)
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		g.aperture.EndPan()
	}

	m := input.Mouse()
	if m.CursOn(g.maingrid.Bounds) > 0 {
		// reserved for cell selection wiring
	}
	g.presstime(assignedKeys)
}

/*
	func highlightcursor(m input.EZmouse, effectBounds []int) ebiten.Image {
		if m.CursOn(effectBounds) > 0 {
			IMG:=ebiten.NewImage(12,12)
			IMG.Fill(color.RGBA{60,50,10,60})

		}
	}
*/
func (g *GameSim) presstime(kbKeys []ebiten.Key) {
	for _, key := range kbKeys {
		//= repeat behavior
		intime := inpututil.KeyPressDuration(key)
		if intime == 1 || (intime > 20 && intime%6 == 0) {
			g.callKey(key)
		}
	}
}
func (g *GameSim) fastInitializeDev() {
	g.callKey(ebiten.KeyG)
	g.callKey(ebiten.KeyR)
}
func (g *GameSim) callKey(k ebiten.Key) {
	switch k {
	case ebiten.KeyG:
		g.maingrid.Visible = !g.maingrid.Visible
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
		} else if g.SimSpeed < 20 {
			g.SimSpeed -= 2
		} else if g.SimSpeed < 40 {
			g.SimSpeed -= 4
		} else {
			g.SimSpeed -= 8
		}

	case ebiten.KeyArrowLeft:
		entity.CutoffDown()
	case ebiten.KeyArrowRight:
		entity.CutoffUp()
	case ebiten.KeyEscape:
		g.close = true
	}
}
