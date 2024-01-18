package ISO8601

const (
	SecondsInMinute    = 60
	MinutesInHour      = 60
	HoursInDay         = 24
	AverageDaysInMonth = 30.42
	MaxDaysInMonth     = 31
	DaysInYear         = 365
	AverageDaysInYear  = 365.24
	DaysInLeapYear     = 366
	MaxDaysInYear      = 366
	DaysInWeek         = 7
	MonthsInYear       = 12
	//Minimum hour value of a timezone according to ISO 8601
	//(note that the -ve sign is supplied in the ISO8601_TIMEZONE class).
	MinTimezoneHour = 12
	//Maximum hour value of a timezone according to ISO 8601.
	MaxTimezoneHour = 14
	//Used for conversions of durations containing months to days and / or seconds.
	NominalDaysInMonth = 30.42
	//Used for conversions of durations containing years to days and / or seconds.
	NominalDaysInYear = 365.24
)
