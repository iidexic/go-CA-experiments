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

// Entity interface for all game entities
type Entity interface {
	NewInstance()
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
	base         color.RGBA
	x, y, tx, ty float64
}

// SetProperties of BaseEntity object
func (e *BaseEntity) SetProperties(i int) {
	_ = i
}
