# Go Real-Time ASCII Screen Sharing

This project implements real-time ASCII screen sharing using Go. It captures screen content, converts it to ASCII format, and shares it over a network in real-time.

## Project Structure

```
.
├── cmd
│   └── main
│       └── main.go
├── go.mod
├── go.sum
└── internal
    ├── display
    │   └── display.go
    ├── network
    │   ├── client
    │   │   └── client.go
    │   ├── network.go
    │   └── server
    │       └── server.go
    ├── renderer
    │   ├── ascii_converter.go
    │   └── renderer.go
    └── screen_capture
        └── screen_capture.go
```

## Features

- Real-time screen capture
- ASCII conversion of screen content
- Network-based sharing with client-server architecture
- Display handling
- Rendering capabilities, including ASCII conversion

## Key Data Structures

The project uses the following main data structures for representing screen content:

```go
type Pixel struct {
    Values [4]byte `json:"byte"`
}

type Frame struct {
    Pixels [][]*Pixel `json:"pixels"`
}
```

- `Pixel`: Represents a single pixel with 4 byte values (likely RGBA).
- `Frame`: Represents a full frame as a 2D slice of pixels.

These structures are designed to be easily serializable to JSON for network transmission.
Note: In the future I will not be using Json to reduce lag and latency! 

## Installation

To install the project, follow these steps:

```bash
git clone https://github.com/TeckZin/go-real-time-ascii-screen-sharing.git
cd go-real-time-ascii-screen-sharing
go mod download
```

## Usage

To run the project:

```bash
go run cmd/main/main.go
```

Detailed usage instructions will be added as the project develops.

## Network Implementation

The project currently uses TCP for network communication. We plan to transition to UDP in the future to:

- Reduce latency for real-time screen sharing
- Improve performance for applications that can tolerate some packet loss
- Enhance scalability for multiple concurrent users

## Contributing

Contributions are welcome! Feel free to submit a Pull Request or create an Issue to discuss potential changes or additions.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Project Status

This project is currently under active development. Features and documentation will be updated regularly.

