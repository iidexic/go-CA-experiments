package entity

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// =======================
const (
	simShiftMod int = iota
	simLightVSDark
	simFloorCycle
)
const halfL int = 3 * 0xff / 2

// BaseEntity default entity type/debug entity
type BaseEntity struct {
	img        *ebiten.Image
	r          image.Rectangle
	op         *ebiten.DrawImageOptions
	set, drawn bool
}

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img            *ebiten.Image
	X, Y, Area     uint
	Pixels, PixLum []byte
	Op             ebiten.DrawImageOptions
	Set, Draw      bool
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity { // ? no ptr ok on GridEntity????
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{
		Img: ebiten.NewImage(width, height),
		Op:  ebiten.DrawImageOptions{},
		X:   uint(width), Y: uint(height), Area: uint(width * height),
		Set:    true,
		Pixels: make([]byte, width*height*4),
	}
	grid.Op.GeoM.Translate(float64((gWidth-width)/2), float64((gHeight-height)/2))
	grid.Img.Fill(color.RGBA{R: 155, G: 155, B: 165, A: 255})
	return &grid
}

// wrap val to remain within limit
func wrap(val, limit int) int {
	return ((val % limit) + limit) % limit
}

// TestSimulate is a testing simulation logic step
// will crash if grid.Pixel is not initialized.
func (grid *GridEntity) TestSimulate(shift, modifier int) {
	gX := int(grid.X)
	gY := int(grid.Y)
	testSimType := simShiftMod
	imax := len(grid.Pixels)
	switch testSimType {
	case simShiftMod:
		for i := range grid.Pixels {
			grid.Pixels[i] = grid.Pixels[((i+shift)<<modifier)%imax]
		}
	case simLightVSDark:
		//[Light Luminance vs Dark Luminance]
		//*= Further from 50% Gray = higher chance of victory
		//*= Simplest method, Check up and left for least amt of change to 
		for i := 0; i < len(grid.Pixels); i += 4 {
			csum := int(grid.Pixels[i]) + int(grid.Pixels[i+1]) + int(grid.Pixels[i+2])
			isLight:=csum>382
			if isLight{
			
			up := i - gX
			down := i + gX
			left := i - 4
			right := i + 4
			//>---(WRAP)---
			//* Wrap: Wrap top-bottom and also left-right
			if up < 0 {
				up = wrap(up, imax)
			} else if down >= imax {
				down = wrap(down, imax)
			}
			if left%gX == gX-1 {
				left += gX
			} else if right%gX == gX {
				right -= gX
			}
			//>---(!WRAP)---
			//up-down-left-right color sums
			csumU := grid.Pixels[up] + grid.Pixels[up+1] + grid.Pixels[up+2]
			csumD := grid.Pixels[down] + grid.Pixels[down+1] + grid.Pixels[down+2]
			csumL := grid.Pixels[left] + grid.Pixels[left+1] + grid.Pixels[left+2]
			csumR := grid.Pixels[right] + grid.Pixels[right+1] + grid.Pixels[right+2]

		}
	case simFloorCycle:

	}
}

func (grid *GridEntity) simShiftModulo(shift, modifier int) {

}

// ======================================================
// ======================================================

/*//** unused

//* this was sudocode for wrap in comparisons
//~	//first perform initial :
//~		u = i-wline, d = i+wline, l = i-1, r = i+1
//~	//then do checks
//~		if i < wline or i (is on last line) >=len-wline (i think) //VERTICAL WRAP
//~			then do mod +len mod
//~	//This made harder by the fact that we have 4val per pixel
//~ // Maybe we can do:
//~		if i%len<4 then add wlen to i (if left wrap add a row)
//~		if i%len>(len-4) then subtract wlen from i
//~	//should be good. Checking for each per pix

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
