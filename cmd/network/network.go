package network

import (
	"encoding/json"
	"fmt"
	"main/cmd/renderer"
)

type Packages struct {
	Frames []renderer.Frame `json:"frames"`
}

func (p *Packages) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	// jsonDataNewWithNewLine := append(jsonData, '\n')
	fmt.Printf("jsonData: %s\n", string(jsonData))

	return jsonData, err

}
