package main

import (
	"fmt"
	"main/cmd/network/client"
	"main/cmd/network/server"
	"main/cmd/screen_capture"

	"github.com/vova616/screenshot"
)

func main() {

	var input string
	fmt.Print("Host connection[y/n]:")
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		return
	}
	if input == "y" {

		var width int32
		var height int32
		img, err := screenshot.CaptureScreen()
		if err != nil {
			return
		}
		width = int32(img.Bounds().Dx())
		height = int32(img.Bounds().Dy())
		serverStruct := server.InitServerClient(width, height)
		fmt.Println("your hosting")
		serverStruct.StartServer()
	} else if input == "n" {
		var url string
		var port string
		var usr string
		var width int32
		var height int32

		fmt.Print("enter url: ")
		_, err := fmt.Scanf("%s\n", &url)
		if err != nil {
			return
		}

		fmt.Print("enter port: ")
		_, err = fmt.Scanf("%s\n", &port)
		if err != nil {
			return
		}

		fmt.Print("enter username: ")
		_, err = fmt.Scanf("%s\n", &usr)
		if err != nil {
			return
		}

		img, err := screenshot.CaptureScreen()
		if err != nil {
			return
		}
		width = int32(img.Bounds().Dx())
		height = int32(img.Bounds().Dy())

		clientSruct := client.InitClient(usr, url, port, "tcp", width, height)
		screen_capture.InitRender(clientSruct)



	}

}
