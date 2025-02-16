package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Entity denotes any game entity
type Entity interface {
	GetImg() *ebiten.Image
	GetOpt() *ebiten.DrawImageOptions
	GetGeom() *ebiten.GeoM
}

// =======================

// =========Movement Testing========
func (b *BaseEntity) testMovements() {
	// dunno
}

// =================================

// BaseEntity default entity type/debug entity
type BaseEntity struct {
	w, h       int
	Img        *ebiten.Image
	R          image.Rectangle
	Opt        *ebiten.DrawImageOptions
	GeoM       *ebiten.GeoM
	Set, Drawn bool
}

// GetImg returns entity image
func (b *BaseEntity) GetImg() *ebiten.Image {
	return b.Img
}

// GetOpt returns entity image draw options
func (b *BaseEntity) GetOpt() *ebiten.DrawImageOptions {
	return b.Opt
}

// GetGeom returns entity's drawoptions GeoM
func (b *BaseEntity) GetGeom() *ebiten.GeoM {
	return &b.Opt.GeoM
}

// NewBaseEntity does its thing
func NewBaseEntity(w, h int) *BaseEntity {
	b := BaseEntity{
		Img: ebiten.NewImage(w, h),
		w:   w,
		h:   h,
		Opt: &ebiten.DrawImageOptions{},
	}
	b.GeoM = &b.Opt.GeoM
	return &b
}

// wrap val to remain within limit
func wrap(val, limit int) int {
	return ((val % limit) + limit) % limit
}

// TODO: Messy implementation. Fix later if desired.
func sidewrap(index, move, width int) int {
	//[Treat as if wrapping around 1 single row]
	irownum := index / width
	iWrap := index % width
	wrapMoved := ((iWrap+(move%width))%width + width) % width
	return wrapMoved + (irownum * width)

}

func bavg(b ...byte) (avg byte) {
	var rval int = 0
	for i := range b {
		rval += int(b[i])
	}
	return byte(rval / len(b))
}
func erch(e error) {
	if e != nil {
		panic(e)
	}
}
