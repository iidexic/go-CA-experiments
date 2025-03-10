package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img                     *ebiten.Image
	X, Y                    uint
	Bounds                  []int
	modAdd, modMult, Area   int
	Px, mem, Highlight, rng []byte
	memsize, Fails          int
	Op                      ebiten.DrawImageOptions
	Draw, Debug             bool
	pxsize                  byte
	reload                  func()
}

// this is a type alias:
type outcome = int

// outcome enum:
const (
	ineut outcome = iota
	ilose
	iwin
	istale
	ifriend
	imine
)

var testCutoff byte = 128

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity {
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{

		Img:    ebiten.NewImage(width, height),
		Bounds: make([]int, 4),
		Op:     ebiten.DrawImageOptions{},
		X:      uint(width), Y: uint(height), Area: width * height,
		Px:     make([]byte, width*height*4),
		rng:    make([]byte, 300),
		modAdd: 0, modMult: 1, //> Unknown if needed
		Highlight: []byte{127, 127, 127, 160, 255, 40, 44, 40, 0, 1, 191, 40},
		pxsize:    1,
	}
	grid.reload = gfx.Fbytes(grid.rng)
	grid.Bounds[0] = (gWidth - width) / 2
	grid.Bounds[2] = grid.Bounds[0] + width

	grid.Bounds[1] = (gHeight - height) / 2
	grid.Bounds[3] = grid.Bounds[1] + height
	grid.Op.GeoM.Translate(float64(grid.Bounds[0]), float64(grid.Bounds[1]))
	grid.Img.Fill(gfx.PaletteGP[gfx.Dark])

	return &grid
}

// XY returns grid GeoM tx, ty screen location
func (grid *GridEntity) XY() (int, int) {
	return int(grid.Op.GeoM.Element(0, 2)), int(grid.Op.GeoM.Element(1, 2))
}

/*
	func (grid *GridEntity) rngv() byte {
		if len(grid.crng) < 4{
		}
		return <-grid.crng
	}
*/
func (grid *GridEntity) getrng(i int) byte {
	return grid.rng[i%290]
}
func (grid *GridEntity) exec1v1(o outcome, ipx, epx int) {
	i := ipx * 4
	e := epx * 4
	switch o {
	case ilose:
		sliceToward(grid.Px[i:i+3], grid.Px[e:e+3], 150)
	case iwin:
		sliceToward(grid.Px[e:e+3], grid.Px[i:i+3], 150)
	case ifriend:
		//! Think through this - not workin great
		difP := make([]int, 3)
		var iMaxDiff int
		for n := range 3 {
			difP[n] = int(grid.Px[e+n]) - int(grid.Px[i+n])
			if difP[n] < 0 {
				difP[n] = -difP[n]
			}
			if difP[n] > difP[iMaxDiff] {
				iMaxDiff = n
			}
		}
		dT := grid.Px[i+iMaxDiff]
		grid.Px[i+iMaxDiff] = grid.Px[e+iMaxDiff]
		grid.Px[e+iMaxDiff] = dT
	case imine:

	case istale:
		//grid.Px[i]
	case ineut:

	}

}

// SimstepLVSD performs one cycle/screen of checks and updates
// for the center-distance intensity comparison sim ("Light VS Dark")
func (grid *GridEntity) SimstepLVSD(pixLock bool) {
	grid.reload() //roll rng
	for i := range grid.Area {
		irng := grid.getrng(i)
		up := wrap(i-int(grid.X), grid.Area)
		lft := sidewrap(i, -1, int(grid.X))
		iR := i * 4
		upR := up * 4
		lftR := lft * 4

		ival := bavg(grid.Px[iR : iR+3]...) // averaged value of pixel colors
		uval := bavg(grid.Px[upR : upR+3]...)
		lval := bavg(grid.Px[lftR : lftR+3]...)

		results := versusLVSD(irng, ival, uval, lval) // (currently) return slice of lightWin bools.
		standinrng := uval ^ lval
		if standinrng > 127 {
			grid.exec1v1(results[1], i, lft)
			grid.exec1v1(results[0], i, up)
		} else {
			grid.exec1v1(results[0], i, up)
			grid.exec1v1(results[1], i, lft)
		}
	}
}

