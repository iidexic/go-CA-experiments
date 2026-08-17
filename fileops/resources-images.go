package fileops

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func ImageFromFile(path string) (*ebiten.Image, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
