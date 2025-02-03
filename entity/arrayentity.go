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
} /*can we do*/ /*multiple per line?*/
