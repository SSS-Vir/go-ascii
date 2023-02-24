package main

import (
	"go-ascii/flags"
	internalGif "go-ascii/gif"
	internalJpeg "go-ascii/jpeg"
	internalPng "go-ascii/png"
	"path/filepath"
)

func main() {
	parameters := flags.Get()
	extension := filepath.Ext(parameters.Filepath)
	switch extension {
	case ".gif":
		{
			internalGif.ASCII(parameters)
		}
	case ".jpg", ".jpeg":
		{
			internalJpeg.ASCII(parameters)
		}
	case ".png":
		{
			internalPng.ASCII(parameters)
		}
	}

}
