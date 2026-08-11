package input

import "github.com/hajimehoshi/ebiten/v2"

type Binding interface {
	Exec()
}
type KeyContainer interface {
	IsUsed(ebiten.Key) bool
}

type Momentary interface {
	Triggered() bool
}

type Sustained interface {
	Momentary
	State() bool
}

type Latching interface {
	Binding
	Reset()
	state() bool
}
