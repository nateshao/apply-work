package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//	// 发起Get请求
	//	resp, err := http.Get("http://www.weather.com.cn/data/sk/101280601.html")
	//	// 发起Post请求 可以访问https://www.baidu.com
	resp, err := http.Post("http://www.weather.com.cn/data/sk/101280601.html",
		"application/x-www-form-urlencoded", strings.NewReader("name=cjb"))

	if err != nil {
		fmt.Println("http.Get err=", err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body) // 读取
	if err != nil {
		fmt.Println("ioutil.ReadAll err=", err)
		return
	}
	fmt.Println(string(bytes))
}

//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"time"
//)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容

//const url = "https://www.baidu.com"
//func Get(url string) string {
//
//	// 超时时间：5秒
//	client := &http.Client{Timeout: 5 * time.Second}
//	resp, err := client.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	var buffer [512]byte
//	result := bytes.NewBuffer(nil)
//	for {
//		n, err := resp.Body.Read(buffer[0:])
//		result.Write(buffer[0:n])
//		if err != nil && err == io.EOF {
//			break
//		} else if err != nil {
//			panic(err)
//		}
//	}
//
//	return result.String()
//}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
//func Post(url string, data interface{}, contentType string) string {
//
//	// 超时时间：5秒
//	client := &http.Client{Timeout: 5 * time.Second}
//	jsonStr, _ := json.Marshal(data)
//	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	result, _ := ioutil.ReadAll(resp.Body)
//	return string(result)
//}
