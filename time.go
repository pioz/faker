package factory

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
	timezone, _ := GetData("timezone", "text")
	return timezone.(string)
}

func TimeZoneAbbr() string {
	timezone, _ := GetData("timezone", "abbr")
	return timezone.(string)
}

func TimeZoneFull() string {
	timezone, _ := GetData("timezone", "full")
	return timezone.(string)
}

func TimeZoneOffset() float32 {
	offsetString, _ := GetData("timezone", "offset")
	offset, _ := strconv.ParseFloat(offsetString.(string), 32)
	return float32(offset)
}

func TimeZoneRegion() string {
	timezone, _ := GetData("timezone", "region")
	return timezone.(string)
}
