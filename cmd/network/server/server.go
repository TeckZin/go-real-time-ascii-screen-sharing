package server

import (
	"bufio"
	"fmt"
	"main/cmd/display"
	"main/cmd/network"
	"net"
)

type ServerStuct struct {
	URL  string
	Port string
	Type string
}

func InitServerClient() *ServerStuct {
	serverStuct := ServerStuct{
		URL:  "localhost",
		Port: "8080",
		Type: "tcp",
	}

	return &serverStuct
}

func (s *ServerStuct) StartServer() {
	// Listen on TCP port 8080
	listener, err := net.Listen(s.Type, ":"+s.Port)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(con net.Conn) {
	defer con.Close()
	sc := bufio.NewScanner(con)
	for sc.Scan() {
		message := sc.Bytes()
		if len(message) == 0 {
			break
		}
		pkg := network.Packages{
			Frames: make([]display.Frame, len(message)),
		}
		copy(pkg.Frames, message)
		fmt.Printf("Received message from %s: %s\n", con.RemoteAddr().String(), message)

		// Echo the message back
		_, err := con.Write([]byte("Server: " + string(message) + "\n"))
		if err != nil {
			fmt.Printf("Error sending response: %v\n", err)
			return
		}
	}

	err := sc.Err()
	if err != nil {
		return
	}

}
