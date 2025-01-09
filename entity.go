package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// =======================
// INTERFACES
// =======================

// entityProperties holds values corresponding to game entity
type entityProperties interface {
	SetProperties()
}

// Entity interface for all game entities.
type Entity interface {
	EmitEntity()
}

// =======================
// STRUCTS
// =======================

// BaseEntity default entity type/debug entity
type BaseEntity struct {
	img        *ebiten.Image
	r          image.Rectangle
	op         *ebiten.DrawImageOptions
	set, drawn bool
}

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	img       *ebiten.Image
	subs      []ebiten.Image
	r         image.Rectangle
	op        ebiten.DrawImageOptions
	set, draw bool
}

func (g *GridEntity) populateGrid(size, spacing int) func() gen

// TextEntity for defining and drawing text
type TextEntity struct {
	textFaceSource *text.GoTextFaceSource
}

// MapEntity planned for array of instances of entity with set locations
type MapEntity struct {
}

// ?-------------------byval ok?
func makeGridDefault() GridEntity {
	width := (3 * gameWidth) / 4
	height := (3 * gameHeight) / 4
	grid := GridEntity{
		img: ebiten.NewImage(width, height),
		op:  ebiten.DrawImageOptions{},
		set: true,
	}
	grid.op.GeoM.Translate(float64((gameWidth-width)/2), float64((gameHeight-height)/2))
	grid.img.Fill(color.RGBA{R: 155, G: 155, B: 165, A: 255})
	return grid
}
func borderFill(grid *GridEntity) {
	//TODO
}
func resetGrid(grid *GridEntity) {
	//TODO
}

// SetProperties of BaseEntity object
func (e *BaseEntity) SetProperties(i int) {
	_ = i

}
