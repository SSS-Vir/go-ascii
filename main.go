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

	file, _ := os.Open(parameters.Filename)
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
				frame = imaging.Resize(frame, terminalSize.Width/2, terminalSize.Height, imaging.NearestNeighbor)
			}

			for i := 0; i < frame.Bounds().Dy(); i++ {
				middleImage(&final)
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

func middleImage(final *string) {
	terminalSize, _ := tsize.GetSize()
	for t := 0; t < terminalSize.Width/4; t++ {
		*final += " "
	}
}
