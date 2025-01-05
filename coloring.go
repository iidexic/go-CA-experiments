package main

import (
	"crypto/rand"
	"image/color"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/iidexic/go-CA-experiments/utils"
)

/* unused interface/structs
// sequence is an interface for sequences. stepped/slewed sequences of color changes/gradients
type sequence interface {
	apply(g gobject)
}
*/

// colorchange - struct for defining a shift from start -> end color in stepcount steps
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

// randcolors returns quantity colors requested
func randcolors(size int) []color.RGBA {
	var cs []color.RGBA = make([]color.RGBA, size)
	for i := range size {
		cs[i] = Randcolor()
	}
	return cs
}

// Randfade generates gradient between random color and black in `stepcount` steps
// TODO rewrite to use colorchange struct
func Randfade(stepcount uint8) []color.RGBA {
	var cs []color.RGBA = make([]color.RGBA, stepcount)
	cs[0] = Randcolor()
	//fade := make([]byte, 4)
	for i := uint8(1); i < stepcount; i++ {

	}
	return cs
}

// Randcolor returns a color.RGBA with pseudo-random red, green, blue, alpha
func Randcolor() color.RGBA {
	var i []byte = make([]byte, 4)
	_, _ = rand.Read(i)
	return color.RGBA{i[0], i[1], i[2], i[3]}
}

//======================================================
//======================================================
//=== Not currently used ===============================
/*

 */
