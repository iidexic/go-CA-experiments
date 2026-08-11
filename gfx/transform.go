package gfx

import "github.com/hajimehoshi/ebiten/v2"

type Visible interface {
	Img() *ebiten.Image
}
type Transformable interface {
	ImgOpt() *ebiten.DrawImageOptions
}
type imageInitializable interface {
	InitImg()
}

func Move(t Transformable, x, y int) {
	t.ImgOpt().GeoM.Translate(float64(x), float64(y))
}
func Scale(t Transformable, x, y float64) {
	t.ImgOpt().GeoM.Scale(x, y)
}
func Rotate(t Transformable, angle float64) {
	t.ImgOpt().GeoM.Rotate(angle)
}
func Flip(t Transformable, x, y int) {
	t.ImgOpt().GeoM.Scale(float64(x), float64(y))
}
