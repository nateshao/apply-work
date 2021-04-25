package main

import (
	"encoding/json"
	"fmt"
)

/*
	json转map
*/

func main() {
	b := []byte(`{"IP":"127.0.0.1","name":"SKY"}`) // json字符串一定是要 用 Esc下面那个 ` 来包（这个符号就是golang中用来包tag的） 然后字符串里面key要双引号
	m := make(map[string]string)
	err := json.Unmarshal(b, &m) // json.Unmarshal不区分json字段 的大小写，只要字母一样即可，不区分大小写
	// 所以大小写不影响反序列化，但是结构体序列化成json的时候大小写必须和tag一致
	if err != nil {
		fmt.Println("Unmarshal failed:", err)
		return
	}
	fmt.Println("m:", m)
	for key, value := range m {
		fmt.Println("key = ", key, "value = ", value)

	}
}
