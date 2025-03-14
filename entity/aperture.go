package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// called in NewAperture
func initCellMatrix(width, height int) *cellMatrix {
	cm := cellMatrix{
		Area: width * height,
		Px:   make([]byte, width*height*4),
	}
	return &cm
}

// --------------------------------------------------
// UNUSED - well the whold file is.
// ruleset/syscomponents intended pieces of new system, alongside aperture.
// keep all parts separate and modular. ruleset to define buckets and logic
// syscomponents struct originally intended to link pieces within the system
type ruleset interface {
	set(syscomponents)
	step()
}
type syscomponents struct {
}

// --------------------------------------------------
type cellMatrix struct {
	Area int
	Px   []byte
}

// Aperture is the window to view the actual cell grid through.
// allows zooming/panning
type Aperture struct {
	CM          *cellMatrix
	frame, view *ebiten.Image
	gm          ebiten.GeoM
	w, h        int
	zoom        byte
	area        []byte //x0,y0,x1,y1. like a rect but I don't need the other shit
	Px          []byte
	frameSizer  sizer //is there any point to this
}
type sizer func(int, int, byte) []int

// NewAperture returns a pointer to a properly initialized Aperture object
func NewAperture(width, height int, flags byte) *Aperture {
	ap := Aperture{
		w: width,
		h: height, frameSizer: simplesizer,
		CM: initCellMatrix(width, height),
	}
	ap.frameSizer(width, height, 0)
	return &ap
}

// ? There is a better way to do this for current use case
func simplesizer(w, h int, fl byte) []int {
	//if fl == 0 {
	return []int{w / 4, h / 9, (w * 3) / 4, (h * 8) / 9}
	//}

}
