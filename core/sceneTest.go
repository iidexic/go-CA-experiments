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
	pal             []color.RGBA
	sqr             *entity.BaseEntity
	cyc             cycler
}

// cycler counts cycles to be sent to anywhere game timing is needed
type cycler struct {
	ticks, frames int
}

// GameTestInit function to initialize Game obj and start gameloops from main
func GameTestInit(width, height int) *GameTest {
	g := &GameTest{
		gWidth:  width,
		gHeight: height,
		sqr:     makeSquare(20, 20),
		pal:     gfx.PaletteGP,
	}
	return g
}

// Update Game Method
func (g *GameTest) Update() error {
	g.cyc.ticks++
	g.debugUpdate()

	return nil
}

// Draw Game method
func (g *GameTest) Draw(screen *ebiten.Image) {
	g.cyc.frames++
	util.DbgCountFrames()
	screen.Fill(g.pal[gfx.GrayLight])

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
