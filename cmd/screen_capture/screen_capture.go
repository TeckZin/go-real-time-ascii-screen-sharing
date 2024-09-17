package screen_capture

import (
	"fmt"
	"main/cmd/network"
	"main/cmd/network/client"
	"main/cmd/renderer"
	"math"
	"time"

	"github.com/vova616/screenshot"
)
const IMG_RATIO float64 = 50.0


func InitRender(clientStruct *client.ClientStruct) {
	fmt.Println("start rendering")
	targetFPS := 30
	frameDuration := time.Second / time.Duration(targetFPS)
	var packages network.Packages
	sendingInProgress := false
	// lastSendTime := time.Now()
	// minimumWaitTime := time.Second

	for {
		startTime := time.Now()

		frame, err := CaptureScreen(clientStruct.Width, clientStruct.Height,
			clientStruct.ServerWidth, clientStruct.Height)
		packages.Frames = append(packages.Frames, *frame)

		if err != nil {
			fmt.Println("Error capturing screen:", err)
		}

		if !sendingInProgress {
			sendingInProgress = true
			go func(p network.Packages) {
				fmt.Println("sending packet")
				err := clientStruct.SendTCPMessage(&p)
				if err != nil {
					fmt.Println("Failed to send packet:", err)
				}
				sendingInProgress = false
				// lastSendTime = time.Now()
			}(packages)
			packages = network.Packages{}
		}

		elapsed := time.Since(startTime)
		sleepTime := frameDuration - elapsed
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}

func  CaptureScreen(targetWidth int32, targetHeight int32, serverWidth int32,
    serverHeight int32) (*renderer.Frame, error) {

	// fmt.Println("capture frame")
	img, err := screenshot.CaptureScreen()
	if err != nil {
		return nil, err
	}
	// fmt.Println("screen shoted")
	renderImage, err := renderer.ReadImage(img)
	if err != nil {
		return nil, err
	}
	// fmt.Println("image read")
	// fmt.Println(renderImage)
    var tWidth int32
    var tHeight int32
    var newFrame *renderer.Frame

    if serverWidth != 0 && serverHeight != 0 {
        tWidth = int32(float32(serverWidth)/35)
        tHeight = int32(float32(serverHeight)/35)
    } else {
        tWidth  = int32(math.Round(float64(targetWidth)/IMG_RATIO))
        tHeight  = int32(math.Round(float64(targetHeight)/IMG_RATIO))
    }

    newFrame, err = renderImage.GetFrame(tWidth, tHeight)
	if err != nil {
		return nil, err
	}


	return newFrame, nil
}
