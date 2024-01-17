package response

import "encoding/json"

type ReadLine struct {
	Line string `json:"line"`
}

func (u *ReadLine) Marshal() ([]byte, error) {
	return json.Marshal(u)
}
