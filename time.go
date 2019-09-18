package golib

import (
	"time"
	"strings"
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

