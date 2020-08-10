package faker

import (
	"strconv"
	"time"
)

// DurationInRange will build a random time.Duration between min and max
// included.
func DurationInRange(min, max time.Duration) time.Duration {
	return time.Duration(Int64InRange(int64(min), int64(max)))
}

// Duration will build a random time.Duration.
func Duration() time.Duration {
	return time.Duration(Int64())
}

// TimeInRange will build a random time.Time between min and max included.
func TimeInRange(min, max time.Time) time.Time {
	if min.After(max) {
		return min
	}
	d := DurationInRange(0, max.Sub(min))
	return min.Add(d)
}

// Time will build a random time.Time.
func Time() time.Time {
	return time.Now().Add(Duration())
}

// NanoSecond will build a random nano second.
func NanoSecond() int {
	return IntInRange(0, 999999999)
}

// Second will build a random second.
func Second() int {
	return IntInRange(0, 59)
}

// Minute will build a random minute.
func Minute() int {
	return IntInRange(0, 59)
}

// Hour will build a random hour in military time.
func Hour() int {
	return IntInRange(0, 23)
}

// Day will build a random day between 1 - 31.
func Day() int {
	return IntInRange(1, 31)
}

// WeekDay will build a random weekday string (Monday-Sunday).
func WeekDay() string {
	return time.Weekday(IntInRange(0, 6)).String()
}

// Month will build a random month string.
func Month() string {
	return time.Month(IntInRange(1, 12)).String()
}

// Year will build a random year between 1900 - 2100.
func Year() int {
	return IntInRange(1900, 2100)
}

// TimeZone will build a random timezone string.
func TimeZone() string {
	timezone, err := GetData("timezone", "text")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

// TimeZoneAbbr will build a random abbreviated timezone string.
func TimeZoneAbbr() string {
	timezone, err := GetData("timezone", "abbr")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

// TimeZoneFull will build a random full timezone string.
func TimeZoneFull() string {
	timezone, err := GetData("timezone", "full")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

// TimeZoneOffset will build a random timezone offset.
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

// TimeZoneRegion will build a random timezone region string.
func TimeZoneRegion() string {
	timezone, err := GetData("timezone", "region")
	if err != nil {
		panic(err)
	}
	return timezone.(string)
}

// Builder functions

func durationInRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxDuration(params...)
	if err != nil {
		return nil, err
	}
	return DurationInRange(min, max), nil
}

func durationBuilder(params ...string) (interface{}, error) {
	return Duration(), nil
}

func timeBuilder(params ...string) (interface{}, error) {
	return Time(), nil
}

func nanoSecondBuilder(params ...string) (interface{}, error) {
	return NanoSecond(), nil
}

func secondBuilder(params ...string) (interface{}, error) {
	return Second(), nil
}

func minuteBuilder(params ...string) (interface{}, error) {
	return Minute(), nil
}

func hourBuilder(params ...string) (interface{}, error) {
	return Hour(), nil
}

func dayBuilder(params ...string) (interface{}, error) {
	return Day(), nil
}

func weekDayBuilder(params ...string) (interface{}, error) {
	return WeekDay(), nil
}

func monthBuilder(params ...string) (interface{}, error) {
	return Month(), nil
}

func yearBuilder(params ...string) (interface{}, error) {
	return Year(), nil
}

func timeZoneBuilder(params ...string) (interface{}, error) {
	return TimeZone(), nil
}

func timeZoneAbbrBuilder(params ...string) (interface{}, error) {
	return TimeZoneAbbr(), nil
}

func timeZoneFullBuilder(params ...string) (interface{}, error) {
	return TimeZoneFull(), nil
}

func timeZoneOffsetBuilder(params ...string) (interface{}, error) {
	return TimeZoneOffset(), nil
}

func timeZoneRegionBuilder(params ...string) (interface{}, error) {
	return TimeZoneRegion(), nil
}
