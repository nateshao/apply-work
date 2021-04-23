package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	var m Message
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":43142342}`)

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("json Unmarshal error")
	}
	fmt.Println(m.Name, m.Time, m.Body)
	fmt.Println(m.Name)
	fmt.Println(m.Body)
	fmt.Println(m.Time)
}
