package faker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleDurationInRange() {
	faker.SetSeed(41)
	fmt.Println(faker.DurationInRange(100, 200))
	// Output: 158ns
}

func ExampleDuration() {
	faker.SetSeed(42)
	fmt.Println(faker.Duration())
	// Output: -1606331h18m2.623497133s
}

func TestTimeInRange(t *testing.T) {
	faker.SetSeed(43)
	var minTime, maxTime, value time.Time

	minTime, _ = time.ParseInLocation(time.RFC3339, "1984-10-21T22:08:41+00:00", time.UTC)
	maxTime, _ = time.ParseInLocation(time.RFC3339, "1980-10-21T22:08:41+00:00", time.UTC)
	value = faker.TimeInRange(minTime, maxTime)
	t.Log(value)
	assert.Equal(t, "1984-10-21 22:08:41 +0000 UTC", value.String())

	minTime, _ = time.ParseInLocation(time.RFC3339, "1984-10-21T22:08:41+00:00", time.UTC)
	maxTime, _ = time.ParseInLocation(time.RFC3339, "2020-10-21T22:08:41+00:00", time.UTC)
	value = faker.TimeInRange(minTime, maxTime)
	t.Log(value)
	assert.Equal(t, "1992-10-10 23:04:44.977812265 +0000 UTC", value.String())
}

func TestTime(t *testing.T) {
	faker.SetSeed(44)
	value := faker.Time()
	t.Log(value)
	assert.NotEmpty(t, value)
}

func ExampleNanoSecond() {
	faker.SetSeed(45)
	fmt.Println(faker.NanoSecond())
	// Output: 299642157
}

func ExampleSecond() {
	faker.SetSeed(46)
	fmt.Println(faker.Second())
	// Output: 51
}

func ExampleMinute() {
	faker.SetSeed(47)
	fmt.Println(faker.Minute())
	// Output: 40
}

func ExampleHour() {
	faker.SetSeed(48)
	fmt.Println(faker.Hour())
	// Output: 9
}

func ExampleDay() {
	faker.SetSeed(49)
	fmt.Println(faker.Day())
	// Output: 5
}

func ExampleWeekDay() {
	faker.SetSeed(50)
	fmt.Println(faker.WeekDay())
	// Output: Saturday
}

func ExampleMonth() {
	faker.SetSeed(51)
	fmt.Println(faker.Month())
	// Output: July
}

func ExampleYear() {
	faker.SetSeed(52)
	fmt.Println(faker.Year())
	// Output: 2032
}

func ExampleTimeZone() {
	faker.SetSeed(53)
	fmt.Println(faker.TimeZone())
	// Output: Venezuela Standard Time
}

func ExampleTimeZoneAbbr() {
	faker.SetSeed(54)
	fmt.Println(faker.TimeZoneAbbr())
	// Output: MEDT
}

func ExampleTimeZoneFull() {
	faker.SetSeed(55)
	fmt.Println(faker.TimeZoneFull())
	// Output: (UTC-03:00) Cayenne, Fortaleza
}

func ExampleTimeZoneOffset() {
	faker.SetSeed(56)
	fmt.Println(faker.TimeZoneOffset())
	// Output: -1
}

func ExampleTimeZoneRegion() {
	faker.SetSeed(57)
	fmt.Println(faker.TimeZoneRegion())
	// Output: Africa/Accra
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
