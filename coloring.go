package main

import (
	"crypto/rand"
	"image/color"

	"github.com/iidexic/go-CA-experiments/utils"
)

/* unused interface/structs
// sequence is an interface for sequences. stepped/slewed sequences of color changes/gradients
type sequence interface {
	apply(g gobject)
}
*/
// TODO Determine best method of bridging gap between byte vals and int stepcount
// colorchange - struct for defining a shift from start -> end color in given stepcount
// max stepcount 255 - implement scale/dither if want longer
type colorchange struct {
	start, end                color.RGBA
	delta, stepdelta, stepmod []byte
	result                    []color.RGBA
	stepcount                 uint8
}

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

// Randcolor64 returns a pseudo-random generated 64x64 2d slice of color.RGBA
func Randcolor64() []color.RGBA {
	var bg []byte = make([]byte, 12288)
	fa := utils.Bytesum(bg[bg[1]:(bg[1] + 24)]) //semirandom start point for alpha to try and cut down random gens at least a little, will investigate necessity later
	_, err := rand.Read(bg)

	utils.CheckPants(err)

	//*Single-Array
	cbytes := [][]byte{bg[:4096], bg[4096:8192], bg[8192:12288], bg[fa : fa+4096]}

	return arrayToColor(cbytes)
}

// Randcolors returns quantity colors requested
func Randcolors(size int) []color.RGBA {
	var cs []color.RGBA = make([]color.RGBA, size)
	for i := range size {
		cs[i] = Randcolor()
	}
	return cs
}

// Randfade generates gradient between random color and black in `stepcount` steps
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

// ArrayToColor takes structured [][]byte array and loads into colors.
func arrayToColor(bg [][]byte) []color.RGBA {
	acolor := make([]color.RGBA, len(bg))
	for i, row := range bg { //TODO double-check direction of operation
		acolor[i] = color.RGBA{row[0], row[1], row[2], row[3]}
	}
	return acolor
}
