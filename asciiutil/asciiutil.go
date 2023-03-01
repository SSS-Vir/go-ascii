package asciiutil

import (
	"go-ascii/pixelutil"
	"image"
	"os"
)

func ImageToASCII(image image.Image, colored bool) []string {
	ret := make([]string, image.Bounds().Dy())
	for i := 0; i < image.Bounds().Dy(); i++ {
		for j := 0; j < image.Bounds().Dx(); j++ {
			pixel := pixelutil.NewPixel(image.At(j, i))
			var gradientSymbol string
			if !colored {
				gradientSymbol = pixel.GradientSymbol()
			} else {
				gradientSymbol = pixel.ColoredSymbol()
			}
			ret[i] += gradientSymbol
		}
		ret[i] += "\n"
	}
	return ret
}

func FImageToASCII(image image.Image, file *os.File, colored bool) {
	for i := 0; i < image.Bounds().Dy(); i++ {
		var line string
		for j := 0; j < image.Bounds().Dx(); j++ {
			pixel := pixelutil.NewPixel(image.At(j, i))
			var gradientSymbol string
			if !colored {
				gradientSymbol = pixel.GradientSymbol()
			} else {
				gradientSymbol = pixel.ColoredSymbol()
			}
			line += gradientSymbol
		}
		file.WriteString(line + "\n")
	}
}
