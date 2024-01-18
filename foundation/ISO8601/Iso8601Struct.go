package ISO8601

import (
	"github.com/nav-inc/datetime"
	"time"
)

func ValidYear(y int) bool {
	return true
}

func ValidMonth(m int) bool {
	return m >= 1 && m <= MonthsInYear
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func ValidDay(y, m, d int) bool {
	return d >= 1 && d <= daysIn(time.Month(m), y)
}

func ValidHour(h, m, s int) bool {
	return (h >= 0 && h < HoursInDay) || (h == HoursInDay && m == 0 && s == 0)
}

func ValidMinute(m int) bool {
	return m >= 0 && m < MinutesInHour
}

func ValidSecond(s int) bool {
	return s >= 0 && s < SecondsInMinute
}

func ValidFractionalSecond(fs float32) bool {
	return fs >= 0.0 && fs < 1.0
}

func ValidIso8601Date(s string) bool {
	_, err := datetime.ParseLocal(s)
	return err == nil
}

func ValidIso8601Time(s string) bool {
	return ValidIso8601Date(s)
}

func ValidIso8601DateTime(s string) bool {
	return ValidIso8601Date(s)
}

func ValidIso8601Duration(s string) bool {
	_, err := time.ParseDuration(s)
	return err == nil
}
