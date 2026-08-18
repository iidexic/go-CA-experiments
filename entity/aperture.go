package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Aperture is the view-transform layer between a GridEntity and the screen.
// It owns the mapping from grid-pixel space to screen-pixel space and the
// pan/zoom navigation state. See docs/aperture.md for the design.
type Aperture struct {
	grid         *GridEntity
	gridTopLeft  image.Point     // top-left of visible region, in grid-pixel coords
	screenBounds image.Rectangle // where on screen the aperture draws
	zoom         float32         // screen pixels per grid pixel
	minZoom      float32
	maxZoom      float32

	dragging   bool
	dragAnchor image.Point // cursor position at drag start (screen coords)
	dragOrigin image.Point // gridTopLeft at drag start
}

// NewAperture returns an Aperture that draws the given grid within screenBounds.
// Initial zoom is 1 and the view is anchored at the grid's top-left corner.
func NewAperture(grid *GridEntity, screenBounds image.Rectangle) *Aperture {
	a := &Aperture{
		grid:         grid,
		screenBounds: screenBounds,
		zoom:         1,
		minZoom:      1,
		maxZoom:      32,
	}
	a.clampTopLeft()
	return a
}

// Pan shifts the visible region by (dxGrid, dyGrid) in grid-pixel coordinates.
func (a *Aperture) Pan(dxGrid, dyGrid int) {
	a.gridTopLeft.X += dxGrid
	a.gridTopLeft.Y += dyGrid
	a.clampTopLeft()
}

// Zoom applies a delta zoom step around a screen-coord anchor.
// The grid point currently under the anchor stays under the anchor after scaling.
func (a *Aperture) Zoom(delta float32, anchor image.Point) {
	a.SetZoom(a.zoom+delta, anchor)
}

// SetZoom applies an absolute zoom around a screen-coord anchor.
func (a *Aperture) SetZoom(z float32, anchor image.Point) {
	newZoom := z
	if newZoom < a.minZoom {
		newZoom = a.minZoom
	}
	if newZoom > a.maxZoom {
		newZoom = a.maxZoom
	}
	if newZoom == a.zoom {
		return
	}
	gx, gy, _ := a.ScreenToGrid(anchor.X, anchor.Y)
	a.zoom = newZoom
	// Re-anchor gridTopLeft so (gx, gy) maps back under the same screen anchor.
	a.gridTopLeft.X = gx - int(float32(anchor.X-a.screenBounds.Min.X)/a.zoom)
	a.gridTopLeft.Y = gy - int(float32(anchor.Y-a.screenBounds.Min.Y)/a.zoom)
	a.clampTopLeft()
}

// BeginPan records the drag start state. Called on right-button just-pressed.
func (a *Aperture) BeginPan(cursor image.Point) {
	a.dragging = true
	a.dragAnchor = cursor
	a.dragOrigin = a.gridTopLeft
}

// UpdatePan updates gridTopLeft from cursor delta relative to the drag start.
func (a *Aperture) UpdatePan(cursor image.Point) {
	if !a.dragging {
		return
	}
	dxScreen := cursor.X - a.dragAnchor.X
	dyScreen := cursor.Y - a.dragAnchor.Y
	a.gridTopLeft.X = a.dragOrigin.X - int(float32(dxScreen)/a.zoom)
	a.gridTopLeft.Y = a.dragOrigin.Y - int(float32(dyScreen)/a.zoom)
	a.clampTopLeft()
}

// EndPan clears the drag state. Called on right-button release.
func (a *Aperture) EndPan() { a.dragging = false }

// ScreenToGrid converts a screen coord to a grid coord. The third return
// reports whether the point falls within the aperture's screen bounds AND the grid.
func (a *Aperture) ScreenToGrid(sx, sy int) (int, int, bool) {
	gx := a.gridTopLeft.X + int(float32(sx-a.screenBounds.Min.X)/a.zoom)
	gy := a.gridTopLeft.Y + int(float32(sy-a.screenBounds.Min.Y)/a.zoom)
	inside := image.Pt(sx, sy).In(a.screenBounds) &&
		gx >= 0 && gx < int(a.grid.X) && gy >= 0 && gy < int(a.grid.Y)
	return gx, gy, inside
}

// GridToScreen converts a grid coord to a screen coord.
func (a *Aperture) GridToScreen(gx, gy int) (int, int) {
	sx := a.screenBounds.Min.X + int(float32(gx-a.gridTopLeft.X)*a.zoom)
	sy := a.screenBounds.Min.Y + int(float32(gy-a.gridTopLeft.Y)*a.zoom)
	return sx, sy
}

// Draw blits the visible portion of the grid image to the screen at the
// current zoom and screen position.
func (a *Aperture) Draw(screen *ebiten.Image) {
	visW := int(float32(a.screenBounds.Dx()) / a.zoom)
	visH := int(float32(a.screenBounds.Dy()) / a.zoom)
	if visW <= 0 || visH <= 0 {
		return
	}
	src := image.Rect(a.gridTopLeft.X, a.gridTopLeft.Y,
		a.gridTopLeft.X+visW, a.gridTopLeft.Y+visH)
	src = src.Intersect(image.Rect(0, 0, int(a.grid.X), int(a.grid.Y)))
	if src.Empty() {
		return
	}
	sub, ok := a.grid.Img.SubImage(src).(*ebiten.Image)
	if !ok {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(a.zoom), float64(a.zoom))
	// Offset by how much of the requested source was clipped by grid bounds,
	// so the visible region still lands at screenBounds.Min.
	offsetX := float64(a.screenBounds.Min.X) + float64(src.Min.X-a.gridTopLeft.X)*float64(a.zoom)
	offsetY := float64(a.screenBounds.Min.Y) + float64(src.Min.Y-a.gridTopLeft.Y)*float64(a.zoom)
	op.GeoM.Translate(offsetX, offsetY)
	screen.DrawImage(sub, op)
}

// clampTopLeft keeps gridTopLeft within [0, grid.size - visibleSize], or 0 if
// the grid is smaller than the visible region on that axis.
func (a *Aperture) clampTopLeft() {
	visW := int(float32(a.screenBounds.Dx()) / a.zoom)
	visH := int(float32(a.screenBounds.Dy()) / a.zoom)
	maxX := int(a.grid.X) - visW
	maxY := int(a.grid.Y) - visH
	if maxX < 0 {
		maxX = 0
	}
	if maxY < 0 {
		maxY = 0
	}
	if a.gridTopLeft.X < 0 {
		a.gridTopLeft.X = 0
	} else if a.gridTopLeft.X > maxX {
		a.gridTopLeft.X = maxX
	}
	if a.gridTopLeft.Y < 0 {
		a.gridTopLeft.Y = 0
	} else if a.gridTopLeft.Y > maxY {
		a.gridTopLeft.Y = maxY
	}
}
