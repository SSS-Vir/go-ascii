package asciiutil

import (
	"go-ascii/pixelutil"
	"image"
	"os"
)

func ImageToASCII(image image.Image) []string {
	ret := make([]string, image.Bounds().Dy())
	for i := 0; i < image.Bounds().Dy(); i++ {
		for j := 0; j < image.Bounds().Dx(); j++ {
			pixel := pixelutil.NewPixel(image.At(j, i))
			gradientSymbol := pixel.GradientSymbol()
			ret[i] += gradientSymbol
		}
		ret[i] += "\n"
	}
	return ret
}

func FImageToASCII(image image.Image, file *os.File) {
	for i := 0; i < image.Bounds().Dy(); i++ {
		var line string
		for j := 0; j < image.Bounds().Dx(); j++ {
			pixel := pixelutil.NewPixel(image.At(j, i))
			gradientSymbol := pixel.GradientSymbol()
			line += gradientSymbol
		}
		file.WriteString(line + "\n")
	}
}
