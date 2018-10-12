package utils

import (
    "fmt"
    "time"
    "strings"
)

// DateFormat for format time.Time struct to string
func DateFormat(t time.Time, format string) string {
    var result string

    result = strings.Replace(format, "MM", t.Format("01"), -1)
    result = strings.Replace(result, "M", t.Format("1"), -1)

    result = strings.Replace(result, "DD", t.Format("02"), -1)
    result = strings.Replace(result, "D", t.Format("2"), -1)

    result = strings.Replace(result, "YYYY", t.Format("2006"), -1)
    result = strings.Replace(result, "YY", t.Format("06"), -1)

    result = strings.Replace(result, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
    result = strings.Replace(result, "H", fmt.Sprintf("%d", t.Hour()), -1)

    result = strings.Replace(result, "hh", t.Format("03"), -1)
    result = strings.Replace(result, "h", t.Format("3"), -1)

    result = strings.Replace(result, "mm", t.Format("04"), -1)
    result = strings.Replace(result, "m", t.Format("4"), -1)

    result = strings.Replace(result, "ss", t.Format("05"), -1)
    result = strings.Replace(result, "s", t.Format("5"), -1)

    return result
}

// Timestamp to return base unix timestamp
func Timestamp() int64 {
    return time.Now().Unix()
}
