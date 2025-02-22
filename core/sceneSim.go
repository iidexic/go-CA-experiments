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

// GameSim struct - ebiten
type GameSim struct {
	maingrid                         *entity.GridEntity
	pal                              []color.RGBA
	gWidth, gHeight, pWidth, pHeight int
	SimSpeed, modAdd, modMult, uTix  int
	sqr                              *entity.BaseEntity
	ticks                            uint16
	rngen                            *gfx.QuickRNG
}

// GameSimInit returns GameSim pointer for main sim scene with default settings
func GameSimInit(GameSimWidth, GameSimHeight int) *GameSim {
	rng := gfx.GetQuickRNG(64)
	g := &GameSim{
		SimSpeed: -8,
		modAdd:   1,
		modMult:  4,
		gWidth:   GameSimWidth,
		gHeight:  GameSimHeight,
		pal:      gfx.PaletteGP,
		rngen:    &rng,
	}
	g.maingrid = entity.MakeGridDefault(g.gWidth, g.gHeight, g.rngen.C)
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

func centr(width float64, height float64, tx float64, ty float64) (float64, float64) {
	return (width + tx) / 2, (height + ty) / 2
}
func (g *GameSim) testSquarePosition() {
	//** to do not corner rotation, shift - 1/2 of bounds xy, then shift back.
	w, h := g.sqr.Img.Bounds().Dx(), g.sqr.Img.Bounds().Dy() //grab bounds
	g.sqr.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)   //center the origin
	g.sqr.GeoM.Rotate(float64(1) / 96.0 * math.Pi / 6)       // perform rotate.
	g.sqr.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)     //put back to proper location
}

// Update function
func (g *GameSim) Update() error {
	g.debugUpdate()
	g.ticks++
	if g.SimSpeed > 0 && g.isSimTick() {
		g.maingrid.SetMod(g.modAdd, g.modMult)
		g.maingrid.SimstepLVSD(true)
		g.maingrid.Img.WritePixels(g.maingrid.Px)
	}
	inputActions(g)
	return nil
}

// Draw screen
func (g *GameSim) Draw(screen *ebiten.Image) { //^DRAW
	util.DbgCountFrames()
	screen.Fill(g.pal[gfx.GrayDark])

	if g.maingrid.Draw {
		screen.DrawImage(g.maingrid.Img, &g.maingrid.Op)
	}
	g.testSquarePosition()
	//== test sqr draw
	screen.DrawImage(g.sqr.Img, g.sqr.Opt)

	ebitenutil.DebugPrintAt(screen, util.Dbg.Output, 120, 0)

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
	input.GetInKB() //DEBUG USE
	util.Dbg.UpdateDetail = fmt.Sprintf("||SPD:%d Cutoff:%d", g.SimSpeed, entity.CutoffIs())
}
