package pqinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseYears(t *testing.T) {
	i, err := parse("3 years")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, i.yrs, 3, "parsed interval yrs")
	assert.EqualValues(t, i.hrs, 0, "parsed interval hrs")
	assert.EqualValues(t, i.us, 0, "parsed interval us")
}

func TestParseNegativeYears(t *testing.T) {
	i, err := parse("-3 years")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, yrSignBit|3, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 0, i.us, "parsed interval us")
}

func TestParseMonths(t *testing.T) {
	i, err := parse("6 mons")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, i.yrs, 0, "parsed interval yrs")
	assert.EqualValues(t, i.hrs, 24*30*6, "parsed interval hrs")
	assert.EqualValues(t, i.us, 0, "parsed interval us")
}

func TestParseNegativeMonths(t *testing.T) {
	i, err := parse("-8 mons")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, -8*30*24, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 0, i.us, "parsed interval us")
}

func TestParseDays(t *testing.T) {
	i, err := parse("11 days")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 11*24, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 0, i.us, "parsed interval us")
}

func TestParseNegativeDays(t *testing.T) {
	i, err := parse("-43 days")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, -43*24, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 0, i.us, "parsed interval us")
}

func TestParseHours(t *testing.T) {
	i, err := parse("12:00:00")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 12, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 0, i.us, "parsed interval us")
}

func TestParseNegativeHours(t *testing.T) {
	i, err := parse("-04:00:00")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs&^usSignBit, "parsed interval yrs")
	assert.EqualValues(t, -4, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 0, i.us, "parsed interval us")
}

func TestParseMinutes(t *testing.T) {
	i, err := parse("00:43:00")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, int64(43)*60*1000000, i.us, "parsed interval us")
}

func TestParseNegativeMinutes(t *testing.T) {
	i, err := parse("-00:07:00")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, int64(usSignBit), i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 7*60*1000000, i.us, "parsed interval us")
}

func TestParseSeconds(t *testing.T) {
	i, err := parse("00:00:33")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 33*1000000, i.us, "parsed interval us")
}

func TestParseNegativeSeconds(t *testing.T) {
	i, err := parse("-00:00:41")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, int64(usSignBit), i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 41*1000000, i.us, "parsed interval us")
}

func TestParseMicroseconds(t *testing.T) {
	i, err := parse("00:00:00.003456")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 3456, i.us, "parsed interval us")
}

func TestParseMicrosecondsMissingPlaces(t *testing.T) {
	i, err := parse("00:00:00.3456")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 0, i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 345600, i.us, "parsed interval us")
}

func TestParseNegativeMicroseconds(t *testing.T) {
	i, err := parse("-00:00:00.0011")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, int64(usSignBit), i.yrs, "parsed interval yrs")
	assert.EqualValues(t, 0, i.hrs, "parsed interval hrs")
	assert.EqualValues(t, 1100, i.us, "parsed interval us")
}

func TestParseCombined(t *testing.T) {
	i, err := parse("2 years 7 mons 9 days 07:44:18.472719")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, 2, i.yrs, "parsed interval yrs")
	assert.EqualValues(
		t,
		7*30*24+9*24+7,
		i.hrs,
		"parsed interval hrs",
	)
	assert.EqualValues(
		t,
		int64(44)*60*1000000+18*1000000+472719,
		i.us,
		"parsed interval us",
	)
}

func TestParseCombinedNegative(t *testing.T) {
	i, err := parse("-14 years -2 mons -8 days -11:22:33.456789")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, int64(yrSignBit)|usSignBit|14, i.yrs, "parsed interval yrs")
	assert.EqualValues(
		t,
		-2*30*24-8*24-11,
		i.hrs,
		"parsed interval hrs",
	)
	assert.EqualValues(
		t,
		22*60*1000000+33*1000000+456789,
		i.us,
		"parsed interval us",
	)
}

func TestParseMixedSigns(t *testing.T) {
	i, err := parse("-7 years 4 mons -2 days 11:22:33.456789")
	assert.Nil(t, err, "parse error")

	assert.EqualValues(t, yrSignBit|7, i.yrs, "parsed interval yrs")
	assert.EqualValues(
		t,
		4*30*24-2*24+11,
		i.hrs,
		"parsed interval hrs",
	)
	assert.EqualValues(
		t,
		22*60*1000000+33*1000000+456789,
		i.us,
		"parsed interval us",
	)
}
