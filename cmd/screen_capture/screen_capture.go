package screen_capture

import (
	"fmt"
	"main/cmd/network"
	"main/cmd/network/client"
	"main/cmd/renderer"
	"time"

	"github.com/vova616/screenshot"
)

func InitRender(clientStruct *client.ClientStruct) {
	fmt.Println("start rendering")
	targetFPS := 30	// currFps := 0
	frameDuration := time.Second / time.Duration(targetFPS)
    currFps := 0

	var packages network.Packages

	for {
		startTime := time.Now()

		elasped := time.Since(startTime)


		sleepTime := frameDuration - elasped
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
        frame, err := CaptureScreen(clientStruct.TargetWidth, clientStruct.TargetHeight)
        packages.Frames = append(packages.Frames, *frame)
        if err != nil {
            err = clientStruct.SendTCPMessage(&packages)
            if err != nil {
                fmt.Println("fail")
                fmt.Println(err)
                continue
            }
            fmt.Println("packetSent")
            packages = network.Packages{}
        }

        if currFps == targetFPS {
            fmt.Println("call capture frame")

            fmt.Println("sending packet")
            err = clientStruct.SendTCPMessage(&packages)
            if err != nil {
                fmt.Println("fail")
                fmt.Println(err)
                continue
            }
            fmt.Println("packetSent")
            packages = network.Packages{}
            currFps = 0
        }
        currFps++


	}

}
func CaptureScreen(targetWidth int32, targetHeight int32) (*renderer.Frame, error) {

	fmt.Println("capture frame")
	img, err := screenshot.CaptureScreen()
	if err != nil {
		return nil, err
	}
	fmt.Println("screen shoted")
	renderImage, err := renderer.ReadImage(img)
	if err != nil {
		return nil, err
	}
	fmt.Println("image read")
	// fmt.Println(renderImage)
	newFrame, err := renderImage.GetFrame(targetWidth, targetHeight)
	if err != nil {
		return nil, err
	}


	return newFrame, nil
}
