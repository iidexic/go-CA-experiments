package gfx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gamedata"
)

// WindowConfig holds window and game sizes.
// Ebiten supports one window per process, so MainWindow is the only instance.
type WindowConfig struct {
	winPos    pt
	winSize   pt
	gameSize  pt
	gamescale float32
	title     string
	monitor   *ebiten.MonitorType
}

var mainWindow *WindowConfig

func InitWindow() {
	if mainWindow == nil {
		mainWindow = &WindowConfig{
			winSize:  gamedata.WindowSize(),
			gameSize: gamedata.GameSize(),
			monitor:  ebiten.Monitor(),
			title:    gamedata.WindowName(),
		}
		mainWindow.SetScale(float32(gamedata.WindowScale()))
		ApplyWindowConfig(mainWindow)
	}
}

func MainWindow() *WindowConfig {
	InitWindow()
	return mainWindow
}

func (d *WindowConfig) LastSizes() ([2]int, [2]int) {
	return d.winSize, d.gameSize
}

func (d *WindowConfig) SetScale(scale float32) {
	d.gamescale = scale
	if d.winSize[0] > 0 && d.winSize[1] > 0 {
		d.updateGameSize()
	}
}

func (d *WindowConfig) updateGameSize() {
	d.gameSize[0] = int(float32(d.winSize[0]) / d.gamescale)
	d.gameSize[1] = int(float32(d.winSize[1]) / d.gamescale)
}

func (d *WindowConfig) isInitialized() bool {
	ws := d.winSize[0] > 0 && d.winSize[1] > 0
	gs := d.gameSize[0] > 0 && d.gameSize[1] > 0
	ss := d.gamescale > 0
	return ws && gs && ss
}

func ApplyWindowConfig(d *WindowConfig) {
	if d.isInitialized() {
		ebiten.SetWindowSize(d.winSize[0], d.winSize[1])
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
		ebiten.SetWindowTitle(d.title)
		ebiten.SetWindowPosition(d.winPos[0], d.winPos[1])
	}
}

type pt [2]int
