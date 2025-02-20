package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法测试未通过")
	} else {
		t.Log("测试一通过！")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("xixi不让你通过")
}
