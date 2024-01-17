package time

type ITimeDefinitions interface {
	ValidYear(y int) bool
	ValidMonth(m int) bool
	ValidDay(y, m, d int) bool
	ValidHour(h, m, s int) bool
	ValidMinute(m int) bool
	ValidSecond(s int) bool
	ValidIso8601Date(s string) bool
	ValidIso8601Time(s string) bool
	ValidIso8601DateTime(s string) bool
	ValidIso8601Duration(s string) bool
}
