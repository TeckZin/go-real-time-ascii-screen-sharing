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

	"github.com/vova616/screenshot"
)

type ServerStruct struct {
	URL  string
	Port string
	Type string

}

func InitServerClient() *ServerStruct {
	serverStruct := ServerStruct{
		URL:  "localhost",
		Port: "8080",
		Type: "tcp",

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
        // fmt.Println("new connection")
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *ServerStruct)handleConnection(conn net.Conn) {
	defer conn.Close()

    img, err := screenshot.CaptureScreen()
		if err != nil {
			return
		}
    width := int32(img.Bounds().Dx())
    height := int32(img.Bounds().Dy())

	packages, err := receivePackages(conn)
	if err != nil {
		fmt.Printf("Error receiving packages: %v\n", err)
		return
	}
    s.DisplayPackage(packages)

	// fmt.Printf("Received packages: %+v\n", packages)

    message := strconv.Itoa(int(width)) + ":" + strconv.Itoa(int(height))+ "\n"
	err = sendResponse(conn, message)
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

func sendResponse(conn net.Conn, message string ) error {

    _, err :=  conn.Write([]byte(message))
	return err
}

func (server *ServerStruct) DisplayPackage(packages *network.Packages) {

	targetFPS := 30
	frameDuration := time.Second / time.Duration(targetFPS)


    fmt.Print("\033[H")

    for frameIndex, frame := range packages.Frames {
        startTime := time.Now()

        var asciiImage display.AsciiImage
        asciiImage.PixelMap = frame.Pixels
        asciiImage.GetAnsiEncoding()

        // If it's not the first frame, move cursor back to top
        if frameIndex > 0 {
            fmt.Printf("\033[%dA", len(asciiImage.PixelMap))
        }

        for y, pixelRow := range asciiImage.PixelMap {
            // Clear the current line
            fmt.Print("\033[2K")

            for x, pixel := range pixelRow {
                fmt.Print(string(*asciiImage.ANSIEncodingMap[y][x]))
                fmt.Print(string(pixel.Values[3]))
                fmt.Print("\033[0m")
            }

            // Move to the next line
            if y < len(asciiImage.PixelMap)-1 {
                fmt.Print("\r\n")
            }
        }

        elapsed := time.Since(startTime)
        sleepTime := frameDuration - elapsed
        if sleepTime > 0 {
            time.Sleep(sleepTime)
        }
    }
}

