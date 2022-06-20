package time

import (
	"time"
)

const TIME_STYLE_STR string = "2006-01-02T15:04:05Z"
const ZERO_TIME_STR string = "0000-00-00 00:00:00"
const NULL_TIME_STR string = "0001-01-01 00:00:00"
const DB_DEFAULT_TIME_STR string = "2006-01-02 15:04:05"

// TimeNowCST return now beijing time
func TimeNowCST() (time.Time, error) {
	var curNow time.Time
	timelocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return curNow, err
	}
	time.Local = timelocal
	curNow = time.Now().Local()
	return curNow, nil
}

// TransToCST trans to beijing time
// cstStr is beijing time string like: 2017-06-20 18:16:15
func TransToCST(cstStr string) (time.Time, error) {
	if cstStr == "" {
		return ZEROTime(), nil
	}
	t := time.Time{}
	local, _ := time.LoadLocation("Asia/Chongqing")
	tmp, err := time.ParseInLocation(TIME_STYLE_STR, cstStr, local)
	if err != nil {
		return t, err
	}
	t = tmp
	return t, nil
}

// TransToUTC trans to beijing time
// cstStr is beijing time string like: 2017-06-20 18:16:15
func TransToUTC(utcStr string) (time.Time, error) {
	if utcStr == "" {
		return ZEROTime(), nil
	}
	t := time.Time{}
	local := time.UTC
	tmp, err := time.ParseInLocation(TIME_STYLE_STR, utcStr, local)
	if err != nil {
		return t, err
	}
	t = tmp
	return t, nil
}

// CSTTransToA trans to CST to string
// cst is a beijing time
func CSTTransToA(cst time.Time) string {
	return cst.Format(TIME_STYLE_STR)
}

func ZEROTime() time.Time {
	z, _ := TransToCST(ZERO_TIME_STR)
	return z
}

func IsNULLTime(t time.Time) bool {
	if t.Format(TIME_STYLE_STR) == NULL_TIME_STR {
		return true
	}
	return false
}

func NewDBDefaultTime() time.Time {
	z, _ := TransToCST(TIME_STYLE_STR)
	return z
}

func IsDBDefaultTime(t time.Time) bool {
	if t.Format(TIME_STYLE_STR) == TIME_STYLE_STR {
		return true
	}
	return false
}

func DBDefaultTimeTransZEROTime(t time.Time) time.Time {
	if IsDBDefaultTime(t) {
		return ZEROTime()
	}
	return t
}

func ZEROTimeTransDBDefaultTime(t time.Time) time.Time {
	if IsNULLTime(t) {
		return NewDBDefaultTime()
	}
	return t
}
