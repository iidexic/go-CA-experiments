package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img                   *ebiten.Image
	X, Y                  uint
	Bounds                []int
	modAdd, modMult, Area int
	Px, Highlight, rng    []byte
	Op                    ebiten.DrawImageOptions
	Draw, Debug           bool
	reload                func()
	DebugString           string
}

var testCutoff byte = 128

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity {
	borderTot := (gWidth / 16)
	width := gWidth - borderTot   //(31 * gWidth) / 32
	height := gHeight - borderTot //(31 * gHeight) / 32
	grid := GridEntity{

		Img:    ebiten.NewImage(width, height),
		Bounds: make([]int, 4),
		Op:     ebiten.DrawImageOptions{},
		X:      uint(width), Y: uint(height), Area: width * height,
		Px:     make([]byte, width*height*4),
		rng:    make([]byte, width*height),
		modAdd: 0, modMult: 1, //> Unknown if needed
		Highlight: []byte{127, 127, 127, 160, 255, 40, 44, 40, 0, 1, 191, 40},
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

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//****************************************************************************
//****************************************************************************

func pxisort(pix []byte) []int {
	ts := make([]int, 3)
	if pix[0] > pix[1] {
		ts[0] = 1
		ts[1] = 0
	} else {
		ts[0] = 0
		ts[1] = 1
	}
	if pix[ts[1]] > pix[2] {
		if pix[ts[0]] > pix[2] {
			ts[2] = ts[1]
			ts[1] = ts[0]
			ts[0] = 2
			return ts[:3]
		}
		ts[2] = ts[1]
		ts[1] = 2
		return ts[:3]
	}
	ts[2] = 2
	return ts
}

// XY returns grid GeoM tx, ty screen location
func (grid *GridEntity) XY() (int, int) {
	return int(grid.Op.GeoM.Element(0, 2)), int(grid.Op.GeoM.Element(1, 2))
}

func (grid *GridEntity) getrng(i int) byte { return grid.rng[i%grid.Area] }

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

func (grid *GridEntity) exec1v1(o outcome, ipx, epx int) {
	i := ipx * 4
	e := epx * 4
	switch o {
	case ilose:
		sliceToward(grid.Px[i:i+3], grid.Px[e:e+3], 150)
	case iwin:
		sliceToward(grid.Px[e:e+3], grid.Px[i:i+3], 150)
	case ifriend:
		grid.interactFriend(i, e)
	case imine:
		grid.interactMine(i, e)
	case istale:
		grid.interactStalemate(i, e)
	case ineut:
		grid.interactNeutral(i, e)
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

func (grid *GridEntity) interactMine(i, m int) {
	ip := grid.Px[i : i+4]
	ival := bavg(ip[:3]...) //TODO: prevent needing to calc multiple times per pixel per frame
	light := ival > 128
	mp := grid.Px[m : m+4]
	s := pxisort(mp[:3])
	//[Mine Behavior]
	//- For the pixel being mined (m); each byte val is resource.
	//- 3-color is most energy/hardest to mine
	//- going down to 1-color,  easiest to mine
	if light {
		mineLV1 := mp[s[2]] - mp[s[1]]
		mineLV2 := mp[s[1]] - mp[s[0]]
		switch {
		case mineLV1 > 15:
			mp[s[2]], ip[s[2]] = bmov(mp[s[2]], ip[s[2]], 16)
		case mineLV2 > 3:
			mp[s[2]], ip[s[2]] = bmov(mp[s[2]], ip[s[2]], 4)
			mp[s[1]], ip[s[1]] = bmov(mp[s[1]], ip[s[1]], 4)
		default:
			bsladd(ip[:3], 1)
			bslsub(mp[:3], 1)
		}
	} else {
		mineLV1 := (255 - mp[s[0]]) - (255 - mp[s[1]])
		mineLV2 := (255 - mp[s[1]]) - (255 - mp[s[2]])
		switch {
		case mineLV1 > 15:
			ip[s[0]], mp[s[0]] = bmov(ip[s[0]], mp[s[0]], 16)
		case mineLV2 > 3:
			ip[s[0]], mp[s[0]] = bmov(ip[s[0]], mp[s[0]], 4)
			ip[s[1]], mp[s[1]] = bmov(ip[s[1]], mp[s[1]], 4)
		default:
			bsladd(mp[:3], 1)
			bslsub(ip[:3], 1)
		}
	}
}

func bmov(src, dest byte, amt byte) (byte, byte) {
	lim := min(src, 255-dest, amt)
	src -= lim
	dest += lim
	return src, dest
}
func (grid *GridEntity) interactNeutral(i, e int) {

}
func (grid *GridEntity) interactStalemate(i, e int) { //! FIX CRASH :)
	ipx := grid.Px[i : i+4]
	epx := grid.Px[e : e+4]
	srng := int(grid.getrng(i+2)) + int(grid.getrng(i+1))
	xtrarng := grid.getrng(i + 666)
	if xtrarng > 252 { //makes colored noise, lower the threshold = more noise
		//do the XOR
		for n := range ipx {
			ipx[(srng*e+n)%3] ^= epx[(srng*i+n)%3]
			epx[(srng*i+n)%3] ^= ipx[(srng*e+2+n)%3]
		}
	} else if bavg(ipx[:3]...) > 128 {
		bsladd(ipx[:3], 16)
		bslsub(epx[:3], 16)
	} else if bavg(ipx[:3]...) < 127 {
		bsladd(epx[:3], 16)
		bslsub(ipx[:3], 16)
	}
}
func (grid *GridEntity) interactFriend(i, e int) {
	//! untested
	ip := grid.Px[i : i+3]
	ep := grid.Px[e : e+3]

	ipavg := acdbavg(ip...) //TODO: rewrite
	epavg := acdbavg(ep...)

	alignment := bavg(ip...) > 128
	if ipavg > epavg {
		sliceToward(ip, ep, 64)
	} else if epavg > ipavg {
		sliceToward(ep, ip, 64)
	} else { //? possibly add a re-ordering of RGB values to match i's
		si := pxisort(ip)
		se := pxisort(ep)
		if alignment {
			if ip[si[2]] > ep[se[2]] {
				temp := ep[si[2]]
				ep[si[2]] = ep[se[2]]
				ep[se[2]] = temp
			} else {
				temp := ip[se[2]]
				ip[se[2]] = ip[si[2]]
				ip[si[2]] = temp
			}
			//bsladd(ip, 128) //bsladd(ep, 128)
		} else {
			if ip[si[0]] < ip[se[0]] {
				temp := ep[si[0]]
				ep[si[0]] = ep[se[0]]
				ep[se[0]] = temp
			} else {
				temp := ip[se[0]]
				ip[se[0]] = ip[si[0]]
				ip[si[0]] = temp
			}
			//bslsub(ip, 128) //bslsub(ep, 128)
		}
	}
}
func moveToward(from, to byte, amount byte) byte {
	dist := int(to) - int(from)
	dx := (dist * int(amount)) / 255
	return byte(int(from) + dx)
	// max |dx| == |dist|
}

// given two slices, moves one toward the other, specified by byte
// to will loop if from larger than to
func sliceToward(from, to []byte, amount byte) {
	lto := len(to)
	for i := range from {
		from[i] = moveToward(from[i], to[i%lto], amount)
	}
}

var (
	overlayRed  []byte = []byte{240, 100, 100, 120}
	overlayBlue        = []byte{0, 0, 180, 140}
	overlayMid         = []byte{100, 129, 100, 205}
)

// ApplyDbgOverlay does that.
func (grid *GridEntity) ApplyDbgOverlay(mode int) []byte {
	overlay := make([]byte, len(grid.Px))
	var r int
	for i := range grid.Area {
		r = i * 4
		ba := bavg(grid.Px[r : r+3]...)
		switch {
		case ba > 128:
			mto(grid.Px[r:r+4], overlayRed, overlay[r:r+4])
		case ba < 127:
			mto(grid.Px[r:r+4], overlayBlue, overlay[r:r+4])
		case ba == 127 || ba == 128:
			mto(grid.Px[r:r+4], overlayMid, overlay[r:r+4])
		}
	}
	return overlay
}

func mto(bs1, bs2, dest []byte) {
	for i := range dest[:3] {
		dest[i] = moveToward(bs1[i], bs2[i], bs2[3])
	}
	dest[3] = 255
}

// moves color from one px to another.
// * up==true will move color from px1 to px2 and vice versa
func cmov(px1, px2 []byte, up bool, amt ...byte) {
	for i, a := range amt {
		if up {
			lim := min(px1[i], 255-px2[i], a)
			px1[i] -= lim
			px2[i] += lim
		} else {
			lim := min(px2[i], 255-px1[i], a)
			px1[i] += lim
			px2[i] -= lim
		}
	}
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
