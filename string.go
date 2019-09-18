package golib

import (
	"math"
	"strings"
)

func Substr(str string, start int, length int) string {
	isStrrev := 0
	begin := int(math.Abs(float64(start)))
	if len(str) < begin {
		return str
	}
	if start < 0 {
		isStrrev = 1
		str = Strrev(str)
	}

	switch {
	case length == 0:
		str = ""
		break
	case length == -1:
		str = str[begin:]
		break
	default:
		end := int(begin) + length
		if end > len(str) {
			end = len(str)
		}
		str = str[begin:end]
		break
	}

	if isStrrev == 1 {
		return Strrev(str)
	}
	return str
}

func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}
