package main

import (
	"flag"
	"fmt"
)

func main() { // 同一个目录下不能有多个package main
	var h bool
	flag.BoolVar(&h, "h", false, "this help")
	flag.Parse()
	fmt.Println(h)
}
