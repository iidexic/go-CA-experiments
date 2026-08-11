package gfx

import (
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/fileops"
)

type Sprite struct {
	*ebiten.Image
	Opt *ebiten.DrawImageOptions
}

func (s *Sprite) Img() *ebiten.Image {
	return s.Image
}
func (s *Sprite) ImgOpt() *ebiten.DrawImageOptions {
	return s.Opt
}

type SpriteLib struct {
	Sprites map[string]*Sprite
}

func (s *SpriteLib) GetSprite(name string) *Sprite {
	if s.Sprites[name] == nil {
		return nil
	}
	return s.Sprites[name]
}

func NewSprite(name string) *Sprite {
	s := Sprite{
		Image: ebiten.NewImage(16, 16),
		Opt:   &ebiten.DrawImageOptions{},
	}
	s.Opt.GeoM.Translate(8, 8)
	LoadedSprites.Sprites[name] = &s
	return &s
}

// NewSpriteFromFile loads a sprite from a file. Returns nil if error.
func NewSpriteFromFile(path string) (*Sprite, string) {
	img, err := fileops.ImageFromFile(path)
	if err != nil {
		return nil, ""
	}
	s := Sprite{
		Image: img,
		Opt:   &ebiten.DrawImageOptions{},
	}
	s.Opt.GeoM.Translate(8, 8)
	name := filepath.Base(path)
	LoadedSprites.Sprites[name] = &s

	return &s, name
}
func DrawSprites(screen *ebiten.Image, names []string) {
	for _, name := range names {
		if s := LoadedSprites.GetSprite(name); s != nil {
			screen.DrawImage(s.Image, s.Opt)
		}
	}
}

var LoadedSprites = SpriteLib{Sprites: make(map[string]*Sprite)}

func init() {

}
