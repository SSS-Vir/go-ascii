package main

import (
	"ascii-go/flags"
	"ascii-go/pixelutil"
	"github.com/disintegration/imaging"
	tsize "github.com/kopoli/go-terminal-size"
	"image"
	"image/gif"
	"os"
	"os/exec"
	"time"
)

func main() {
	parameters := flags.Get()

	file, err := os.Open(parameters.Filename)
	if err != nil {
		panic(err)
	}
	decodedGif, err := gif.DecodeAll(file)
	if err != nil {
		return
	}

	clearConsole()

	for true {
		for frameIndex := 0; frameIndex < len(decodedGif.Image); frameIndex++ {

			var final string
			start := time.Now()

			var frame image.Image = decodedGif.Image[frameIndex]

			if parameters.IsResized() {
				frame = imaging.Resize(frame, parameters.Width, parameters.Height, imaging.NearestNeighbor)
			} else {
				terminalSize, _ := tsize.GetSize()
				frame = imaging.Resize(frame, terminalSize.Width/2, terminalSize.Height, imaging.MitchellNetravali)
			}

			for i := 0; i < frame.Bounds().Dy(); i++ {
				middleImage(&final, &parameters)
				for j := 0; j < frame.Bounds().Dx(); j++ {
					pixel := pixelutil.NewPixel(frame.At(j, i))
					gradientSymbol := pixel.GradientSymbol()
					final += gradientSymbol
				}
				final += "\n"
			}
			println(final)
			end := time.Now().Sub(start).Milliseconds()
			if end < parameters.TimeForFrame() {
				time.Sleep(time.Duration(parameters.TimeForFrame()-end) * time.Millisecond)
			}

		}
	}

}

func clearConsole() { //Windows
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func middleImage(final *string, parameters *flags.ApplicationParameters) {
	var count int
	terminalSize, _ := tsize.GetSize()
	if parameters.IsResized() {
		count = (terminalSize.Width - parameters.Width) / 2
	} else {
		count = terminalSize.Width / 4
	}
	for t := 0; t < count; t++ {
		*final += " "
	}
}
