package main

import (
	"fmt"
	"main/cmd/network/client"
	"main/cmd/network/server"
)

func main() {

	var input string
	fmt.Print("Host connection[y/n]:")
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		return
	}
	if input == "y" {
		serverStruct := server.InitServerClient()
		fmt.Println("your hosting")
		serverStruct.StartServer()
	} else if input == "n" {
		var url string
		var port string
		var usr string

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

		clientSruct := client.InitClient(usr, url, port, "tcp")
		clientSruct.ConnectTcp()

	}

}
