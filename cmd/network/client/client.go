package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"main/cmd/network"
	"net"
	"strconv"
	"strings"
)

type ClientStruct struct {
	Username     string
	ServerUrl    string
	ServerPort   string
	ServerType   string
	Width  int32
	Height int32
    ServerWidth int32
    ServerHeight int32
	Connection   net.Conn
}

func InitClient(username string, serverUrl string,
	serverPort string, serverType string, width int32, height int32) *ClientStruct {
	var clientStruct ClientStruct
	clientStruct.Username = username
	clientStruct.ServerUrl = serverUrl
	clientStruct.ServerPort = serverPort
	clientStruct.ServerType = serverType
	clientStruct.Width = width
	clientStruct.Height = height
    clientStruct.ServerWidth = int32(0)
    clientStruct.ServerHeight = int32(0)

	return &clientStruct
}

func (c *ClientStruct) SendTCPMessage(packages *network.Packages)  error {
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
    width, height, err := c.ResponseTCPMessage(conn)

	if err != nil {
		return err
	}

    c.ServerWidth = width

    c.ServerHeight = height



	return nil
}


func (c *ClientStruct)ResponseTCPMessage(conn net.Conn) (int32, int32, error) {
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return 0,0, err
	}

    resp = resp[:len(resp)-1]
    res := strings.Split(resp, ":")

	fmt.Printf("Response: %s\n",resp)
    if len(res) == 2 {
        width, err := strconv.Atoi(res[0])
        if err != nil {
            return 0, 0, err
        }
        height, err := strconv.Atoi(res[1])
        if err != nil {
            return 0, 0, err
        }
        return int32(width), int32(height), nil

    }

	return 0, 0, nil
}
