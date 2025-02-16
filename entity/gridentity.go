package entity

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img                   *ebiten.Image
	X, Y                  uint
	modAdd, modMult, Area int
	Px                    []byte
	Op                    ebiten.DrawImageOptions
	Draw                  bool // probably not in use
	statflags             byte // ()()()()()()()()
}

// this is a type alias:
type outcome = int

const (
	ineut outcome = iota
	ilose
	iwin
	istale
	ifriend
	imine
)

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity {
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{
		Img: ebiten.NewImage(width, height),
		Op:  ebiten.DrawImageOptions{},
		X:   uint(width), Y: uint(height), Area: width * height,

		Px:      make([]byte, width*height*4),
		modAdd:  0,
		modMult: 1,
	}
	grid.Op.GeoM.Translate(float64((gWidth-width)/2), float64((gHeight-height)/2))
	grid.Img.Fill(gfx.PaletteGP[gfx.Dark])
	return &grid
}

// SetMod sets modulation values modAdd and modMult
func (grid *GridEntity) SetMod(modAdd, modMult int) {
	grid.modAdd = modAdd
	grid.modMult = modMult
}
func (grid *GridEntity) exec1v1(o outcome, ipx, epx int) {
	i := ipx * 4
	e := epx * 4
	//removing all but win/lose to troubleshoot for now
	switch o {
	case ilose:
		for n := range 3 {

			grid.Px[i+n] = moveToward(grid.Px[i+n], grid.Px[e+n], 150)
		}
	case iwin:
		for n := range 3 {

			grid.Px[e+n] = moveToward(grid.Px[e+n], grid.Px[i+n], 150)
		}
	case ifriend:
		// trade biggest diff:
		difP := make([]int, 3)
		var iMaxDiff int
		for n := range difP {
			difP[n] = int(grid.Px[e+n]) - int(grid.Px[i+n])
			if difP[n] < 0 {
				difP[n] = -difP[n]
			}
			if difP[n] > difP[iMaxDiff] {
				iMaxDiff = n
			}
		}
		//^TEMPORARY!!:
		dT := grid.Px[i+iMaxDiff]
		grid.Px[i+iMaxDiff] = grid.Px[e+iMaxDiff]
		grid.Px[e+iMaxDiff] = dT
	}
}

