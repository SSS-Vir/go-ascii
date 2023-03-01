package png

import (
	"github.com/disintegration/imaging"
	"go-ascii/asciiutil"
	"go-ascii/flags"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ASCII(parameters flags.ApplicationParameters) {
	basename := strings.TrimSuffix(filepath.Base(parameters.Filepath), filepath.Ext(parameters.Filepath))
	imageFile, err := os.Open(parameters.Filepath)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	image, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	if parameters.IsResized() {
		image = imaging.Resize(image, parameters.Width, parameters.Height, parameters.ResampleFilter)
	}
	outfile, err := os.Create("./" + basename + ".txt")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	asciiutil.FImageToASCII(image, outfile, false)
}
