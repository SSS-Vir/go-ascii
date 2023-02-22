package flags

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type ApplicationParameters struct {
	Filename string
	Width    int
	Height   int
	Fps      uint
}

func (p *ApplicationParameters) IsResized() bool {
	return p.Width != 0 || p.Height != 0
}

func (p *ApplicationParameters) TimeForFrame() int64 {
	return int64(1000 / p.Fps)
}

var (
	fileName              string
	size                  string
	fps                   uint
	applicationParameters ApplicationParameters
)

func init() {
	flag.StringVar(&fileName, "file", "", "path to file in quotes e.g. \".\\path\\to\\file\"")
	flag.StringVar(&size, "size", "", "non negative WIDTHxHEIGHT")
	flag.UintVar(&fps, "fps", 18, "non negative num")

	flag.Parse()

	if !fs.ValidPath(fileName) || len(fileName) == 0 {
		fmt.Printf("Not valid file path, %s", fileName)
		os.Exit(0)
	}
	applicationParameters.Filename = fileName
	width, height, err := parseSize(size)
	if err == nil {
		// width and height can be 0
		applicationParameters.Width = width
		applicationParameters.Height = height
	}
	applicationParameters.Fps = fps

	return
}

func parseSize(size string) (int, int, error) {
	split := strings.Split(size, "x")
	if len(split) != 2 {
		return 0, 0, errors.New("failed to parse size")
	}
	width, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}
	height, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}
	if width < 0 || height < 0 {
		return 0, 0, errors.New("negative nums on size")
	}
	return width, height, nil
}

func Get() ApplicationParameters {
	return applicationParameters
}
