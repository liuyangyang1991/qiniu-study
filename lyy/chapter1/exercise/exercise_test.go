package chapter1

import "testing"

func Test_Exercise_1(t *testing.T) {
	if i, e := Cheng(2, 3); i != 6 || e != nil {
		t.Error("出错了!")
	} else {
		t.Log("OK")
	}
}
func Test_Exercise_2(t *testing.T) {
	if i, e := Cheng(0, 3); i != 0 || e != nil {
		t.Log("OK")
	} else {
		t.Error("出错了!")
	}
}
