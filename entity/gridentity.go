package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img                   *ebiten.Image
	X, Y                  uint
	modAdd, modMult, Area int
	Pixels, PixLum        []byte
	Op                    ebiten.DrawImageOptions
	Set, Draw             bool
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity { // ? no ptr ok on GridEntity????
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{
		Img: ebiten.NewImage(width, height),
		Op:  ebiten.DrawImageOptions{},
		X:   uint(width), Y: uint(height), Area: width * height,
		Set:     true,
		Pixels:  make([]byte, width*height*4),
		modAdd:  0,
		modMult: 1,
	}
	grid.Op.GeoM.Translate(float64((gWidth-width)/2), float64((gHeight-height)/2))
	grid.Img.Fill(color.RGBA{R: 155, G: 155, B: 165, A: 255})
	return &grid
}

// SetMod sets modulation values modAdd and modMult
func (grid *GridEntity) SetMod(modAdd, modMult int) {
	grid.modAdd = modAdd
	grid.modMult = modMult
}

// SimstepLVSD performs one cycle/screen of checks and updates
// for the center-distance intensity comparison sim ("Light VS Dark")
func (grid *GridEntity) SimstepLVSD() {
	// Fixed by doing -3 to sent len value in functionMod.
	//TODO: Now all combinations of modifiers (add/mult) will have grid go to white. Is that how it shold work?...,
	for i := 0; i < len(grid.Pixels); i += 4 {
		//? any benefit to using i++ in range area *4 for index? yes prob
		toIndex := functionMod(i, grid.modAdd, grid.modMult, len(grid.Pixels)) //FIXME
		/*//! Function mod not keeping in bounds (most likely)
		LIMIT = grid.width*grid.height ~ 518400 for now.

		*/
		grid.pxGoToward(i, grid.Pixels[toIndex:toIndex+3])
	}
}

func (grid *GridEntity) pxGoToward(indexR int, toPx []byte) {
	for i := range 3 {
		if grid.Pixels[indexR+i] > toPx[i] {
			grid.Pixels[indexR+i] -= (grid.Pixels[indexR+i] - toPx[i]) / 2
		} else {
			grid.Pixels[indexR+i] += (toPx[i] - grid.Pixels[indexR+i]) / 2
		}
	}
}
func (grid *GridEntity) pxReplace(indexR int, new []byte) {
	grid.Pixels[indexR] = new[0]
	grid.Pixels[indexR+1] = new[1]
	grid.Pixels[indexR+2] = new[2]
}

// pxTransplant overwrites 1px (3 indices) in grid.Pixels, starting at index
// write uses values at indices R,G,B without changes made during the function call
func (grid *GridEntity) pxTransplant(index int, R, G, B int) {
	var nR byte = grid.Pixels[R]
	var nG byte = grid.Pixels[G]
	var nB byte = grid.Pixels[B]
	grid.Pixels[index] = nR
	grid.Pixels[index+1] = nG
	grid.Pixels[index+2] = nB
}

func functionMod(start, add, mult int, limit int) int {
	return wrap(((start + add) * mult), limit-3)
}
