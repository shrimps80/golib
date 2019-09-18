package golib

import (
	"math"
	"strings"
	"strconv"
	"time"
	"math/rand"
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

func ParseInt(b string, defInt int) int {
	id, err := strconv.Atoi(b)
	if err != nil {
		return defInt
	} else {
		return id
	}
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
