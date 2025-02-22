package input

//STUB - need to rethink input handling

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// maybe just make a struct for these
var keysPressed []ebiten.Key = make([]ebiten.Key, 0, 16)
var keysJustPressed []ebiten.Key = make([]ebiten.Key, 0, 16)

// GetInKB is DEBUG Key List
func GetInKB() {
	keysPressed = inpututil.AppendPressedKeys(keysPressed[:0])
}

// GetJustPressedKeys just middlemans right now
func GetJustPressedKeys() []ebiten.Key {
	keysJustPressed = inpututil.AppendJustPressedKeys(keysJustPressed[:0])
	return keysJustPressed
}

// KeysOut (Debug use)
func KeysOut() *[]ebiten.Key {
	return &keysPressed
}

// KeyHandler
type keyhandler struct {
	keys  []ebiten.Key
	binds []int
}
