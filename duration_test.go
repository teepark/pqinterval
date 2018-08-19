package pqinterval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationValue(t *testing.T) {
	i := new(Duration)
	_ = i.Scan("3 years 182 days 01:22:33.456789")

	val, err := i.Value()
	assert.Nil(t, err, "Duration.Value() error")
	assert.EqualValues(
		t,
		"3 years 182 days 1 hours 22 minutes 33 seconds 456 milliseconds 789 microseconds",
		val,
		"Duration.Value() result")

	j := time.Duration(30) * time.Minute
	k := Duration(j)
	val, err = k.Value()
	assert.Nil(t, err, "Duration.Value() error")
	assert.EqualValues(
		t,
		"30 minutes",
		val,
		"Duration.Value() compatibility with time.Duration")
}

func TestZeroDuration(t *testing.T) {
	i := new(Duration)
	assert.EqualValues(t, time.Duration(0), *i, "Duration.Scan() result")

	val, err := i.Value()
	assert.Nil(t, err, "Duration.Value() error")
	assert.EqualValues(t, "0 microseconds", val, "Duration.Value() result")

	assert.NoError(t, i.Scan("00:00:00"), "Duration.Scan() error")
	assert.EqualValues(t, time.Duration(0), *i, "Duration.Scan() result")

	val, err = i.Value()
	assert.Nil(t, err, "Duration.Value() error")
	assert.EqualValues(t, "0 microseconds", val, "Duration.Value() result")
}
