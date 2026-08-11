package input

import "github.com/hajimehoshi/ebiten/v2"

type InputHandler struct {
	*KeyMap
	Keymaps map[string]KeyBindManager
}

var mainhandler = InputHandler{
	KeyMap:  &KeyMap{binds: make(map[ebiten.Key]KeyBind)},
	Keymaps: make(map[string]KeyBindManager),
}

func NewKeymap(name string) *KeyMap {
	km := KeyMap{binds: make(map[ebiten.Key]KeyBind)}
	mainhandler.Keymaps[name] = &km
	return &km
}

// Update all input
func (ih *InputHandler) Update() {
	go ih.KeyMap.Update()
	for _, km := range ih.Keymaps {
		go km.Update()
	}
}
