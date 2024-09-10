package network

import (
	"encoding/json"
	"main/cmd/display"
)

type Packages struct {
	Frames []display.Frame `json:"frames"`
}

func (p *Packages) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	jsonDataNewWithNewLine := append(jsonData, '\n')

	return jsonDataNewWithNewLine, err

}
