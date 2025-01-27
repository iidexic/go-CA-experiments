package gfx

import (
	"image/color"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/util"
)

// Palette holds  main palette colors. Commment shows color/use
var Palette []color.RGBA = []color.RGBA{
	{32, 29, 31, 255}, //background
}

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

//=== Not currently used ===============================

//func gradientRemainder(n, steps,sd, sr int) int{
// 	not useful time spend to make more good
// 	if n in first half (,steps//2) and stepnum % steps/sr == 0)
// 	if in same range on other side of halfway
//}
