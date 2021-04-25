package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherInfoJson struct { // 天气信息json
	Weatherinfo WeatherinfoObject
}

type WeatherinfoObject struct { // 天气包含的内容
	City    string // 城市
	CityId  string // 城市id
	Temp    string // 温度
	WD      string // 西部数据（Western Digital）
	WS      string // 气象站（weather station）
	SD      string // 标准偏差（standard deviation）
	Time    string // 时间
	IsRadar string // 是雷达
	Radar   string // 雷达
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)                              // 设置提示信息，及发生的时间和地点
	resp, err := http.Get("http://www.weather.com.cn/data/sk/101280601.html") // 调用HTTP发起Get请求
	if err != nil {
		log.Fatal(err) // 打印日志
	}

	defer resp.Body.Close()                 // 请求关闭
	input, err := ioutil.ReadAll(resp.Body) // 读取所有信息

	var jsonWeather WeatherInfoJson
	json.Unmarshal(input, &jsonWeather) // json转换
	log.Printf("Results:%v\n", jsonWeather)

	log.Println(jsonWeather.Weatherinfo.City)
	log.Println(jsonWeather.Weatherinfo.WD)
	log.Println(jsonWeather.Weatherinfo.IsRadar)
	log.Println(jsonWeather.Weatherinfo.SD)
	log.Println(jsonWeather.Weatherinfo.Temp)
	log.Println(jsonWeather.Weatherinfo.Radar)
	log.Println(jsonWeather.Weatherinfo.Time)

}
