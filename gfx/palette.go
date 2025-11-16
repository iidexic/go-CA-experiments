package gfx

import "image/color"

type ColorLookup interface {
	GetColor(string) color.RGBA
}
type Palette struct {
	Name         string
	Colors       map[string]color.RGBA
	DefaultColor string
}

func (p *Palette) GetColor(name string) color.RGBA {
	if c, ok := p.Colors[name]; ok {
		return c
	}
	return p.Colors[p.DefaultColor]
}

func NewPalette(name string, colors map[string]color.RGBA, defaultColor string) *Palette {
	if defaultColor == "" {
		var low uint8 = 255
		for k, v := range colors {
			if v.R < low {
				low = v.R
				defaultColor = k
			}
		}
	}
	return &Palette{
		Name:         name,
		Colors:       colors,
		DefaultColor: defaultColor,
	}
}
