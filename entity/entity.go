package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Entity denotes any game entity
type Entity interface {
	GetImg() *ebiten.Image
	GetOpt() *ebiten.DrawImageOptions
	XY() (int, int)
}

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

func sidewrap(index, move, width int) int {
	//[Treat as if wrapping around 1 single row, also this code could use cleanup]
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

// absolute center difference byte average.
func acdbavg(b ...byte) byte { //TODO: REWRITE - Can't tell if even functioning as intended
	tot := 0
	for _, v := range b {
		tot += int(v) - 127
	}
	a := tot / len(b)
	if a < 0 {
		a = -a
	}
	return byte(a)
}

// byte slice limit add (probably should combine with bslsub)
func bsladd(b []byte, add byte) {
	for i := range b {
		if add > 255-b[i] {
			b[i] = 255
		} else {
			b[i] += add
		}
	}
}

func bslsub(b []byte, subtract byte) {
	for i := range b {
		if b[i] < subtract {
			b[i] = 0
		} else {
			b[i] -= subtract
		}
	}
}

// Erch check error do checking
func Erch(e error) {
	if e != nil {
		panic(e)
	}
}
