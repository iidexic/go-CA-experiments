package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// type ezkb struct{}

// EZmouse handles cursor and mouse buttons
type EZmouse struct {
	cX, cY                int
	lmb, rmb, mmb, m4, m5 int
	hl                    ebiten.Image
}

// MBState refreshes state of mouse buttons
func (m *EZmouse) MBState() {
	// Duration in tick. 1 on mousedown
	m.lmb = inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft)
	m.rmb = inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight)
	m.mmb = inpututil.MouseButtonPressDuration(ebiten.MouseButtonMiddle)
	m.m4 = inpututil.MouseButtonPressDuration(ebiten.MouseButton3)
	m.m5 = inpututil.MouseButtonPressDuration(ebiten.MouseButton4)
}

var m EZmouse

// Mouse returns EZmouse obj
func Mouse() *EZmouse {
	return &m
}

// XY returns last-recorded coordinates of mouse cursor
func (m *EZmouse) XY() (int, int) {
	return m.cX, m.cY
}

// CursOn intakes a list of rectangles indicating
func (m *EZmouse) CursOn(bounds []int) int {
	m.cX, m.cY = ebiten.CursorPosition()
	if bounds[0] <= m.cX && m.cX <= bounds[2] && bounds[1] <= m.cY && m.cY <= bounds[3] {
		return 1
	}
	return 0
}
