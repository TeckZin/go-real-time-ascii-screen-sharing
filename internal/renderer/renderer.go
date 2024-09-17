package renderer

type RendererColor struct {
	R, G, B uint8
}

// 0 r 1 b 2 g 3 ascii

type Pixel struct {
	Values [4]byte `json:"byte"`
}

type Frame struct {
	Pixels [][]*Pixel `json:"pixels"`
}

func (img *RenderImage) GetFrame(width int32, height int32) (*Frame, error) {
	var newFrame Frame
	// newFrame.Pixels = make([][]*Pixel, height)

	// img.ScaleImageRatio(IMG_SCALE)

    img.ScaleImageBounds(width, height)

  	img.Width = int(width)
	img.Height = int(height)
	// fmt.Println("got scale")

	img.GetBrightness()
	// fmt.Println("got brighness")

	img.GetColorMap()
	// fmt.Println("got color map")

	// fmt.Println("image")
	// fmt.Println(img.BrightnessMap)

	newFrame.convertBrightnessToAscii(img.BrightnessMap, width, height)
	// fmt.Println(img.BrightnessMap)
	// fmt.Println("converted to brightness")
	// fmt.Println(newFrame)
	newFrame.getImageColor(img)
	// fmt.Println("got img color")

	return &newFrame, nil

}

func (frame *Frame) convertBrightnessToAscii(brightnessMap [][]int32, width int32, height int32) {

	asciiCharacters := []rune{'.', ':', '-', '=', '+', '*', '#', '%', '$', '@'}
	// fmt.Println(height)
	// fmt.Println(width)
	// fmt.Println(brightnessMap)
	frame.Pixels = make([][]*Pixel, height)

	for y, bRow := range brightnessMap {
		frame.Pixels[y] = make([]*Pixel, width*2) // Initialize each row
		for x, b := range bRow {
			ascii := rune(asciiCharacters[int(b)*(len(asciiCharacters)-1)/255])

			pixel := &Pixel{}
			pixel.Values[3] = byte(ascii)
			frame.Pixels[y][x*2] = pixel
			frame.Pixels[y][x*2+1] = pixel
		}

		// fmt.Println(out)

	}
}

func (frame *Frame) getImageColor(img *RenderImage) {
	for y, row := range frame.Pixels {
		for x, p := range row {
			// fmt.Println("pixles rgb")
			// fmt.Println(*frame.Pixels[y][x])
			// (*p).Values[0] = byte(img.RedColorMap[y][x])
			// (*p).Values[1] = byte(img.BlueColorMap[y][x])
			// (*p).Values[2] = byte(img.GreenColorMap[y][x])

             originalX := x / 2

            (*p).Values[0] = byte(img.RedColorMap[y][originalX])
            (*p).Values[1] = byte(img.BlueColorMap[y][originalX])
            (*p).Values[2] = byte(img.GreenColorMap[y][originalX])
		}
	}

}
