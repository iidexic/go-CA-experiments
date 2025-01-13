package entity

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// =======================
// STRUCTS

// BaseEntity default entity type/debug entity
type BaseEntity struct {
	img        *ebiten.Image
	r          image.Rectangle
	op         *ebiten.DrawImageOptions
	set, drawn bool
}

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img       *ebiten.Image
	r         image.Rectangle //intended to house width/height. Not needed atm
	Op        ebiten.DrawImageOptions
	Set, Draw bool
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) GridEntity { // ? no ptr ok on GridEntity????
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{
		Img: ebiten.NewImage(width, height),
		Op:  ebiten.DrawImageOptions{},
		Set: true,
	}
	grid.Op.GeoM.Translate(float64((gWidth-width)/2), float64((gHeight-height)/2))
	grid.Img.Fill(color.RGBA{R: 155, G: 155, B: 165, A: 255})
	return grid
}

// ======================================================
// ======================================================
// * unused
/*
// -----------------------
// INTERFACES

// entityProperties holds values corresponding to game entity
type entityProperties interface {SetProperties()}

// Entity interface for all game entities.
type Entity interface {Emit()}

// -----------------------
// Structs

// TextEntity for defining and drawing text
type TextEntity struct {textFaceSource *text.GoTextFaceSource}

// MapEntity planned for array of instances of entity with set locations
type MapEntity struct {}

// -----------------------
// Functions
func borderFill(grid *GridEntity) {} //TODO
func resetGrid(grid *GridEntity) {} //TODO
*/
