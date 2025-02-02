package core

import (
	"fmt"
	"image/color"
	"math"

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

type numeric interface {
	int | ~uint | ~float64
}

// GameSim struct - ebiten
type GameSim struct {
	maingrid                         *entity.GridEntity
	entlist                          []entity.Entity
	pal                              []color.RGBA
	gWidth, gHeight, pWidth, pHeight int
	RunSimulation, modAdd, modMult   int
	sqr                              *entity.BaseEntity
}

// GameSimInit returns GameSim pointer for main sim scene with default settings
func GameSimInit(GameSimWidth, GameSimHeight int) *GameSim {
	g := &GameSim{
		RunSimulation: 0,
		modAdd:        1,
		modMult:       4,
		gWidth:        GameSimWidth,
		gHeight:       GameSimHeight,
		maingrid:      entity.MakeGridDefault(GameSimWidth, GameSimHeight),
		pal:           gfx.PaletteWB,
	}
	//==== TESTING STUFF ====
	g.sqr = makeSquare(16, 16)

	//=======================
	return g
}

// ===============================================================
// Draw/movement testing
// Trying to use interface to make a more broad approach.
// But probably just use the struct methods
func makeSquare(width, height int) *entity.BaseEntity {
	sq := entity.NewBaseEntity(width, height)

	sq.Img.Fill(color.RGBA{255, 40, 230, 255})
	//sq.GeoM.Scale(1.1, 1.1)
	sq.GeoM.Translate(20, 20)
	return sq
}

// ** This seems wack
func centr[N1, N2, N3, N4 numeric](width N1, height N2, tx N3, ty N4) (float64, float64) {
	return (float64(width) + float64(tx)) / 2, (float64(height) + float64(ty)) / 2
}
func (g *GameSim) tsqRotAroundCenter(rad float64) {
	w, h := g.sqr.Img.Bounds().Dx(), g.sqr.Img.Bounds().Dy()
	tx, ty := g.sqr.GeoM.Element(2, 0), g.sqr.GeoM.Element(2, 1)
	dcentrX, dcentrY := centr(w, h, tx, ty)
	g.sqr.GeoM.Translate(dcentrX, dcentrY) //center the origin
	g.sqr.GeoM.Rotate(rad)

}
func (g *GameSim) testSquarePosition() {
	//** to do not corner rotation, shift - 1/2 of bounds xy, then shift back.
	//* This would have to be done before translation happens every step??
	// Do we like un-translate right after the fucking draw or what
	w, h := g.sqr.Img.Bounds().Dx(), g.sqr.Img.Bounds().Dy() //grab bounds
	g.sqr.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)   //center the origin
	g.sqr.GeoM.Rotate(float64(1) / 96.0 * math.Pi / 6)       // perform rotate. uncertain exact purpose of div/mult
	g.sqr.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)     //put back to proper location
	//? Can we include current geom translate? I think so
	//? Next line somehow did not fuck things up in flappy, seems like the actual translation? I can't tell.
	// Could also be the sprite change? Seems super unlikely though.
	//g.sqr.GeoM.Translate(float64(20.0/16.0), float64(20.0/16.0))
}

//===============================================================

// Update GameSim - GameSim logic, assume locked at 60TPS.
func (g *GameSim) Update() error {
	g.debugUpdate()

	if g.RunSimulation > 0 {
		g.maingrid.SetMod(g.modAdd, g.modMult)
		g.maingrid.SimstepLVSD(true)
		g.maingrid.Img.WritePixels(g.maingrid.Pixels)
	}
	// rotation/movement stuff happn within draw? how driven?
	//== test box draw move rotate

	//==
	inputActions(g)
	return nil
}

// Draw screen
func (g *GameSim) Draw(screen *ebiten.Image) { //^DRAW
	util.DbgCountFrames()
	screen.Fill(g.pal[0])

	if g.maingrid.Draw {
		screen.DrawImage(g.maingrid.Img, &g.maingrid.Op)
	}
	//*= entitylist temporarily not in use.
	// g.drawEntityList(screen)
	g.testSquarePosition()
	//== test sqr draw
	screen.DrawImage(g.sqr.Img, g.sqr.Opt)

	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)

}
func (g *GameSim) drawEntityList(screen *ebiten.Image) {
	for _, ent := range g.entlist {
		screen.DrawImage(ent.GetImg(), ent.GetOpt())
	}
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
	input.GetInKB() //DEBUG USE
	util.Dbg.UpdateDetail = fmt.Sprintf("||U:add/shift=%d, multBitmax=%d", g.modAdd, g.modMult)
}
