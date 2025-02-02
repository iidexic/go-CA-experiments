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

// GameTest - test scene Game struct
type GameTest struct {
	gWidth, gHeight int
	//? should the game not have screen in it?
	pal []color.RGBA //RGBA bytes - same as ebiten Pixel slice
	sqr *entity.BaseEntity
}

// GameTestInit function to initialize Game obj and start gameloops from main
func GameTestInit(width, height int) *GameTest {
	g := &GameTest{
		gWidth:  width,
		gHeight: height,
		sqr:     makeSquare(20, 20),
		pal:     gfx.PaletteWB,
	}
	return g
}

// Update Game Method
func (g *GameTest) Update() error {
	g.debugUpdate()

	return nil
}

// Draw Game method
func (g *GameTest) Draw(screen *ebiten.Image) {
	util.DbgCountFrames()
	screen.Fill(g.pal[gfx.WhiteTan])

	// g.drawEntityList(screen)

	screen.DrawImage(g.sqr.Img, g.sqr.Opt)

	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)
}

// Layout Game method
func (g *GameTest) Layout(wWidth, wHeight int) (gameX, gameY int) {
	return g.gWidth, g.gHeight
}
func (g *GameTest) debugUpdate() {
	defer util.Dbg.DebugBuildOutput()
	util.DbgCountTicks()
	input.GetInKB() //DEBUG USE
	util.Dbg.UpdateDetail = fmt.Sprintf("")
}
