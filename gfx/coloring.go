package gfx

import (
	"image/color"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/hajimehoshi/ebiten/v2"
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
	if err != nil {
		panic(err)
	}
	for i := 3; i < len(b); i += 4 {
		b[i] = 255
	}
	return b
}

// Imagenoise directly generates and writes rand colors to existing ebiten img.
func Imagenoise(img *ebiten.Image) {
	area := uint(img.Bounds().Dx() * img.Bounds().Dy())
	img.WritePixels(Randpx(area))
}

//=== PALETTES ===============================

// > --------Palette WoodBlock-----------------------
/*
type indexWB int

const (
	WBDark indexWB = iota
	WBBrown
	WBPeach
	WBWhiteTan
	WBBlue
	WBMutedTeal
	WBSkyBlue
	WBGrayWarm
	WBRed
	WBOrange
	WBYellow
	WBGreen
)

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
*/
//> ----------------------------------------------------

// > -------Pallet Gummy Pickles-------------------------
type indexGP int

// Color names for PalleteGP
const (
	Yellow indexGP = iota
	Tan
	Green
	GrayDark
	GrayMid
	GrayLight
	White
	LightPink
	Pink
	Maroon
	Salmon
	PeachCoral
	Mustard
	DarkOrange
	MidRed
	Brown
	Lilac
	Orange
	LightOrange
	CornflowerBlue
	SeaBlue
	BluPurp
	DeepPurp
	Dark
)

// PaletteGP holds Gummy Pickles pallete for use throughout
var PaletteGP []color.RGBA = []color.RGBA{
	// Gummy-Pickles (24-color)
	{199, 175, 66, 255},  //yellow
	{165, 132, 73, 255},  //Tan mid (slight grn)
	{73, 106, 75, 255},   //Green
	{54, 73, 98, 255},    //GrayDark (muted blue)
	{99, 121, 140, 255},  //GrayMid
	{179, 197, 194, 255}, //GrayLight
	{239, 236, 232, 255}, //White
	{216, 163, 220, 255}, //LightPink
	{212, 98, 158, 255},  //Pink
	{133, 57, 91, 255},   //Maroon
	{206, 82, 99, 255},   //Salmon
	{241, 142, 116, 255}, //Peach
	{211, 147, 84, 255},  //MustardYlw
	{199, 102, 82, 255},  //DarkOrng
	{164, 72, 87, 255},   //MidRed
	{154, 92, 76, 255},   //Brown
	{120, 76, 173, 255},  //Lilac
	{229, 136, 82, 255},  //Orang
	{254, 204, 128, 255}, //LightOrange
	{142, 163, 230, 255}, //CornflowerBlue
	{58, 113, 166, 255},  //SeaBlue
	{75, 58, 166, 255},   //BluPurp
	{53, 41, 89, 255},    //DeepPurp
	{6, 13, 35, 255},     //Dark
}

// BPaletteGP is byte slice version of palette
// used when colors are needed to pull directly into pixel data
var BPaletteGP [][]byte = [][]byte{
	// BSPalette has an un-exported thingy
	// Gummy-Pickles (24-color)
	{199, 175, 66, 255},  //yellow
	{165, 132, 73, 255},  //Tan mid (slight grn)
	{73, 106, 75, 255},   //Green
	{54, 73, 98, 255},    //GrayDark (muted blue)
	{99, 121, 140, 255},  //GrayMid
	{179, 197, 194, 255}, //GrayLight
	{239, 236, 232, 255}, //White
	{216, 163, 220, 255}, //LightPink
	{212, 98, 158, 255},  //Pink
	{133, 57, 91, 255},   //Maroon
	{206, 82, 99, 255},   //Salmon
	{241, 142, 116, 255}, //Peach
	{211, 147, 84, 255},  //MustardYlw
	{199, 102, 82, 255},  //DarkOrng
	{164, 72, 87, 255},   //MidRed
	{154, 92, 76, 255},   //Brown
	{120, 76, 173, 255},  //Lilac
	{229, 136, 82, 255},  //Orang
	{254, 204, 128, 255}, //LightOrange
	{142, 163, 230, 255}, //CornflowerBlue
	{58, 113, 166, 255},  //SeaBlue
	{75, 58, 166, 255},   //BluPurp
	{53, 41, 89, 255},    //DeepPurp
	{6, 13, 35, 255},     //Dark
}
