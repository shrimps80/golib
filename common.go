package golib

import (
	"time"
	"strings"
	"math"
	"regexp"
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

func Strtotime(strtime string) (int64, error) {
	re, _ := regexp.Compile("[0-9]+");
	allIndex := re.FindAllIndex([]byte(strtime), -1);
	timeFormart := []string{"01", "02", "15", "04", "05"}
	newStr := strtime
	i := 0
	for k := range allIndex {
		if allIndex[k][1]-allIndex[k][0] == 4 {
			if allIndex[k][0] == 0 {
				newStr = "2006" + newStr[allIndex[k][1]:]
			} else {
				newStr = newStr[:allIndex[k][0]] + "2006" + newStr[allIndex[k][1]:]
			}
		}

		if allIndex[k][1]-allIndex[k][0] == 2 {
			if allIndex[k][0] == 0 {
				newStr = timeFormart[i] + newStr[allIndex[k][1]:]
			} else {
				newStr = newStr[:allIndex[k][0]] + timeFormart[i] + newStr[allIndex[k][1]:]
			}
			i++
		}
	}
	t, err := time.Parse(newStr, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
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
