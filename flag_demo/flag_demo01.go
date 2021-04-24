// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	var args = os.Args
// 	fmt.Println(args)

// 	return
// }
package main

import (
	"flag"
	"fmt"
)

func main() {
	var port = flag.Int("port", 1, "帮助")
	flag.Parse()
	fmt.Println(port)
}
