package main

import (
	"flag"
	"fmt"
)

func main() {
	var tangshi int
	// var songci int
	flag.IntVar(&tangshi, "tangshi", 1, "好雨知时节，当春乃发生。")
	// flag.IntVar(&songci, "songci", 2, "宋词:迟日江山丽，春风花草香。")
	flag.Parse()
	fmt.Println("唐诗：好雨知时节，当春乃发生。")

	// flag.IntVar(&port, "port", 2, "宋词")
	// flag.Parse()
	// fmt.Println("宋词:迟日江山丽，春风花草香。", songci)
}
