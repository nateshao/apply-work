package main

import (
	"flag"
	"fmt"
)

var (
	h  string
	a  string
	ip = flag.Int("flagname", 1234, "help message for flagname")
)

func main() {
	fmt.Println("da")
	fmt.Println(ip)
}
