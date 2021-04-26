package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func httpGet(url string) (string, error) {
	response, err := http.Get("http://www.weather.com.cn/data/sk/101280601.html")
	if err != nil {
		log.Println("get error")
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("ioutil read error")
	}
	return string(body), err
}

//func httpPost(url string, jsonStr string) (string, error) {
//	var json = []byte(jsonStr)
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
//	req.Header.Set("Content-Type", "application/json")
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("post error")
//	}
//	defer resp.Body.Close()
//	body, err2 := ioutil.ReadAll(resp.Body)
//	if err2 != nil {
//		log.Println("ioutil read error")
//	}
//	return string(body), err
//}
