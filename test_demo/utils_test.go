package main

import (
	"testing"
	"time"
)

func TestGetDayGap(t *testing.T) {
	type args struct {
		now    time.Time
		before time.Time
	}
	var formatStr = "2006-01-02 15:04:05"
	day1, _:= time.Parse(formatStr, "2019-07-10 11:11:11")
	day2, _:= time.Parse(formatStr, "2019-07-10 22:11:11")
	day3, _:= time.Parse(formatStr, "2019-07-13 22:11:11")
	tests := []struct {
		name string
		args args
		want int
	}{
		{"同一天", args{day1, day2}, 0}, //在这里写入测试的数据
		{"差3天", args{day1, day3}, 3}, // 如果有多个条件，就写多个
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayGap(tt.args.now, tt.args.before); got != tt.want {
				t.Errorf("GetDayGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
