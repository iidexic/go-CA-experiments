package core

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/input"
	"github.com/iidexic/go-CA-experiments/util"
)

func testMove(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(1, 1)
}

// Game struct - ebiten
type Game struct {
	maingrid                         *entity.GridEntity
	pal                              []color.RGBA
	gWidth, gHeight, pWidth, pHeight int
	RunSimulation, modAdd, modMult   int
}

// GameInit returns Game pointer for main sim scene with default settings
func GameInit(gameWidth, gameHeight int) *Game {
	g := &Game{
		RunSimulation: 0,
		modAdd:        1,
		modMult:       4,
		gWidth:        gameWidth,
		gHeight:       gameHeight,
		maingrid:      entity.MakeGridDefault(gameWidth, gameHeight),
		pal:           gfx.Palette,
	}
	return g
}

//===============================================================

// Update game - game logic, assume locked at 60TPS.
func (g *Game) Update() error {
	g.debugUpdate()

	if g.RunSimulation > 0 {
		g.maingrid.SetMod(g.modAdd, g.modMult)
		g.maingrid.SimstepLVSD(true)
		g.maingrid.Img.WritePixels(g.maingrid.Pixels)
	}
	inputActions(g)
	return nil
}

// Draw screen
func (g *Game) Draw(screen *ebiten.Image) { //^DRAW
	util.DbgCountFrames()
	screen.Fill(g.pal[0])

	if g.maingrid.Draw {
		screen.DrawImage(g.maingrid.Img, &g.maingrid.Op)
	}
	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)

}

// Layout of game window (screen/game)
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	util.DbgCountLayout()
	//TODO: Write functionality for scaling.
	return g.gWidth, g.gHeight
}

func (g *Game) debugUpdate() {
	defer util.Dbg.DebugBuildOutput()
	util.DbgCountTicks()
	input.GetInKB() //DEBUG USE
	util.Dbg.UpdateDetail = fmt.Sprintf("||U:add/shift=%d, multBitmax=%d", g.modAdd, g.modMult)
}
