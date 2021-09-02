package tools

import "time"

func ToDate(timestamp int) string {

	timeobj := time.Unix(int64(timestamp), 0)

	return timeobj.Format("2006-01-02 15:04:05")
}

func Show(s1, s2 string) string {
	return s1 + s2
}
