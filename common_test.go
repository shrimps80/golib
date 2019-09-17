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

func TestStrtotime(t *testing.T) {
	time, _ := Strtotime("2012/12/12 12:12:12")
	fmt.Println(time)
	date := Date("m/d/Y H:i:s", time)
	fmt.Println(date)
}

func TestSubstr(t *testing.T) {
	str := "一二三四五六七八九十"
	fmt.Println(Substr(str, 0, 15))
	fmt.Println(Substr(str, 15, -1))
	fmt.Println(Substr(str, -15, 15))
}

func TestInArray(t *testing.T) {
	peopel := []string{"Bill", "Steve", "Mark", "David"}
	if InArray("Mark", peopel) {
		fmt.Println("in slice")
	} else {
		fmt.Println("not in slice")
	}

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	if InArray("罗马", countryCapitalMap) {
		fmt.Println("in map")
	} else {
		fmt.Println("not in map")
	}

	var number = [5]int{1, 2, 3, 4}
	if InArray(3, number) {
		fmt.Println("in array")
	} else {
		fmt.Println("not in array")
	}
}

func TestStrpos(t *testing.T) {
	fmt.Println(Strpos("You love php, I love php too!", "p1hp", 10))
}
