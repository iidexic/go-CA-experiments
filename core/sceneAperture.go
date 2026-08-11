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
	D           *gfx.WindowConfig
	pal         []color.RGBA
	sqr         *entity.BaseEntity
	cyc         cycler
	spriteNames []string
}

var gamescene *GameScene

func GetScene(pxWidth, pxHeight int) *GameScene {
	gamescene = &GameScene{
		D:           gfx.MainWindow(),
		pal:         gfx.PaletteGP,
		cyc:         cycler{ticks: 0},
		spriteNames: make([]string, 0),
	}
	gamescene.init()
	return gamescene
}
func (g *GameScene) init() {
	g.sqr = entity.NewBaseEntity(16, 16)
	g.cyc = cycler{ticks: 0}
	// test sprite
	sprite, sp2 := gfx.NewSpriteFromFile("Resources/sp2_test.png")
	if sprite == nil {
		util.Dbg.AddErrorF("sprite load error(%s)", sp2)
	} else {
		util.Dbg.AddErrorF("sprite '%s' loaded: %+v", sp2, sprite)
	}
	gfx.Scale(sprite, 0.5, 0.5)

}

// Update Game Method
func (g *GameScene) Update() error {
	debugUpdate(g)
	g.cyc.ticks++

	return nil
}

// Draw Game method
func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(g.pal[gfx.Dark])
	// g.drawEntityList(screen)

	//screen.DrawImage(g.sqr.Img, g.sqr.Opt)
	gfx.DrawSprites(screen, g.spriteNames)
	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)
}

// Layout Game method
func (g *GameScene) Layout(wWidth, wHeight int) (gameX, gameY int) {
	return g.D.LayoutGameSize(wWidth, wHeight)
}

func debugUpdate(g *GameScene) {
	defer util.Dbg.DebugBuildOutput()
	ws, gs := g.D.LastSizes()
	util.Dbg.SetValues(gs[0], gs[1], ws[0], ws[1])
	util.DbgCountTicks()
	//input.GetInKB() //DEBUG USE
}
