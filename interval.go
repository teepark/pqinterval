package pqinterval

import (
	"errors"
	"time"
)

// Interval can represent the full range of PostgreSQL's interval type.
type Interval struct {
	// the top bit is the sign for the microseconds,
	// bottom 29 are the signed year.
	yrs uint32

	hrs int32

	// it takes 33 bits (ouch) to fit microseconds-per-hour with sign,
	// but we have extra space in 'yrs' so the top bit there is the
	// sign for these microseconds
	us uint32
}

// New creates an Interval.
func New(years, days, hours, minutes, seconds, microseconds int) Interval {
	if years > maxYear || years < minYear || hours > maxHour || hours < minHour {
		panic("interval outside range")
	}
	microseconds += seconds*usPerSec + minutes*usPerMin

	hours += days * 24
	hours += microseconds / usPerHr
	microseconds %= usPerHr

	if years < 0 {
		years = (-years) | yrSignBit
	}
	if microseconds < 0 {
		years |= usSignBit
		microseconds *= -1
	}

	return Interval{
		yrs: uint32(years),
		hrs: int32(hours),
		us:  uint32(microseconds),
	}
}

// Years returns the number of years in the interval.
func (ival Interval) Years() int32 {
	years := int32(ival.yrs & (yrSignBit - 1))
	if ival.yrs&yrSignBit != 0 {
		years *= -1
	}
	return years
}

// Hours returns the number of hours in the interval.
func (ival Interval) Hours() int32 {
	return ival.hrs
}

// Microseconds returns the number of microseconds in the interval,
// up to the number of microseconds in an hour.
func (ival Interval) Microseconds() int64 {
	us := int64(ival.us)
	if ival.yrs&usSignBit != 0 {
		us *= -1
	}
	return us
}

// Scan implements sql.Scanner.
func (ival *Interval) Scan(src interface{}) error {
	var s string
	switch x := src.(type) {
	case string:
		s = x
	case []byte:
		s = string(x)
	default:
		return errors.New(
			"pqinterval: converting driver.Value type %T (%q) to string: invalid syntax",
		)
	}

	result, err := parse(s)
	if err != nil {
		return err
	}

	*ival = result
	return nil
}

const (
	// the year range allowed in PostgreSQL intervals.
	maxYear = 0xaaaaaaa
	minYear = -0xaaaaaaa

	maxHour = 1 << 31
	minHour = -1 << 31

	yrSignBit = 0x10000000
	usSignBit = 0x80000000

	usPerSec = 1000000
	usPerMin = usPerSec * 60
	usPerHr  = usPerMin * 60

	// assumptions embedded in PostgreSQL's EXTRACT(EPOCH FROM <interval>)
	daysPerYr  = 365.25
	daysPerMon = 30

	hrsPerYr = daysPerYr * 24
	nsPerYr  = int64(hrsPerYr * time.Hour)
)
