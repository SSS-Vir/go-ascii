package pixelutil

import (
	"image/color"
	"math"
)

const (
	gradient     = " .:!/r(l1Z4H9W8$@"
	gradientSize = len(gradient)
	gradientStep = uint8(math.MaxUint8 / gradientSize)
)

type Pixel struct {
	R, B, G, A uint8
}

func (p *Pixel) Luma() uint8 {
	return uint8(0.2125*float64(p.R)) + uint8(0.7154*float64(p.G)) + uint8(0.0721*float64(p.B))
}

func (p *Pixel) GradientSymbol() string {
	pixel := p.Luma()
	if pixel == math.MaxUint8 {
		pixel--
	}
	return string(gradient[pixel/gradientStep])
}

func NewPixel(color color.Color) *Pixel {
	var r, g, b, a = color.RGBA()
	return &Pixel{uint8(r), uint8(g), uint8(b), uint8(a)}
}
