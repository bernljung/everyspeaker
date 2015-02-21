package main

import (
	"bytes"
	"encoding/json"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Speech  string `json:"speech"`
	Link    string `json:"link"`
}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	s = string(b)
	return
}
