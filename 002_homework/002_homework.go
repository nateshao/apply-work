package main

import "fmt"

func main() { // 同一个目录下不能有多个package main
	var oneortwo int
	fmt.Println("输入 1或2")
	fmt.Scanln(&oneortwo)
	if oneortwo == 1 {
		fmt.Println("唐诗：好雨知时节，当春乃发生。")
	} else if oneortwo == 2 {
		fmt.Println("宋词:迟日江山丽，春风花草香。")
	}

}
