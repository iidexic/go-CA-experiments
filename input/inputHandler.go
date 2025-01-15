package input

//STUB - need to rethink input handling

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// !===Temp Input Handler==================================
var keyAssignments []func()

//!=======================================================
//==Intended Kb Handler - Waiting on help=================

// HandlerKBM Manages KB + Mouse Inputs
type handlerKBM struct {
	kbDown, kbUp []ebiten.Key
	msBtn        []ebiten.MouseButton
}

// *temp-----
var hkey handlerKBM = handlerKBM{
	kbDown: make([]ebiten.Key, 159),
	kbUp:   make([]ebiten.Key, 159),
}

func (h *handlerKBM) GetState() {
	//h.kbDown = make([]ebiten.Key, 6)
	//h.kbUp = make([]ebiten.Key, 6)
	inpututil.AppendPressedKeys(h.kbDown)
	inpututil.AppendJustReleasedKeys(h.kbUp)

	// for future brain remembering: we are wiping every time
	//totally thought this was gonna work
	// for i := 0; h.kbDown[i] != 0; i++ {
	for i, v := range h.kbDown {
		if v > 0 {
			fmt.Printf("KeyDown: index=%d, key# %d key:%s\n", i, v, ebiten.Key(v).String())
		}
	}
}

// UpdateKeys is test func to trigger the internal global handler object (hkey) GetState method
func UpdateKeys() {
	hkey.GetState()
}

//maybe someday
/*
	padID := make([]ebiten.GamepadID, 6)
	ebiten.AppendGamepadIDs(padID)

	//inPad := make([]ebiten.GamepadButton, 60)
	inpututil.AppendPressedGamepadButtons()
	inpututil.AppendPressedStandardGamepadButtons()
*/
