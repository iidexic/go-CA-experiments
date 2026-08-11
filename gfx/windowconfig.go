package gfx

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gamedata"
)

// WindowConfig holds window and game sizes, because ebiten is limited to one window, most situations will only require use of MainWindow()
type WindowConfig struct {
	id            string
	winPos        pt
	winSize       pt
	gameSize      pt
	gamescale     float32
	title         string
	monitor       *ebiten.MonitorType
	maintainScale bool
}

var gamewindows map[string]*WindowConfig = make(map[string]*WindowConfig)
var mainWindow *WindowConfig

func InitWindow() {
	if mainWindow == nil {
		mainWindow = &WindowConfig{
			winSize:       gamedata.WindowSize(),
			gameSize:      gamedata.GameSize(),
			monitor:       ebiten.Monitor(),
			id:            "main",
			title:         gamedata.WindowName(),
			maintainScale: true,
		}
		mainWindow.SetScale(float32(gamedata.WindowScale()))
		gamewindows["main"] = mainWindow
		ApplyWindowConfig(mainWindow)
	}
}

func MainWindow() *WindowConfig {
	InitWindow()
	return mainWindow
}

func stringidentifier(id string) string { return strings.ToLower(strings.ReplaceAll(id, " ", "-")) }

// NewWindow creates a new window. Unless doing weird stuff this will probably never be used
func NewWindow(width, height, scale int, id string, title string) *WindowConfig {
	d := WindowConfig{
		winSize: pt{width, height},
		monitor: ebiten.Monitor(),
		title:   "CellularEmulator",
	}
	if d.id == "" {
		if _, ok := gamewindows[stringidentifier(d.title)]; ok {
		}
	}
	monX, monY := d.monitor.Size()
	//TODO: Change Back from Debug
	d.winPos = startpos(pt{monX, monY}, d.winSize)
	d.SetScale(float32(scale))
	return &d
}

func startpos(monsize pt, winsize pt) pt {
	return pt{monsize[0]/2 - winsize[0]/2, monsize[1]/2 - winsize[1]/2}
}

func (d *WindowConfig) LayoutGameSize(winWidth, winHeight int) (int, int) {
	if !(winWidth == d.winSize[0] && winHeight == d.winSize[1]) {
		d.winSize[0] = winWidth
		d.winSize[1] = winHeight
		if d.maintainScale {
			d.updateGameSize()
		}
	}
	return d.gameSize[0], d.gameSize[1]
}

func (d *WindowConfig) LastSizes() ([2]int, [2]int) {
	return d.winSize, d.gameSize
}

func (d *WindowConfig) SetWindowPosition(x, y int) {
	d.winPos = pt{x, y}
}
func (d *WindowConfig) SetWindowSizePx(width, height int) {
	d.winSize = pt{width, height}
	if d.gamescale == 0 {
		d.gamescale = 2
	}
	d.updateGameSize()
}

func (d *WindowConfig) SetScale(scale float32) {
	d.gamescale = scale
	if d.winSize[0] > 0 && d.winSize[1] > 0 {
		d.updateGameSize()
	}
}

// careful with these. Commenting out winsize for now
func (d *WindowConfig) updateGameSize() {
	d.gameSize[0] = int(float32(d.winSize[0]) / d.gamescale)
	d.gameSize[1] = int(float32(d.winSize[1]) / d.gamescale)
}

// func (d *Display) updateWinSize() {
// 	d.winSize[0] = int(float32(d.gameSize[0]) * d.gamescale)
// 	d.winSize[1] = int(float32(d.gameSize[1]) * d.gamescale)
// }

func (d *WindowConfig) isInitialized() bool {
	ws := d.winSize[0] > 0 && d.winSize[1] > 0
	gs := d.gameSize[0] > 0 && d.gameSize[1] > 0
	ss := d.gamescale > 0
	return ws && gs && ss && d.id != ""
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

// type xy struct {
// 	x, y int
// }
