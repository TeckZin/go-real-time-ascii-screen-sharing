package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"main/cmd/display"
	"main/cmd/network"
	"net"
	"strconv"
	"time"
)

type ServerStruct struct {
	URL  string
	Port string
	Type string
    Height int32
    Width int32
}

func InitServerClient(width int32, height int32) *ServerStruct {
	serverStruct := ServerStruct{
		URL:  "localhost",
		Port: "8080",
		Type: "tcp",
        Height: height,
        Width: width,
	}
	return &serverStruct
}

func (s *ServerStruct) StartServer() {
	port, err := strconv.Atoi(s.Port)
	if err != nil {
		fmt.Printf("Error converting port: %v\n", err)
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.URL, port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s:%d\n", s.URL, port)

	for {
		conn, err := listener.Accept()
        fmt.Println("new connection")
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *ServerStruct)handleConnection(conn net.Conn) {
	defer conn.Close()

	packages, err := receivePackages(conn)
	if err != nil {
		fmt.Printf("Error receiving packages: %v\n", err)
		return
	}
    s.DisplayPackage(packages)

	// fmt.Printf("Received packages: %+v\n", packages)

	err = sendResponse(conn)
	if err != nil {
		fmt.Printf("Error sending response: %v\n", err)
	}
}

func receivePackages(conn net.Conn) (*network.Packages, error) {
	reader := bufio.NewReader(conn)
	jsonData, err := reader.ReadBytes('\n')
    fmt.Println("reading")
	if err != nil {
		return nil, err
	}

	var packages network.Packages
	err = json.Unmarshal(jsonData, &packages)
    fmt.Println("unmarshal completed")
	if err != nil {
		return nil, err
	}

	return &packages, nil
}

func sendResponse(conn net.Conn) error {
	_, err := conn.Write([]byte("Package received\n"))
	return err
}

func (server *ServerStruct) DisplayPackage(packages *network.Packages) {

	targetFPS := 30
	// currFps := 0
	frameDuration := time.Second / time.Duration(targetFPS)

    for _, frame := range packages.Frames {

            startTime := time.Now()

            elasped := time.Since(startTime)


            sleepTime := frameDuration - elasped
            if sleepTime > 0 {
                time.Sleep(sleepTime)
            }

            var asciiImage display.AsciiImage
            asciiImage.PixelMap = frame.Pixels

            asciiImage.GetAnsiEncoding()

            for y, pixelRow := range asciiImage.PixelMap {
                for x, pixel := range pixelRow {
                    fmt.Print(string(*asciiImage.ANSIEncodingMap[y][x]))
                    fmt.Print(string(pixel.Values[3]))
                    fmt.Print("\033[0m")

                }
                fmt.Println()
            }
        }

}

