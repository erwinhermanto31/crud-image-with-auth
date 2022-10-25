package utils

import (
	"encoding/base64"
	"strconv"
	"time"
)

const (
	DateFormatRFC3339 = time.RFC3339
)

func FormatDateToRFC3339(t time.Time) string {
	return t.Format(DateFormatRFC3339)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
