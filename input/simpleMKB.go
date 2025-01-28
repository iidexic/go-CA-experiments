package input

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
	type MouseInteract interface {
		Hover()
		IsOver()
	}
*/
// ezmouse handles mouse input
type ezmouse struct {
	cX, cY                int
	lmb, rmb, mmb, m4, m5 int
}

func (m *ezmouse) MBState() {
	// Duration in tick. 1 on mousedown
	m.lmb = inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft)
	m.rmb = inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight)
	m.mmb = inpututil.MouseButtonPressDuration(ebiten.MouseButtonMiddle)
	m.m4 = inpututil.MouseButtonPressDuration(ebiten.MouseButton3)
	m.m5 = inpututil.MouseButtonPressDuration(ebiten.MouseButton4)
}

// CursOn intakes a list of rectangles indicating
func (m *ezmouse) CursOn(rects []image.Rectangle) int {
	cX, cY := ebiten.CursorPosition()
	for i, v := range rects {
		if v.In(image.Rect(cX, cY, cX, cY)) {
			return i
		}

	}
	return 0
}
