package main

import (
	"flag"
	"fmt"
)

// func main() {
// 	// Program Name is always the first (implicit) argument
// 	// cmd := os.Args[0]

// 	// fmt.Printf("Program Name: %s\n", cmd)
// 	fmt.Print("hello ")
// }
var (
	v bool
	p string
)
var ip = flag.Int("flagname", 1234, "help message for flagname")

func main() {
	// 绑定参数
	// flag.StringVar(&p, "p", "8080", "port")
	// flag.Parse()
	// // 打印读取的参数n
	// fmt.Println(p)

	// flag.String(p, "qianu", "qianu ooo")
	// flag.Parse()
	// fmt.Println(p)
	fmt.Println(ip)
}
