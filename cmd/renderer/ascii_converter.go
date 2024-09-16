package renderer

import (
	"fmt"
	"image"

	"golang.org/x/image/draw"
)

type RenderImage struct {
	Width         int
	Height        int
	BrightnessMap [][]int32
	RedColorMap   [][]int32
	BlueColorMap  [][]int32
	GreenColorMap [][]int32
	ImageValue    *image.Image
}

func ReadImage(image image.Image) (*RenderImage, error) {
	var img RenderImage

	img.ImageValue = &image

	bounds := image.Bounds()

	img.Height = bounds.Max.Y
	img.Width = bounds.Max.X
	fmt.Println("image set")

	return &img, nil

}

func (img *RenderImage) ScaleImageRatio(ratio float32) {
	newWidth := int32(float32(img.Width) * ratio / 2)
	newHeight := int32(float32(img.Height) * ratio / 3)

	img.ScaleImageBounds(newWidth, newHeight)

}
func (img *RenderImage) ScaleImageBounds(width int32, height int32) {
	img.Width = int(width)
	img.Height = int(height)
	emptyImage := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	draw.ApproxBiLinear.Scale(emptyImage, emptyImage.Bounds(), *img.ImageValue, (*img.ImageValue).Bounds(), draw.Over, nil)

	newImg := image.Image(emptyImage)
	img.ImageValue = &newImg

}

func (img *RenderImage) GetColorMap() {
	redMap := make([][]int32, 0)
	greenMap := make([][]int32, 0)
	blueMap := make([][]int32, 0)
	for y := 0; y < img.Height; y++ {
		redRow := make([]int32, 0)
		greenRow := make([]int32, 0)
		blueRow := make([]int32, 0)
		for x := 0; x < img.Width; x++ {
			cords := (*img.ImageValue).At(x, y)
			r, g, b, _ := cords.RGBA()
			redRow = append(redRow, int32(r))
			greenRow = append(greenRow, int32(g))
			blueRow = append(blueRow, int32(b))

		}
		redMap = append(redMap, redRow)
		greenMap = append(greenMap, greenRow)
		blueMap = append(blueMap, blueRow)
	}

	img.RedColorMap = redMap
	img.GreenColorMap = greenMap
	img.BlueColorMap = blueMap

}

func (img *RenderImage) GetBrightness() {
	newBrightnessMap := make([][]int32, 0)
	for y := 0; y < img.Height; y++ {
		row := make([]int32, 0)
		for x := 0; x < img.Width; x++ {
			r, g, b, _ := (*img.ImageValue).At(x, y).RGBA()
			brightness := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 257
			row = append(row, int32(brightness))
		}
		newBrightnessMap = append(newBrightnessMap, row)
	}
	img.BrightnessMap = newBrightnessMap

}
