package gif

import (
	"github.com/disintegration/imaging"
	tsize "github.com/kopoli/go-terminal-size"
	"go-ascii/asciiutil"
	"go-ascii/flags"
	"image"
	"image/gif"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ASCII(parameters flags.ApplicationParameters) {
	file, err := os.Open(parameters.Filepath)
	if err != nil {
		panic(err)
	}
	decodedGif, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}

	clearConsole()

	for {
		for frameIndex := 0; frameIndex < len(decodedGif.Image); frameIndex++ {
			start := time.Now()

			var frame image.Image = decodedGif.Image[frameIndex]

			if parameters.IsResized() {
				frame = imaging.Resize(frame, parameters.Width, parameters.Height, parameters.ResampleFilter)
			} else {
				terminalSize, _ := tsize.GetSize()
				frame = imaging.Resize(frame, terminalSize.Width/2, terminalSize.Height, parameters.ResampleFilter)
			}

			ascii := asciiutil.ImageToASCII(frame)
			middleImage(&ascii, &parameters)
			for _, line := range ascii {
				print(line)
			}

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

func middleImage(imageText *[]string, parameters *flags.ApplicationParameters) {
	var count int
	terminalSize, _ := tsize.GetSize()
	if parameters.IsResized() {
		count = (terminalSize.Width - parameters.Width) / 2
	} else {
		count = terminalSize.Width / 4
	}
	spaces := strings.Repeat(" ", count)
	for i, text := range *imageText {
		(*imageText)[i] = spaces + text
	}
}
