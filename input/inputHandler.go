package input

//STUB - need to rethink input handling

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// MKBinput
type MKBinput struct {
	cX, cY                 int
	kPressed, kJustPressed []ebiten.Key
}

// Keys is attempt at non-struct kb handler
var keysPressed []ebiten.Key = make([]ebiten.Key, 0, 16)
var keysJustPressed []ebiten.Key = make([]ebiten.Key, 0, 16)

// GetInKB is non-Handler Key Append.
func GetInKB() {
	keysPressed = inpututil.AppendPressedKeys(keysPressed[:0])
}

// GetJustPressedKeys does almost nothing at this point
func GetJustPressedKeys() []ebiten.Key {
	keysJustPressed = inpututil.AppendJustPressedKeys(keysJustPressed[:0])
	return keysJustPressed
}

// KeysOut (Debug use)
func KeysOut() *[]ebiten.Key {
	return &keysPressed
}

// *=======Handler=========