// ? Do we need to be using 255 or 256 here for these.
func moveToward(from, to byte, amount byte) byte {

	var dfin int
	dist := int(to) - int(from) //*=0 to 255

	if dist < 0 {
		dist = -dist
	}
	dmult := dist * int(amount)
	if dmult >= 255 {
		dfin = dmult / 255
	} else {
		dfin = dist
	}
	return from + byte(dfin)
}

// given two slices, moves one toward the other, specified by byte
// to will loop if from larger than to
func sliceToward(from, to []byte, amount byte) {
	lto := len(to)
	for i := range from {
		from[i] = moveToward(from[i], to[i%lto], amount)
	}
}

func versusLVSD(rng byte, iClr byte, versus ...byte) []outcome {
	wout := make([]outcome, len(versus))

	if iClr == 127 || iClr == 128 {
		return wout
	}
	alignment := iClr > 128
	//===TODO: No reason to go through conditionals here and then do a return and send that to another function to check the conditions again to find out what needs to be ran.
	for i, v := range versus {
		if v == 127 || v == 128 { // mine (future functionality) if v neutral
			wout[i] = imine
		} else if (v > 128) == alignment { // if vs alignment == mc alignment
			wout[i] = ifriend
		} else { // the actual battle
			rval := battlemc(iClr, v, (testCutoff + rng))
			switch {
			case rval > 0:
				wout[i] = iwin
			case rval < 0:
				wout[i] = ilose
			case rval == 0:
				wout[i] = istale
			}
		}
	}

	return wout
}

// battlemc takes mc and enemy, and returns result
// Output int: sign = win/lose, size = by how much.
func battlemc(mainchar, enemy, rng byte) (mcWin int) {
	var victoryLine byte = mainchar + enemy - 128 //r<victoryLine = lightWin
	mcWin = int(rng) - int(victoryLine)           // positive = lightWin
	if mainchar < 127 {                           // if mc is not light, switch lightwin to darkwin
		return -mcWin
	}
	return mcWin
}

// DebugOverlay dispays colors based on the current state of pixels.
// alpha will be used here to use as an overlay on top of the regular grid
func (grid *GridEntity) DebugOverlay() []byte {
	overlay := make([]byte, len(grid.Px))
	for i := range grid.Area {
		icol := i * 4
		avgC := bavg(grid.Px[icol : icol+3]...)
		if avgC == 127 || avgC == 128 { //Neutral

			//sliceToward(overlay[icol:icol+4], grid.Highlight[:4], 128)
			overlay[icol] = 127
			overlay[icol+1] = 127
			overlay[icol+2] = 127
			overlay[icol+3] = 40 //alpha

		} else if avgC > 128 { // light = red?
			//sliceToward(overlay[icol:icol+4], grid.Highlight[4:8], 128)

			overlay[icol] = 255
			overlay[icol+1] = 20
			overlay[icol+2] = 20
			overlay[icol+3] = 40 //alpha

		} else { // dark = blue
			//sliceToward(overlay[icol:icol+4], grid.Highlight[8:], 128)

			overlay[icol] = 20
			overlay[icol+1] = 20
			overlay[icol+2] = 120
			overlay[icol+3] = 10 //alpha

		}
	}
	return overlay
}

// CutoffUp is to manually change victory cutoff to see effects in real-time
func CutoffUp() {
	t := testCutoff
	switch {
	case t < 148 && t > 108:
		t += 2
	case t >= 148 && t < 245:
		t += 8
	case t <= 108:
		t += 8
	}
	testCutoff = t
}

// CutoffDown is to manually change victory cutoff to see effects in real-time
func CutoffDown() {
	t := testCutoff
	switch {
	case t < 148 && t > 108:
		t -= 2
	case t >= 148 && t < 254:
		t -= 8
	case t <= 108 && t > 8:
		t -= 8
	}
	testCutoff = t
}

// CutoffIs returned
func CutoffIs() byte {
	return testCutoff
}
