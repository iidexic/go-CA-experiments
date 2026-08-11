package gfx

import "github.com/hajimehoshi/ebiten/v2"

type layerOrder int

const (
	layerOrderBackground layerOrder = iota
	layerOrderForeground
	layerOrderUI
)

type Layer struct {
	*ebiten.Image
	opts   *ebiten.DrawImageOptions
	thingy []byte
}

func (L *Layer) Draw(x, y int, img *ebiten.Image) {
	for i := 0; i < len(L.thingy); i++ {
		if L.thingy[i] == 1 {
			img.DrawImage(ebiten.NewImageFromImage(img), L.opts)
		}
	}

}
