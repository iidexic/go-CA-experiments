package core

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/util"
)

// GameSim struct - ebiten
type GameSim struct {
	maingrid                         *entity.GridEntity
	pal                              []color.RGBA
	gWidth, gHeight, pWidth, pHeight int
	SimSpeed                         int
	ticks                            uint16
	devFASTSTART, close              bool
}

// GameSimInit returns GameSim pointer for main sim scene with default settings
func GameSimInit(GameSimWidth, GameSimHeight int) *GameSim {
	g := &GameSim{
		SimSpeed: 1,
		gWidth:   GameSimWidth,
		gHeight:  GameSimHeight,
		pal:      gfx.PaletteGP,
	}
	g.maingrid = entity.MakeGridDefault(g.gWidth, g.gHeight)
	//==== TESTING STUFF ====
	g.devFASTSTART = true
	//=======================
	return g
}

// Update function
func (g *GameSim) Update() error {
	g.debugUpdate()
	g.ticks++
	if g.devFASTSTART {
		g.devFASTSTART = false
		g.fastInitializeDev()
	}
	if g.SimSpeed > 0 && g.isSimTick() {
		g.maingrid.SimstepLVSD()
		if g.maingrid.Debug {
			g.maingrid.Img.WritePixels(g.maingrid.ApplyDbgOverlay(0))
		} else {
			g.maingrid.Img.WritePixels(g.maingrid.Px)
		}
	}
	inputActions(g)
	if g.close {
		return ebiten.Termination
	}
	return nil
}

// Draw screen
func (g *GameSim) Draw(screen *ebiten.Image) {
	util.DbgCountFrames()
	screen.Fill(g.pal[gfx.GrayDark])

	if g.maingrid.Draw {
		screen.DrawImage(g.maingrid.Img, g.maingrid.Op)
	}

	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, g.gWidth/16, 0)
}
func (g *GameSim) isSimTick() bool {
	return int(g.ticks)%(g.SimSpeed /*64-g.SimSpeed*/) == 0
}

// Layout of GameSim window (screen/GameSim)
func (g *GameSim) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	util.DbgCountLayout()
	//TODO: Write functionality for scaling.
	return g.gWidth, g.gHeight
}

func (g *GameSim) debugUpdate() {
	defer util.Dbg.DebugBuildOutput()
	util.DbgCountTicks()
	util.Dbg.UpdateDetail = fmt.Sprintf(
		"||SPD:%d Cut:%d",
		g.SimSpeed, entity.CutoffIs())
}
