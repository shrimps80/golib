package golib

import (
	"testing"
	"fmt"
)

func TestTime(t *testing.T) {
	time := Time()
	fmt.Println(time)
}

func TestDate(t *testing.T) {
	time := Time()
	date := Date("Y-m-d H:i:s", time)
	fmt.Println(date)
}

func TestSubstr(t *testing.T) {
	str := "一二三四五六七八九十"
	fmt.Println(Substr(str, 0, 15))
	fmt.Println(Substr(str, 15, -1))
	fmt.Println(Substr(str, -15, 15))
}
