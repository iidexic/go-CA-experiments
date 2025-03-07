package entity

import "github.com/hajimehoshi/ebiten/v2"

// ArrayEntity testing an entity like grid but with per-pixel subimages or images
type ArrayEntity struct {
	Canvas     *ebiten.Image
	CanvasOpts *ebiten.DrawImageOptions
	CanvasGeoM *ebiten.GeoM
	//Below are slices for Cell images
	Cells               *[]ebiten.Image
	Opts                *[]ebiten.DrawImageOptions
	GeoMs               *[]ebiten.GeoM
	X, Y, Width, Height int
	CellColor           []byte
}

// MakeArrayDefault initializes cell array
func MakeArrayDefault(gWidth, gHeight int) *ArrayEntity {
	// square based on height, offset to one side
	ar := &ArrayEntity{}

	return ar
}

// SetMod sets modulation values modAdd and modMult
func (grid *GridEntity) SetMod(modAdd, modMult int) {
	grid.modAdd = modAdd
	grid.modMult = modMult
}

//=================================================================
/*
Insular, initial method of getting some interaction between pixels
Code below either all applies to SimstepValueShift, or is entirely unused.
*/

// SimstepValueShift is home of current more simplistic sim model after move on to full lvsd
// (keep pixlock for now) pixlock false allows offset color value averaging (i.e. blue vs red channel).
func (grid *GridEntity) SimstepValueShift(pixLock bool) {
	if !pixLock {
		for i := 0; i < len(grid.Px); i += 4 {
			//newR := shiftMod(i, grid.modAdd, grid.modMult, len(grid.Px))
			first, last := wrapRange(i, 3, grid.modAdd, grid.modMult, len(grid.Px))
			grid.pxGoToward(i, grid.Px[first:last])
		}
	} else {
		for i := 0; i < grid.Area; i++ {
			iR := i * 4
			newPix := shiftMod(i, grid.modAdd, grid.modMult, grid.Area)
			newR := newPix * 4 //area to RGB, will always land on R val

			grid.pxGoToward(iR, grid.Px[newR:newR+3])
		}

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

func (grid *GridEntity) pxGoToward(indexR int, toPx []byte) {
	for i := range 3 {
		if grid.Px[indexR+i] > toPx[i] {
			grid.Px[indexR+i] -= (grid.Px[indexR+i] - toPx[i]) / 2
		} else {
			grid.Px[indexR+i] += (toPx[i] - grid.Px[indexR+i]) / 2
		}
	}
}
