package structmemoryvisualization

import (
	"image"
	"image/color"
	"os"
	"reflect"
	"testing"

	"github.com/golang/freetype"
)

func BenchmarkDraw(b *testing.B) {
	testCases := []struct {
		name       string
		structType reflect.Type
	}{
		{
			name: "standard",
			structType: reflect.TypeOf(struct {
				Field1 string
				Field2 int32
				Field3 byte
				Field4 string
				Field5 float64
			}{}),
		},
		{
			name: "empty",
			structType: reflect.TypeOf(struct {
			}{}),
		},
		{
			name: "nested",
			structType: reflect.TypeOf(struct {
				Field1 string
				Field2 struct {
					Field3 int32
					Field4 struct {
						Field5 float64
					}
				}
			}{}),
		},
	}

	for _, testCase := range testCases {
		fieldList, memoryMap := detail(testCase.structType)

		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				draw(fieldList, memoryMap, "test.png")
			}
		})
	}

	os.Remove("test.png")
}

func BenchmarkDrawMemoryBlock(b *testing.B) {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{100, 100},
	})

	c := freetype.NewContext()
	c.SetFont(fontData)
	c.SetDst(img)
	c.SetDPI(72)
	c.SetSrc(image.Black)
	c.SetClip(img.Bounds())
	c.SetFontSize(18)

	b.Run("0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			drawMemoryBlock(img, c, 0, 0, 0)
		}
	})
}

func BenchmarkDrawReact(b *testing.B) {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{100, 100},
	})

	color := color.RGBA{0, 0, 0, 0xff}

	b.Run("0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			drawRect(img, color, 10, 10, 80, 80)
		}
	})
}

func BenchmarkDrawNumber(b *testing.B) {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{100, 100},
	})

	c := freetype.NewContext()
	c.SetFont(fontData)
	c.SetDst(img)
	c.SetDPI(72)
	c.SetSrc(image.Black)
	c.SetClip(img.Bounds())
	c.SetFontSize(18)

	b.Run("0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			drawNumber(c, 0, 0, 0)
		}
	})
}

func BenchmarkDrawText(b *testing.B) {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{100, 100},
	})

	c := freetype.NewContext()
	c.SetFont(fontData)
	c.SetDst(img)
	c.SetDPI(72)
	c.SetSrc(image.Black)
	c.SetClip(img.Bounds())
	c.SetFontSize(18)

	b.Run("0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			drawText(c, 0, 0, "0")
		}
	})
}
