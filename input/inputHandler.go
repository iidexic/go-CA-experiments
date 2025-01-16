package input

//STUB - need to rethink input handling

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//==Intended Kb Handler - Waiting on help=================

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

// handlerKBM Manages KB + Mouse Inputs
type handlerKBM struct {
	kbDown, kbUp []ebiten.Key
	msBtn        []ebiten.MouseButton
}

// *temp-----
// InitMain is optional  init for  when handler is
// something, but anyway I am still not understanding this look at diff w-debug
func (h handlerKBM) InitMain() {
	h.kbDown = make([]ebiten.Key, 16)
	_ = h.kbDown
}

// GetState updates handler with held keys
func (h *handlerKBM) GetState() {
	inpututil.AppendPressedKeys(h.kbDown[:0])
	for i, v := range h.kbDown {
		fmt.Println("got here 2")
		fmt.Printf("KeyDown: index=%d, key# %d key:%s\n", i, v, ebiten.Key(v).String())
	}
}

// getHandler returns initialized inputHandler
func getHandler() handlerKBM {
	return handlerKBM{
		kbDown: make([]ebiten.Key, 0, 16),
		kbUp:   make([]ebiten.Key, 0, 16)}
}
