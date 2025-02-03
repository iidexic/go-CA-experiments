package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img                   *ebiten.Image
	X, Y                  uint
	modAdd, modMult, Area int
	Pixels, PixLum        []byte
	Op                    ebiten.DrawImageOptions
	Set, Draw             bool // probably not in use
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity {
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
	grid.Img.Fill(gfx.PaletteGP[gfx.Dark])
	return &grid
}

// SetMod sets modulation values modAdd and modMult
func (grid *GridEntity) SetMod(modAdd, modMult int) {
	grid.modAdd = modAdd
	grid.modMult = modMult
}

// SimstepLVSD performs one cycle/screen of checks and updates
// for the center-distance intensity comparison sim ("Light VS Dark")
func (grid *GridEntity) SimstepLVSD(pixLock bool) {

	for i := 0; i < len(grid.Pixels); i += 4 {
		//newR := shiftMod(i, grid.modAdd, grid.modMult, len(grid.Pixels))

		first, last := wrapRange(i, 3, grid.modAdd, grid.modMult, len(grid.Pixels))
		grid.pxGoToward(i, grid.Pixels[first:last])
	}

}

func shiftMod(start, add, mult int, limit int) int {
	return wrap(((start + add) * mult), limit-3)
}
func wrapRange(start, len, add, mult int, limit int) (int, int) {
	ishift := wrap((start+add)*mult, limit)
	endshift := ishift + (len - 1)
	if endshift < limit {
		return ishift, endshift
	} //* if not then (for now) subtract length
	return ishift - (len - 1), endshift - (len - 1)
}

// SimstepValueShift is home of current more simplistic sim model after move on to full lvsd
// (keep pixlock for now) pixlock false allows offset color value averaging (i.e. blue vs red channel).
func (grid *GridEntity) SimstepValueShift(pixLock bool) {
	if !pixLock {
		for i := 0; i < len(grid.Pixels); i += 4 {
			//newR := shiftMod(i, grid.modAdd, grid.modMult, len(grid.Pixels))
			first, last := wrapRange(i, 3, grid.modAdd, grid.modMult, len(grid.Pixels))
			grid.pxGoToward(i, grid.Pixels[first:last])
		}
	} else {
		for i := 0; i < grid.Area; i++ {
			iR := i * 4
			newPix := shiftMod(i, grid.modAdd, grid.modMult, grid.Area)
			newR := newPix * 4 //area to RGB, will always land on R val

			grid.pxGoToward(iR, grid.Pixels[newR:newR+3])
		}

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

//==Old, probably entirely unnecessary==
