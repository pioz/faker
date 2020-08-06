package faker

import (
	"strconv"
	"time"
)

func DurationInRange(min, max time.Duration) time.Duration {
	return time.Duration(Int64InRange(int64(min), int64(max)))
}

func Duration() time.Duration {
	return time.Duration(Int64())
}

func TimeInRange(min, max time.Time) time.Time {
	if min.After(max) {
		return min
	}
	d := DurationInRange(0, max.Sub(min))
	return min.Add(d)
}

func Time() time.Time {
	return time.Now().Add(Duration())
}

// NanoSecond will generate a random nano second
func NanoSecond() int {
	return IntInRange(0, 999999999)
}

// Second will generate a random second
func Second() int {
	return IntInRange(0, 59)
}

// Minute will generate a random minute
func Minute() int {
	return IntInRange(0, 59)
}

// Hour will generate a random hour - in military time
func Hour() int {
	return IntInRange(0, 23)
}

// Day will generate a random day between 1 - 31
func Day() int {
	return IntInRange(1, 31)
}

// WeekDay will generate a random weekday string (Monday-Sunday)
func WeekDay() string {
	return time.Weekday(IntInRange(0, 6)).String()
}

// Month will generate a random month string
func Month() string {
	return time.Month(IntInRange(1, 12)).String()
}

// Year will generate a random year between 1900 - 2100
func Year() int {
	return IntInRange(1900, 2100)
}

func TimeZone() string {
	timezone, err := GetData("timezone", "text")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

func TimeZoneAbbr() string {
	timezone, err := GetData("timezone", "abbr")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

func TimeZoneFull() string {
	timezone, err := GetData("timezone", "full")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

func TimeZoneOffset() float32 {
	offsetString, err := GetData("timezone", "offset")
	if err != nil {
		panic(err)
	}
	offset, err := strconv.ParseFloat(offsetString.(string), 32)
	if err != nil {
		panic(err)
	}
	return float32(offset)
}

func TimeZoneRegion() string {
	timezone, err := GetData("timezone", "region")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

// Provider functions

func durationInRangeProvider(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxDuration(params...)
	if err != nil {
		return nil, err
	}
	return DurationInRange(min, max), nil
}

func durationProvider(params ...string) (interface{}, error) {
	return Duration(), nil
}

func timeProvider(params ...string) (interface{}, error) {
	return Time(), nil
}

func nanoSecondProvider(params ...string) (interface{}, error) {
	return NanoSecond(), nil
}

func secondProvider(params ...string) (interface{}, error) {
	return Second(), nil
}

func minuteProvider(params ...string) (interface{}, error) {
	return Minute(), nil
}

func hourProvider(params ...string) (interface{}, error) {
	return Hour(), nil
}

func dayProvider(params ...string) (interface{}, error) {
	return Day(), nil
}

func weekDayProvider(params ...string) (interface{}, error) {
	return WeekDay(), nil
}

func monthProvider(params ...string) (interface{}, error) {
	return Month(), nil
}

func yearProvider(params ...string) (interface{}, error) {
	return Year(), nil
}

func timeZoneProvider(params ...string) (interface{}, error) {
	return TimeZone(), nil
}

func timeZoneAbbrProvider(params ...string) (interface{}, error) {
	return TimeZoneAbbr(), nil
}

func timeZoneFullProvider(params ...string) (interface{}, error) {
	return TimeZoneFull(), nil
}

func timeZoneOffsetProvider(params ...string) (interface{}, error) {
	return TimeZoneOffset(), nil
}

func timeZoneRegionProvider(params ...string) (interface{}, error) {
	return TimeZoneRegion(), nil
}
