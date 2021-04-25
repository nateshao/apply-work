package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.String返回的是指针
	// word表示参数名，foo表示默认值，a string表示参数释义（在-h或解析异常时候会看到）
	// flag.Int、flag.Bool与flag.String同理，不再赘述
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// flag.StringVar返回的是非指针
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 进行flag解析
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	// 002_homework_help.exe -h
}