// SimstepLVSD performs one cycle/screen of checks and updates
// for the center-distance intensity comparison sim ("Light VS Dark")
func (grid *GridEntity) SimstepLVSD(pixLock bool) {

	for i := 0; i < grid.Area; i++ {

		up := wrap(i-int(grid.X), grid.Area)
		lft := sidewrap(i, -1, int(grid.X))
		iR := i * 4
		upR := up * 4
		lftR := lft * 4

		ival := bavg(grid.Px[iR : iR+3]...) // averaged value of pixel colors
		uval := bavg(grid.Px[upR : upR+3]...)
		lval := bavg(grid.Px[lftR : lftR+3]...)

		results := versusLVSD(ival, uval, lval) // (currently) return slice of lightWin bools.
		grid.exec1v1(results[0], i, up)
		grid.exec1v1(results[1], i, lft)
		/*//TODO: 2 options to replace/clean this up
		//> 1. Have all operations (wrap, get R index, bavg, versusLVSD, apply results) in a loop
		//> 	1 loop = 1 battle (1 pair pixels vs) [This requires versusLVSD be rewritten]
		//> 2. Just do all of this shit in versusLVSD and return new pixel OR make it a method w/no return.
		*/
		//For now its 2 cases. just do it twice.

	}
}
func (grid *GridEntity) px2px(i, to, amt byte) {
	if grid.Px[i] > grid.Px[to] {
		grid.Px[i] += ((grid.Px[i] - grid.Px[to]) * amt) % 255
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
func versusLVSD(iClr byte, versus ...byte) (versusResult []outcome) {
	rand := make([]byte, 1)

	rndv, e := fastrand.Read(rand) //fix later
	erch(e)
	_ = rndv //> implement after functional
	wout := make([]outcome, len(versus))

	if iClr == 127 || iClr == 128 {
		return wout
	}
	alignment := iClr > 128

	for i, v := range versus {

		if v == 127 || v == 128 { // mine (future functionality) if v neutral
			wout[i] = imine
		} else if (v > 128) == alignment { // if vs alignment == mc alignment
			wout[i] = ifriend
		} else { // the actual battle
			var lightwin bool
			if alignment { //i light
				lightwin = battle(iClr, v)
			} else {
				lightwin = battle(v, iClr)
			}
			if lightwin == alignment {
				wout[i] = iwin
			} else {
				wout[i] = ilose
			}
		}
	}

	return wout
}

// battlemc takes mc and enemy, and returns result
// Output int: sign = win/lose, size = by how much. (for bool: return lightVictor && mainchar>128)
func battlemc(mainchar, enemy, rng byte) (mcWin int) {
	var victoryLine byte = mainchar + enemy - 128 //r<victoryLine = lightWin
	mcWin = int(rng) - int(victoryLine)           // positive = lightWin
	if mainchar < 127 {                           // if mc is not light, switch lightwin to darkwin
		return -mcWin
	}
	return mcWin
}

// we are going to simplify: see battlemc
func battle(lite, dark byte) (lightwin bool) {
	// maybe try a channel to get RNG? (would still need to pass or call internally)
	// worst case: 127-127+1 = 1 , opposite:127+127-1 = 253 * how to factor in the 1-unit dark advantage
	lPwr := lite - 128                                   // min 1 max 127
	dPwr := dark - 127                                   // min -127 max -1
	winpoint := 127 + lPwr + dPwr                        // skews toward dark
	fint := int(lPwr)*int(dark) + int((lite+1)%(dark+1)) //!!!ERROR DIV BY ZERO WITHOUT +1
	frand := byte(fint % 256)
	return frand > winpoint

}
func (grid *GridEntity) applyResult(results []outcome, i int, iVS ...int) {

}
func shiftMod(start, add, mult int, limit int) int {
	return wrap(((start + add) * mult), limit-3)
}
func wrapRange(start, len, add, mult int, limit int) (int, int) {
	ishift := wrap((start+add)*mult, limit)
	endshift := ishift + (len - 1)
	if endshift < limit {
		return ishift, endshift
	} //* if not then (for now) subtract length
	return ishift - (len - 1), endshift - (len - 1)
}

// SimstepValueShift is home of current more simplistic sim model after move on to full lvsd
// (keep pixlock for now) pixlock false allows offset color value averaging (i.e. blue vs red channel).
func (grid *GridEntity) SimstepValueShift(pixLock bool) {
	if !pixLock {
		for i := 0; i < len(grid.Px); i += 4 {
			//newR := shiftMod(i, grid.modAdd, grid.modMult, len(grid.Px))
			first, last := wrapRange(i, 3, grid.modAdd, grid.modMult, len(grid.Px))
			grid.pxGoToward(i, grid.Px[first:last])
		}
	} else {
		for i := 0; i < grid.Area; i++ {
			iR := i * 4
			newPix := shiftMod(i, grid.modAdd, grid.modMult, grid.Area)
			newR := newPix * 4 //area to RGB, will always land on R val

			grid.pxGoToward(iR, grid.Px[newR:newR+3])
		}

	}
}
func (grid *GridEntity) pxGoToward(indexR int, toPx []byte) {
	for i := range 3 {
		if grid.Px[indexR+i] > toPx[i] {
			grid.Px[indexR+i] -= (grid.Px[indexR+i] - toPx[i]) / 2
		} else {
			grid.Px[indexR+i] += (toPx[i] - grid.Px[indexR+i]) / 2
		}
	}
}
func (grid *GridEntity) pxReplace(indexR int, new []byte) {
	grid.Px[indexR] = new[0]
	grid.Px[indexR+1] = new[1]
	grid.Px[indexR+2] = new[2]
}

// pxTransplant overwrites 1px (3 indices) in grid.Px, starting at index
// write uses values at indices R,G,B without changes made during the function call
func (grid *GridEntity) pxTransplant(index int, R, G, B int) {
	var nR byte = grid.Px[R]
	var nG byte = grid.Px[G]
	var nB byte = grid.Px[B]
	grid.Px[index] = nR
	grid.Px[index+1] = nG
	grid.Px[index+2] = nB
}
