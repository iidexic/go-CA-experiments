package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// =======================
// INTERFACES
// =======================

// entityProperties holds values corresponding to game entity
type entityProperties interface {
	SetProperties()
}

// entityState holds information on an entity state
// as well as information to use in transitioning to other states
type entityState interface {
	trigger()
	update()
}
type stateManager interface {
	Define()
}

// REVIEW might want to actually make the Entity just separate from  entityProperties, entityState.

// Entity interface for all game entities.
type Entity interface {
	Defaults()
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
	op        *ebiten.DrawImageOptions
	set, draw bool
}

// MapEntity planned for array of instances of entity with set locations
type MapEntity struct {
}

// Defaults for entity type
func (grid GridEntity) Defaults() {
	width := (3 * pixWidth) / 4
	height := (3 * pixHeight) / 4
	grid.img = ebiten.NewImage(width, height)
	grid.op = &ebiten.DrawImageOptions{}
	grid.op.GeoM.Translate(float64((pixWidth-width)/2), float64((pixHeight-height)/2))
	grid.img.Fill(color.Gray{})
	grid.set = true
}

// ?-------------------byval ok?
func makeGridDefault() GridEntity {
	grid := GridEntity{
		img: ebiten.NewImage((3*pixWidth)/4, (3*pixHeight)/4),
		op:  &ebiten.DrawImageOptions{},
		set: true,
	}
	width := (3 * pixWidth) / 4
	height := (3 * pixHeight) / 4

	grid.op.GeoM.Translate(float64((pixWidth-width)/2), float64((pixHeight-height)/2))
	grid.img.Fill(color.Gray{})
	return grid
}

// SetProperties of BaseEntity object
func (e *BaseEntity) SetProperties(i int) {
	_ = i

}
