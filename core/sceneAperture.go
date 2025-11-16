package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/iidexic/go-CA-experiments/entity"
	"github.com/iidexic/go-CA-experiments/gfx"
	"github.com/iidexic/go-CA-experiments/util"
)

type GameScene struct {
	gWidth, gHeight int
	pal             []color.RGBA
	sqr             *entity.BaseEntity
	cyc             cycler
}

// Update Game Method
func (g *GameScene) Update() error {
	g.cyc.ticks++

	return nil
}

// Draw Game method
func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(g.pal[gfx.GrayDark])

	// g.drawEntityList(screen)

	screen.DrawImage(g.sqr.Img, g.sqr.Opt)

	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)
}

// Layout Game method
func (g *GameScene) Layout(wWidth, wHeight int) (gameX, gameY int) {
	return g.gWidth, g.gHeight
}
