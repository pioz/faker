package faker_test

import (
	"testing"
	"time"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestDurationInRange(t *testing.T) {
	faker.SetSeed(41)
	value := faker.DurationInRange(100, 200)
	t.Log(value)
	assert.Equal(t, time.Duration(158), value)
}

func TestDuration(t *testing.T) {
	faker.SetSeed(42)
	value := faker.Duration()
	t.Log(value)
	expectedDuration, _ := time.ParseDuration("-1606331h18m2.623497133s")
	assert.Equal(t, expectedDuration, value)
}

func TestTimeInRange(t *testing.T) {
	faker.SetSeed(43)
	var minTime, maxTime, value time.Time

	minTime, _ = time.Parse(time.RFC3339, "1984-10-21T22:08:41+00:00")
	maxTime, _ = time.Parse(time.RFC3339, "1980-10-21T22:08:41+00:00")
	value = faker.TimeInRange(minTime, maxTime)
	t.Log(value)
	assert.Equal(t, "1984-10-21 22:08:41 +0000 +0000", value.String())

	minTime, _ = time.Parse(time.RFC3339, "1984-10-21T22:08:41+00:00")
	maxTime, _ = time.Parse(time.RFC3339, "2020-10-21T22:08:41+00:00")
	value = faker.TimeInRange(minTime, maxTime)
	t.Log(value)
	assert.Equal(t, "1992-10-10 23:04:44.977812265 +0000 +0000", value.String())
}

func TestTime(t *testing.T) {
	faker.SetSeed(44)
	value := faker.Time()
	t.Log(value)
	assert.NotEmpty(t, value)
}

func TestNanoSecond(t *testing.T) {
	faker.SetSeed(45)
	value := faker.NanoSecond()
	t.Log(value)
	assert.Equal(t, 299642157, value)
}

func TestSecond(t *testing.T) {
	faker.SetSeed(46)
	value := faker.Second()
	t.Log(value)
	assert.Equal(t, 51, value)
}

func TestMinute(t *testing.T) {
	faker.SetSeed(47)
	value := faker.Minute()
	t.Log(value)
	assert.Equal(t, 40, value)
}

func TestHour(t *testing.T) {
	faker.SetSeed(48)
	value := faker.Hour()
	t.Log(value)
	assert.Equal(t, 9, value)
}

func TestDay(t *testing.T) {
	faker.SetSeed(49)
	value := faker.Day()
	t.Log(value)
	assert.Equal(t, 5, value)
}

func TestWeekDay(t *testing.T) {
	faker.SetSeed(50)
	value := faker.WeekDay()
	t.Log(value)
	assert.Equal(t, "Saturday", value)
}

func TestMonth(t *testing.T) {
	faker.SetSeed(51)
	value := faker.Month()
	t.Log(value)
	assert.Equal(t, "July", value)
}

func TestYear(t *testing.T) {
	faker.SetSeed(52)
	value := faker.Year()
	t.Log(value)
	assert.Equal(t, 2032, value)
}

func TestTimeZone(t *testing.T) {
	faker.SetSeed(53)
	value := faker.TimeZone()
	t.Log(value)
	assert.Equal(t, "Venezuela Standard Time", value)
}

func TestTimeZoneAbbr(t *testing.T) {
	faker.SetSeed(54)
	value := faker.TimeZoneAbbr()
	t.Log(value)
	assert.Equal(t, "MEDT", value)
}

func TestTimeZoneFull(t *testing.T) {
	faker.SetSeed(55)
	value := faker.TimeZoneFull()
	t.Log(value)
	assert.Equal(t, "(UTC-03:00) Cayenne, Fortaleza", value)
}

func TestTimeZoneOffset(t *testing.T) {
	faker.SetSeed(56)
	value := faker.TimeZoneOffset()
	t.Log(value)
	assert.Equal(t, float32(-1), value)
}

func TestTimeZoneRegion(t *testing.T) {
	faker.SetSeed(57)
	value := faker.TimeZoneRegion()
	t.Log(value)
	assert.Equal(t, "Africa/Accra", value)
}

func TestTimeBuild(t *testing.T) {
	faker.SetSeed(301)
	s := &struct {
		Field1  time.Duration `faker:"DurationInRange(58m,1h)"`
		Field2  time.Duration `faker:"duration"`
		Field3  time.Duration
		Field4  time.Time `faker:"time"`
		Field5  time.Time
		Field6  int     `faker:"NanoSecond"`
		Field7  int     `faker:"Second"`
		Field8  int     `faker:"Minute"`
		Field9  int     `faker:"Hour"`
		Field10 int     `faker:"Day"`
		Field11 string  `faker:"WeekDay"`
		Field12 string  `faker:"Month"`
		Field13 int     `faker:"Year"`
		Field14 string  `faker:"TimeZone"`
		Field15 string  `faker:"TimeZoneAbbr"`
		Field16 string  `faker:"TimeZoneFull"`
		Field17 float32 `faker:"TimeZoneOffset"`
		Field18 string  `faker:"TimeZoneRegion"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "59m0.054463842s", s.Field1.String())
	assert.Equal(t, "-1578403h18m50.786383245s", s.Field2.String())
	assert.Equal(t, "-2000121h17m20.98416526s", s.Field3.String())
	assert.NotEmpty(t, s.Field4)
	assert.NotEmpty(t, s.Field5)
	assert.Equal(t, 377478642, s.Field6)
	assert.Equal(t, 23, s.Field7)
	assert.Equal(t, 19, s.Field8)
	assert.Equal(t, 11, s.Field9)
	assert.Equal(t, 19, s.Field10)
	assert.Equal(t, "Monday", s.Field11)
	assert.Equal(t, "June", s.Field12)
	assert.Equal(t, 1919, s.Field13)
	assert.Equal(t, "Central Asia Standard Time", s.Field14)
	assert.Equal(t, "NAST", s.Field15)
	assert.Equal(t, "(UTC+02:00) Tripoli", s.Field16)
	assert.Equal(t, float32(-1), s.Field17)
	assert.Equal(t, "Etc/GMT+5", s.Field18)
}
