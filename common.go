package golib

import (
	"time"
	"strings"
	"math"
)

func Time() int64 {
	return time.Now().Unix()
}

func Date(format string, timestamp int64) string {
	format = strings.ToLower(format)
	changeYears := strings.Replace(format, "y", "2006", -1)
	changeMonths := strings.Replace(changeYears, "m", "01", -1)
	changeDays := strings.Replace(changeMonths, "d", "02", -1)
	changeHours := strings.Replace(changeDays, "h", "15", -1)
	changeMinutes := strings.Replace(changeHours, "i", "04", -1)
	changeSeconds := strings.Replace(changeMinutes, "s", "05", -1)
	return time.Unix(timestamp, 0).Format(changeSeconds)
}

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
