package factory_test

import (
	"testing"
	"time"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestDurationInRange(t *testing.T) {
	factory.SetSeed(41)
	value := factory.DurationInRange(100, 200)
	t.Log(value)
	assert.Equal(t, time.Duration(158), value)
}

func TestDuration(t *testing.T) {
	factory.SetSeed(42)
	value := factory.Duration()
	t.Log(value)
	expectedDuration, _ := time.ParseDuration("-1606331h18m2.623497133s")
	assert.Equal(t, expectedDuration, value)
}

func TestTimeInRange(t *testing.T) {
	factory.SetSeed(43)
	var minTime, maxTime, value time.Time

	minTime, _ = time.Parse(time.RFC3339, "1984-10-21T22:08:41+00:00")
	maxTime, _ = time.Parse(time.RFC3339, "1980-10-21T22:08:41+00:00")
	value = factory.TimeInRange(minTime, maxTime)
	t.Log(value)
	assert.Equal(t, "1984-10-21 22:08:41 +0000 +0000", value.String())

	minTime, _ = time.Parse(time.RFC3339, "1984-10-21T22:08:41+00:00")
	maxTime, _ = time.Parse(time.RFC3339, "2020-10-21T22:08:41+00:00")
	value = factory.TimeInRange(minTime, maxTime)
	t.Log(value)
	assert.Equal(t, "1992-10-10 23:04:44.977812265 +0000 +0000", value.String())
}

func TestTime(t *testing.T) {
	factory.SetSeed(44)
	value := factory.Time()
	t.Log(value)
	assert.NotEmpty(t, value)
}

func TestNanoSecond(t *testing.T) {
	factory.SetSeed(45)
	value := factory.NanoSecond()
	t.Log(value)
	assert.Equal(t, 299642157, value)
}

func TestSecond(t *testing.T) {
	factory.SetSeed(46)
	value := factory.Second()
	t.Log(value)
	assert.Equal(t, 51, value)
}

func TestMinute(t *testing.T) {
	factory.SetSeed(47)
	value := factory.Minute()
	t.Log(value)
	assert.Equal(t, 40, value)
}

func TestHour(t *testing.T) {
	factory.SetSeed(48)
	value := factory.Hour()
	t.Log(value)
	assert.Equal(t, 9, value)
}

func TestDay(t *testing.T) {
	factory.SetSeed(49)
	value := factory.Day()
	t.Log(value)
	assert.Equal(t, 5, value)
}

func TestWeekDay(t *testing.T) {
	factory.SetSeed(50)
	value := factory.WeekDay()
	t.Log(value)
	assert.Equal(t, "Saturday", value)
}

func TestMonth(t *testing.T) {
	factory.SetSeed(51)
	value := factory.Month()
	t.Log(value)
	assert.Equal(t, "July", value)
}

func TestYear(t *testing.T) {
	factory.SetSeed(52)
	value := factory.Year()
	t.Log(value)
	assert.Equal(t, 2032, value)
}

func TestTimeZone(t *testing.T) {
	factory.SetSeed(53)
	value := factory.TimeZone()
	t.Log(value)
	assert.Equal(t, "Venezuela Standard Time", value)
}

func TestTimeZoneAbbr(t *testing.T) {
	factory.SetSeed(54)
	value := factory.TimeZoneAbbr()
	t.Log(value)
	assert.Equal(t, "MEDT", value)
}

func TestTimeZoneFull(t *testing.T) {
	factory.SetSeed(55)
	value := factory.TimeZoneFull()
	t.Log(value)
	assert.Equal(t, "(UTC-03:00) Cayenne, Fortaleza", value)
}

func TestTimeZoneOffset(t *testing.T) {
	factory.SetSeed(56)
	value := factory.TimeZoneOffset()
	t.Log(value)
	assert.Equal(t, float32(-1), value)
}

func TestTimeZoneRegion(t *testing.T) {
	factory.SetSeed(57)
	value := factory.TimeZoneRegion()
	t.Log(value)
	assert.Equal(t, "Africa/Accra", value)
}
