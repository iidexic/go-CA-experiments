package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type anchortype int

const (
	anchorFree anchortype = iota
	anchorCenter
	anchorTop
	anchorBtm
	anchorLeft
	anchorRight
)

type positioning struct {
	anchor    anchortype
	anchorObj vobj
}

// Img2GeoM takes the image rect's Min point and translates the GeoM to match it
func Img2GeoM(img *ebiten.Image, gm *ebiten.GeoM) {
	rct := img.Bounds()
	gm.Translate(float64(rct.Min.X), float64(rct.Min.Y))
}
