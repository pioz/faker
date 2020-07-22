package factory_test

import (
	"testing"
	"time"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestTimeBuild(t *testing.T) {
	factory.SetSeed(301)
	s := &struct {
		DurationInRangeField time.Duration `factory:"DurationInRange(58m,1h)"`
		DurationField        time.Duration `factory:"duration"`
		DefaultDurationField time.Duration
		TimeField            time.Time `factory:"time"`
		DefaultTimeField     time.Time
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "59m0.054463842s", s.DurationInRangeField.String())
	assert.Equal(t, "-1578403h18m50.786383245s", s.DurationField.String())
	assert.Equal(t, "-2000121h17m20.98416526s", s.DefaultDurationField.String())
	assert.NotEmpty(t, s.TimeField)
	assert.NotEmpty(t, s.DefaultTimeField)
}
