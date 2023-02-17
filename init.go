package structmemoryvisualization

import (
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

// Colors
var white = color.RGBA{255, 255, 255, 0xff}
var black = color.RGBA{0, 0, 0, 0xff}
var green = color.RGBA{0, 185, 0, 0xff}
var red = color.RGBA{185, 0, 0, 0xff}

// Font data for texts
var fontData *truetype.Font

// Word size (depends on the cpu architecture. It is 8 for 64bit CPUs and it is 4 for 32bit CPUs)
const wordSize = (32 << uintptr(^uintptr(0)>>63)) / 8

func init() {
	// Load the font
	var err error
	fontData, err = freetype.ParseFont(goregular.TTF)
	if err != nil {
		panic(err)
	}
}
