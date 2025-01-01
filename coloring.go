package main

import (
	"crypto/rand"
	"image/color"

	"github.com/iidexic/go-CA-experiments/utils"
)

func randcolor64() []color.RGBA {
	var bg []byte = make([]byte, 12288)
	fa := utils.Bytesum(bg[bg[1]:(bg[1] + 24)]) //semirandom start point for alpha to try and cut down random gens at least a little, will investigate necessity later
	_, err := rand.Read(bg)

	utils.CheckPants(err)

	//*Single-Array
	cbytes := [][]byte{bg[:4096], bg[4096:8192], bg[8192:12288], bg[fa : fa+4096]}

	return arrayToColor(cbytes)
}

func randcolor() color.RGBA {
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
