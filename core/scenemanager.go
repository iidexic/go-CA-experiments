package core

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

var QuitGame error = fmt.Errorf("quit game")

type SceneManager struct {
	scenes map[string]ebiten.Game
}

var Manager *SceneManager = &SceneManager{scenes: make(map[string]ebiten.Game)}

func (m *SceneManager) AddScene(name string, scene ebiten.Game) {
}

func (m *SceneManager) LaunchScene(name string) error {
	if _, ok := m.scenes[name]; !ok {
		return fmt.Errorf("scene '%s' not found", name)
	}

	return nil
}

func (m *SceneManager) run(name string) error {

	if err := ebiten.RunGame(m.scenes[name]); err != nil {
		if err == QuitGame {
			return nil
		}
		return err
	}
	return nil
}
