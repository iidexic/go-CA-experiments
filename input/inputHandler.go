package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//gotta rethink this one

var x ebiten.Key

func inputInit() {
	//? Slice size allocation arbitrary. Is there a better way?
	padID := make([]ebiten.GamepadID, 6)
	ebiten.AppendGamepadIDs(padID)
	inKey := make([]ebiten.Key, 24)
	inPad := make([]ebiten.GamepadButton, 60)
	inpututil.AppendPressedGamepadButtons()
	inpututil.AppendPressedStandardGamepadButtons()
	ebiten.GamepadAxis()
}
func getState() []ebiten.Key {

	inpututil.AppendPressedKeys(inKey)

}
