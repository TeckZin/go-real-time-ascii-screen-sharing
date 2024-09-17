package screen_capture

import (
	"fmt"
	"main/internal/network"
	"main/internal/network/client"
	"main/internal/renderer"
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

        // Capture 30 frames
        for i := 0; i < 30; i++ {
            frame, err := CaptureScreen(clientStruct.Width, clientStruct.Height,
                clientStruct.ServerWidth, clientStruct.Height)
            if err != nil {
                fmt.Println("Error capturing screen:", err)
                continue
            }
            packages.Frames = append(packages.Frames, *frame)

            // Sleep for the remaining time to maintain frame rate
            elapsed := time.Since(startTime)
            sleepTime := frameDuration*time.Duration(i+1) - elapsed
            if sleepTime > 0 {
                time.Sleep(sleepTime)
            }
        }

        // Send the package of 30 frames
        if !sendingInProgress {
            sendingInProgress = true
            go func(p network.Packages) {
                fmt.Println("sending packet of 30 frames")
                err := clientStruct.SendTCPMessage(&p)
                if err != nil {
                    fmt.Println("Failed to send packet:", err)
                }
                sendingInProgress = false
            }(packages)
            packages = network.Packages{}
        }

        // Wait for the next 30-frame cycle
        elapsed := time.Since(startTime)
        sleepTime := frameDuration*30 - elapsed
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
