package input

//STUB - need to rethink input handling

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Keys is attempt at non-struct kb handler
var keysPressed []ebiten.Key = make([]ebiten.Key, 0, 16)

// GetInKB is non-Handler Key Append.
func GetInKB() {
	keysPressed = inpututil.AppendPressedKeys(keysPressed[:0])
}

// KeysOut (Debug use)
func KeysOut() *[]ebiten.Key {
	return &keysPressed
}

// *=======Handler=========
