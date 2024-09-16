package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"main/cmd/network"
	"net"
)

type ClientStruct struct {
	Username     string
	ServerUrl    string
	ServerPort   string
	ServerType   string
	TargetWidth  int32
	TargetHeight int32
	Connection   net.Conn
}

func InitClient(username string, serverUrl string,
	serverPort string, serverType string, width int32, height int32) *ClientStruct {
	var clientStruct ClientStruct
	clientStruct.Username = username
	clientStruct.ServerUrl = serverUrl
	clientStruct.ServerPort = serverPort
	clientStruct.ServerType = serverType
	clientStruct.TargetWidth = width
	clientStruct.TargetHeight = height
	return &clientStruct
}

func (c *ClientStruct) SendTCPMessage(packages *network.Packages) error {
	conn, err := net.Dial("tcp", c.ServerUrl+":"+c.ServerPort)
	if err != nil {
		return err
	}
	defer conn.Close()

	jsonData, err := json.Marshal(packages)
	if err != nil {
		return err
	}
    jsonData = append(jsonData, '\n')

	_, err = conn.Write(jsonData)
	if err != nil {
		return err
	}

	err = ResponseTCPMessage(conn)
	if err != nil {
		return err
	}

	return nil
}

func ResponseTCPMessage(conn net.Conn) error {
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return err
	}
	fmt.Printf("Response: %s", resp)
	return nil
}
