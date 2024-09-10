package display

type Pixel struct {
	color []byte
	ascii []byte
}

type Frame struct {
	Pixels []Pixel
}
