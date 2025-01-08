package main

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/utils"
)

// Gradientbytes makes a gradient from color c1 to color c2 in number of steps
// currently annotated so I can remember what I did
func gradientbytes(c1 []byte, c2 []byte, steps uint8) []byte {
	//default 4 steps, to avoid divide by 0
	if steps == 0 {
		steps = 4
	}
	//slice+type prep
	var delta byte
	gr := make([]byte, (steps+1)*4)
	s := int(steps) //arg as uint8 to control range.
	for i, v := range c2 {

		gr[i] = c1[i]
		gr[s*i] = c2[i]
		//--------------
		var cstart byte
		if v >= c1[i] {
			delta = v - c1[i]
			cstart = c1[i]
		} else {
			delta = c1[i] - v
			cstart = v
		}
		d := delta / byte(steps)
		r := delta % byte(steps)
		var n int
		for n = 1; n < s; n++ {
			plusr := byte(0)
			if byte(n) <= r {
				plusr++
			}
			gr[i+4*n] = cstart + d*byte(n) + plusr
		}
	}
	return gr
}

// Randcolor returns a []byte with pseudo-random red, green, blue, alpha
func Randcolor() []byte {
	var i []byte = make([]byte, 4)
	_, _ = fastrand.Read(i)
	return i
}

// Randpx uses bytedance fastrand to generate pixelcount random RGB colors.
// as of now, it generates RGBA rand for all pix and replaces A with 255.
func Randpx(pixelcount uint) []byte {
	b := make([]byte, pixelcount*4)
	_, err := fastrand.Read(b)
	for i := 3; i < len(b); i += 4 {
		b[i] = 255
	}
	utils.CheckPants(err)
	return b
}
func imagenoise(img *ebiten.Image) {
	area := uint(img.Bounds().Dx() * img.Bounds().Dy())
	img.WritePixels(Randpx(area))
}

//======================================================
//======================================================
//=== Not currently used ===============================
/*
// sequence is an interface for sequences. stepped/slewed sequences of color changes/gradients
type sequence interface {
	apply(g gobject)
}
*/
//func gradientRemainder(n, steps,sd, sr int) int{
// 	not useful time spend to make more good
// 	if n in first half (,steps//2) and stepnum % steps/sr == 0)
// 	if in same range on other side of halfway
//}
/* original gradient work
type colorchange struct {
	start, end                color.RGBA
	delta, stepdelta, stepmod []byte
	result                    []byte
	stepcount                 uint8
}

// cc.calcDelta generates cc.delta, byte slice of start-end difference of RGBA values
func (cc colorchange) calcDelta() {
	cc.delta = make([]byte, 4)
	cc.delta[0] = cc.start.R - cc.end.R
	cc.delta[1] = cc.start.G - cc.end.G
	cc.delta[2] = cc.start.B - cc.end.B
	cc.delta[3] = cc.start.A - cc.end.A
}

// calcStep will use cc delta and size to determine required step size
func (cc colorchange) calcStep() {
	cc.stepdelta = make([]byte, 4)
	if cc.stepcount == 0 {
		cc.stepcount = 4
	}
	for i, bytedelta := range cc.delta {
		cc.stepdelta[i] = bytedelta / cc.stepcount
		cc.stepmod[i] = bytedelta % cc.stepcount
	}
}
func gradient(start, end color.RGBA, stepcount uint8) colorchange {
	cc := colorchange{start: start, end: end, stepcount: stepcount}
	cc.calcDelta()
	cc.calcStep()

	return cc
}

*/
