package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type inputState struct {
}
type inputKeyboard struct {
	pressed, down, released []ebiten.Key
}
type inputMouse struct {
	pressed, down, rleased []ebiten.MouseButton
	wheel                  int
}
type inputPad struct {
}

var x ebiten.Key

func getState() []ebiten.Key {
	in := []ebiten.Key{}
	inpututil.AppendPressedKeys(in)
	pad := inpututil.AppendPressedGamepadButtons()

}
