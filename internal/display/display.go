package display

import (
	"fmt"
	"main/internal/renderer"
	"math"
)

type AsciiImage struct {
	PixelMap        [][]*renderer.Pixel
	ANSIEncodingMap [][]*string
}





func rgbToANSI(rgb renderer.RendererColor) string {
	r := uint8(math.Round(float64(rgb.R) / 51))
	g := uint8(math.Round(float64(rgb.G) / 51))
	b := uint8(math.Round(float64(rgb.B) / 51))
	return fmt.Sprintf("\033[38;5;%dm", 16+36*r+6*g+b)
}



// final render
func (a *AsciiImage) GenerateDispay() {

}

func (a *AsciiImage) GetAnsiEncoding() {

	ansiMap := make([][]*string, 0)
	for _, row := range a.PixelMap {
		ansiRow := make([]*string, 0)
		for _, pixel := range row {
			rgb := renderer.RendererColor{R: uint8(pixel.Values[0]), G: uint8(pixel.Values[1]), B: uint8(pixel.Values[2])}
			ansi := rgbToANSI(rgb)
			ansiRow = append(ansiRow, &ansi)
		}
		ansiMap = append(ansiMap, ansiRow)

	}

	a.ANSIEncodingMap = ansiMap

}
