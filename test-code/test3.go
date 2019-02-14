package main

import (
	"encoding/json"
	"fmt"
)

// the struct for the value of a "sendMsg"-command
type sendMsg struct {
	user string
	msg  string
}

// The type for the value of a "say"-command
type say string

func main() {
	data := []byte(`{"sendMsg":{"user":"ANisus","msg":"Trying to send a message"},"say":"Hello"}`)

	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	fmt.Printf("%v", objmap)
}
