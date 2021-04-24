// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	var port = flag.Int("port", 0, "帮助")
// 	flag.Parse()
// 	fmt.Println(port)
// }
// package main
package main

import (
	"flag"
	"fmt"
)

func main() {
	var tangshi int
	var songci int
	flag.IntVar(&tangshi, "tangshi", 1, "好雨知时节，当春乃发生。")
	flag.IntVar(&songci, "songci", 2, "宋词")
	flag.Parse()
	fmt.Println("唐诗：好雨知时节，当春乃发生。")

	// flag.IntVar(&port, "port", 2, "宋词")
	// flag.Parse()
	fmt.Println("宋词")
}
