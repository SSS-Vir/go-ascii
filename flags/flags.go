package flags

import (
	"errors"
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"strconv"
	"strings"
)

const ( // resample filters
	NearestNeighbor   = "nearestneighbor"
	Box               = "box"
	Linear            = "linear"
	Hermite           = "hermite"
	MitchellNetravali = "mitchellnetravali"
	CatmullRom        = "catmullrom"
	BSpline           = "bspline"
	Gaussian          = "gaussian"
	Bartlett          = "bartlett"
	Lanczos           = "lanczos"
	Hann              = "hann"
	Hamming           = "hamming"
	Blackman          = "blackman"
	Welch             = "welch"
	Cosine            = "cosine"
)

var resampleFilters = map[string]imaging.ResampleFilter{
	NearestNeighbor:   imaging.NearestNeighbor,
	Box:               imaging.Box,
	Linear:            imaging.Linear,
	Hermite:           imaging.Hermite,
	MitchellNetravali: imaging.MitchellNetravali,
	CatmullRom:        imaging.CatmullRom,
	BSpline:           imaging.BSpline,
	Gaussian:          imaging.Gaussian,
	Bartlett:          imaging.Bartlett,
	Lanczos:           imaging.Lanczos,
	Hann:              imaging.Hann,
	Hamming:           imaging.Hamming,
	Blackman:          imaging.Blackman,
	Welch:             imaging.Welch,
	Cosine:            imaging.Cosine,
}

type ApplicationParameters struct {
	Filename       string
	Width          int
	Height         int
	Fps            uint
	ResampleFilter imaging.ResampleFilter
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
	resampleFilter        string
	applicationParameters ApplicationParameters
)

func init() {
	flag.StringVar(&fileName, "file", "", "path to file in quotes e.g. \".\\path\\to\\file\"")
	flag.StringVar(&size, "size", "", "non negative WIDTHxHEIGHT")
	flag.UintVar(&fps, "fps", 18, "non negative num")
	flag.StringVar(&resampleFilter, "resample", NearestNeighbor, fmt.Sprintf("One of %v", getResampleFiltersKeys()))

	flag.Parse()

	if len(fileName) == 0 {
		println("File name was not given")
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

	applicationParameters.ResampleFilter = resampleFilters[NearestNeighbor]

	val, ok := resampleFilters[strings.ToLower(resampleFilter)]
	if ok {
		applicationParameters.ResampleFilter = val
	}

	resampleFilters = nil
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

func getResampleFiltersKeys() []string {
	var keys []string
	for k := range resampleFilters {
		keys = append(keys, k)
	}
	return keys
}
