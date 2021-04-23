package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func main() {
	t, _ := time.Parse("2021-04-22 8:04:05", "2021-04-23 12:02:09")

	fmt.Println(t.Hour())
	// fmt.Println(t.Clock())
}
