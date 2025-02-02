package gfx

import (
	"image/color"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/util"
)

// Gradientbytes makes a gradient from color c1 to color c2 in number of steps
// Loop operation: (happens once for  each byte in color)
//
//	set start and end point (c1, c2 respectively)
//	calculate difference between c2, c1 and startpt
func Gradientbytes(c1 []byte, c2 []byte, steps uint8) []byte {
	//default 4 steps, to avoid divide by 0
	if steps == 0 {
		steps = 4
	}
	//variable prep
	var delta byte
	gr := make([]byte, (steps+1)*4) //steps = transitions hence +1. *4 for R,G,B,A
	isteps := int(steps)            //arg as uint8 to control range.
	//Loop through bytes in color (will run 4x, for R,G,B,A)
	for i, v := range c2 {
		//Load start color (c1) and end color (c2)
		gr[i] = c1[i]
		gr[isteps*i] = c2[i]
		//--------------
		var cstart byte
		//hardcode abs val to avoid underflow:
		if v >= c1[i] {
			delta = v - c1[i]
			cstart = c1[i] //! Is this correct?
		} else {
			delta = c1[i] - v
			cstart = v
		}
		d := delta / byte(steps)
		r := delta % byte(steps)
		var n int
		for n = 1; n < isteps; n++ {
			plusr := byte(0)
			if byte(n) <= r {
				plusr++
			}
			gr[i+4*n] = cstart + d*byte(n) + plusr
		}
	}
	return gr
}

// Randpx makes `pxcount` pseudo-random colors, all A=255
func Randpx(pxcount uint) []byte {
	b := make([]byte, pxcount*4)
	_, err := fastrand.Read(b)
	for i := 3; i < len(b); i += 4 {
		b[i] = 255
	}
	util.CheckPants(err)
	return b
}

// Imagenoise directly generates and writes rand colors to existing ebiten img.
func Imagenoise(img *ebiten.Image) {
	area := uint(img.Bounds().Dx() * img.Bounds().Dy())
	img.WritePixels(Randpx(area))
}

//=== PALETTES ===============================

// ColorselectWB aliases int for use as PaletteWB index name
type ColorselectWB int

// Color names for Palette
const (
	Dark ColorselectWB = iota
	Brown
	Peach
	WhiteTan
	Blue
	MutedTeal
	SkyBlue
	GrayWarm
	Red
	Orange
	Yellow
	Green
)

// PaletteWB holds primary palette. use with Colorselect
var PaletteWB []color.RGBA = []color.RGBA{
	{43, 40, 33, 255},
	{98, 76, 60, 255},
	{217, 172, 139, 255},
	{227, 207, 180, 255},
	{36, 61, 92, 255},
	{93, 114, 117, 255},
	{92, 139, 147, 255},
	{177, 165, 141, 255},
	{176, 58, 72, 255},
	{212, 128, 77, 255},
	{224, 200, 114, 255},
	{62, 105, 88, 255},
}
