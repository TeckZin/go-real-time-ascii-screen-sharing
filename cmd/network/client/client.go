package client

import (
	"bufio"
	"fmt"
	"main/cmd/network"
	"net"
)

type ClientSturct struct {
	Username   string
	ServerUrl  string
	ServerPort string
	ServerType string
	Connection net.Conn
}

func InitClient(username string, serverUrl string,
	serverPort string, serverType string) *ClientSturct {
	var clientStruct ClientSturct
	clientStruct.Username = username
	clientStruct.ServerUrl = serverUrl
	clientStruct.ServerPort = serverPort
	clientStruct.ServerType = serverType

	return &clientStruct
}

func (c *ClientSturct) ConnectTcp() {
	// Connect to TCP server
	conn, err := net.Dial(c.ServerType, c.ServerUrl+":"+c.ServerPort)
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		return
	}
	c.Connection = conn
	defer conn.Close()

	fmt.Println("Connected to TCP server. Type your message and press Enter.")

	// Start a goroutine to read responses from the server
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("Received from server: %s\n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading from server: %v\n", err)
		}
	}()

}

func (c *ClientSturct) SendMessage(packages *network.Packages) {
	// Read user input and send to server

	// scanner := bufio.NewScanner(os.Stdin)
	// if !scanner.Scan() {
	// 	fmt.Println("nothing")
	// 	return
	// }

	jsonData, err := packages.ToJson()
	if err != nil {
		return
	}
	c.Connection.Write(jsonData)
	// _, err := fmt.Fprintln(c.Connection, message)
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
		return
	}
	// if err := scanner.Err(); err != nil {
	// 	fmt.Printf("Error reading user input: %v\n", err)
	// }
}
