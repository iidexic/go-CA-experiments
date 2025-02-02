package entity

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// =======================
const (
	simShiftMod int = iota
	simLightVSDark
	simFloorCycle
)
const halfL int = 3 * 0xff / 2

// Entity denotes any game entity
type Entity interface {
	GetImg() *ebiten.Image
	GetOpt() *ebiten.DrawImageOptions
	GetGeom() *ebiten.GeoM
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

func moveBhalf(from, to byte) byte {
	return (to - from) / 2
}
func colorDistanceVS(vHi, vLo, center, rng int) bool {
	return vHi+vLo-center > rng // true == vHi win
}
func coloravg(b []byte) int {
	return (int(b[0]) + int(b[1]) + int(b[2])) / 3
}
func bavg(b []byte) int {
	var sum int = 0
	var i int
	for i = range b {
		sum += int(b[i])
	}
	return sum / (i + 1)
}

// TestSimulate removed 2-2-25
