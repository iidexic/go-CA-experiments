package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type ruleset interface {
	set(syscomponents)
	eval([]byte) func()
}
type drawable interface {
}
type syscomponents struct {
}

type pt = image.Point //currently unused

type cellMatrix struct {
	img   *ebiten.Image
	x, y  int
	Area  int
	Px    []byte
	zoom  byte
	rules ruleset
}

func initCellMatrix(width, height int) cellMatrix {
	cm := cellMatrix{
		Area: width * height,
		Px:   make([]byte, width*height*4),
	}
	return cm
}

type vobj struct {
	i  *ebiten.Image
	op ebiten.DrawImageOptions
	g  *ebiten.GeoM
}

func sizeVobj(x, y, w, h int) vobj {
	v := vobj{
		i:  ebiten.NewImage(w, h),
		op: ebiten.DrawImageOptions{},
	}
	v.g = &v.op.GeoM
	v.g.Translate(float64(x), float64(y))
	return v
}

// Aperture is the window to view the actual cell grid through.
// allows zooming/panning
type Aperture struct {
	CM           cellMatrix
	frame, world vobj
	overlay      vobj
	x, y, w, h   int
	zoom         byte
	Px           []byte
}

// NewAperture returns a pointer to a properly initialized Aperture object
func NewAperture(x, y, width, height int, flags byte) *Aperture {
	ap := Aperture{x: x, y: y, w: width, h: height}
	ap.positionScreenAspectR(width, height, 12)
	ap.frame = sizeVobj(0, 0, ap.w, ap.h)

	ap.CM = initCellMatrix(ap.w, ap.h)
	return &ap
}

// ZoomL in/out by amt. -> max single call is zooming half the 255 range.
func (a *Aperture) ZoomL(amt int8) {
	d := int(amt)
	z := int(a.zoom)
	sum := d + z
	if 0 <= sum && sum <= 255 {
		a.zoom = byte(sum)
	} else if 0 > sum {
		a.zoom = 0
	} else if 255 < sum {
		a.zoom = 255
	}
}

func (a *Aperture) positionScreenAspectR(w, h, border int) {
	if w > (h*4)/3 {
		a.h = h - (2 * border)
		a.w = 4 * a.h / 3
	} else {
		a.w = w - (2 * border)
		a.h = 3 * a.w / 4
	}
	//a.op.GeoM.Translate(float64(w-border-a.w), float64(border))
	//~ build out gfx.positioning to take this over
	//~ Settings:
	//~ (placementMethod ==> anchorRight)
	//~ (sizing ==> maxAspect), (padPxY, padPxX ==> 12)
}
