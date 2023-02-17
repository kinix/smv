package structmemoryvisualization

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/math/fixed"
)

// Draw the image for the given field list and memory map
func draw(fieldList []string, memoryMap []int, fileName string) {
	// Spaces + size of blocks
	width := 200 + 34*wordSize

	// Spaces + size of blocks
	height := 100 + 34*int(math.Ceil(float64(len(memoryMap))/wordSize))

	// If field list is longer than image height, use field list height as image height
	if 50+len(fieldList)*28 > height {
		height = 50 + len(fieldList)*28
	}

	// Create image
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Create text writer for the image
	c := freetype.NewContext()
	c.SetFont(fontData)
	c.SetDst(img)
	c.SetDPI(72)
	c.SetSrc(image.Black)
	c.SetClip(img.Bounds())
	c.SetFontSize(18)

	// White background
	drawRect(img, white, 0, 0, width, height)

	// Draw blocks
	for i, val := range memoryMap {
		x := i % wordSize
		y := int(math.Floor(float64(i) / wordSize))
		drawMemoryBlock(img, c, 51+x*34, 21+y*34, val)
	}

	// Draw field list
	for i, val := range fieldList {
		drawText(c, 100+34*wordSize, 42+i*28, fmt.Sprintf("%d: %s", i, val))
	}

	// Encode as PNG.
	f, _ := os.Create(fileName)
	png.Encode(f, img)
}

// Draw a single block
func drawMemoryBlock(img *image.RGBA, writer *freetype.Context, x0, y0, val int) {
	// Outline
	drawRect(img, black, x0, y0, 32, 32)

	if val == -1 {
		// The block is empty (draw red rectangle)
		drawRect(img, red, x0+2, y0+2, 28, 28)
	} else {
		// The block is used (draw green rectangle)
		drawRect(img, green, x0+2, y0+2, 28, 28)

		// Draw the field id in the rectangle
		drawNumber(writer, x0, y0, val)
	}
}

// Draw a rectangle
func drawRect(img *image.RGBA, clr color.Color, x0, y0, width, height int) {
	for x := x0; x < x0+width; x++ {
		for y := y0; y < y0+height; y++ {
			// Set color for a single pixel.
			img.Set(x, y, clr)
		}
	}
}

// Draw a number
func drawNumber(c *freetype.Context, x0, y0 int, num int) {
	var pt fixed.Point26_6

	if num > 9 { // Single digit
		pt = freetype.Pt(x0+7, y0+22)
	} else { // Multiple digit
		pt = freetype.Pt(x0+11, y0+22)
	}

	c.DrawString(fmt.Sprintf("%d", num), pt)
}

// Draw a text
func drawText(c *freetype.Context, x0, y0 int, text string) {
	pt := freetype.Pt(x0, y0)
	c.DrawString(text, pt)
}
