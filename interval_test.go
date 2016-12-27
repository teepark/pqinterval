package pqinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCreatesCorrectInterval(t *testing.T) {
	i := New(1, 2, 3, 4, 5, 6)

	assert.EqualValues(t, 1, i.Years(), "interval years")
	assert.EqualValues(t, 2*24+3, i.Hours(), "interval hours")
	assert.EqualValues(
		t,
		4*60*1000000+5*1000000+6,
		i.Microseconds(),
		"interval microseconds",
	)
}

func TestScanInterval(t *testing.T) {
	i := new(Interval)
	i.Scan("2 days")

	assert.EqualValues(t, 0, i.Years(), "scanned interval years")
	assert.EqualValues(t, 48, i.Hours(), "scanned interval hours")
	assert.EqualValues(t, 0, i.Microseconds(), "scanned interval microseconds")
}
