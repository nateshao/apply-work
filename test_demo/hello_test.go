package test_demo

import "testing"

func TestGetHello(t *testing.T) { // Wrong test signature :错误的测试签名
	s := GetHello()
	if s != "hello 千羽" {
		t.Errorf("FAILL")
	}

}
