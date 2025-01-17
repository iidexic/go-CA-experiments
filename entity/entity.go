package entity

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// =======================
const (
	simShiftMod int = iota
	simNearbyChance
	simFloorCycle
)

// BaseEntity default entity type/debug entity
type BaseEntity struct {
	img        *ebiten.Image
	r          image.Rectangle
	op         *ebiten.DrawImageOptions
	set, drawn bool
}

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img        *ebiten.Image
	X, Y, Area uint
	Pixels     []byte
	Op         ebiten.DrawImageOptions
	Set, Draw  bool
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity { // ? no ptr ok on GridEntity????
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{
		Img: ebiten.NewImage(width, height),
		Op:  ebiten.DrawImageOptions{},
		X:   uint(width), Y: uint(height), Area: uint(width * height),
		Set: true,
	}
	grid.Op.GeoM.Translate(float64((gWidth-width)/2), float64((gHeight-height)/2))
	grid.Img.Fill(color.RGBA{R: 155, G: 155, B: 165, A: 255})
	return &grid
}

// TestSimulate is a testing simulation logic step
// It isn't working because its out of range du
func (grid *GridEntity) TestSimulate(shift, modifier int) {
	testSimType := simShiftMod
	imax := len(grid.Pixels)
	switch testSimType {
	case simShiftMod:
		for i := range grid.Pixels {
			grid.Pixels[i] = grid.Pixels[((i+shift)<<modifier)%imax]
		}
	case simNearbyChance:

	case simFloorCycle:

	}
}

func (grid *GridEntity) simShiftModulo(shift, modifier int) {

}

// ======================================================
// ======================================================

/*//* unused

//first test sim code. not working because goes oob. also bitshift prob not the best anyway

func rightShiftWrap(val, nbits, limit) val {}

for i := range grid.Pixels {
		s := i >> 1
		lpix := len(grid.Pixels)
		if s >= lpix {
			s = s % lpix
		} else if s < 0 {

		}

		if i%2 == 0 {
			grid.Pixels[i] = grid.Pixels[i>>1]
		} else {
			grid.Pixels[i] = grid.Pixels[i<<1]
		}
	}
// -----------------------
// INTERFACES

// entityProperties holds values corresponding to game entity
type entityProperties interface {SetProperties()}

// Entity interface for all game entities.
type Entity interface {Emit()}

// -----------------------
// STRUCTS

// TextEntity for defining and drawing text
type TextEntity struct {textFaceSource *text.GoTextFaceSource}

// MapEntity planned for array of instances of entity with set locations
type MapEntity struct {}

// -----------------------
// FUNCTIONS
func borderFill(grid *GridEntity) {} //TODO
func resetGrid(grid *GridEntity) {} //TODO
*/
