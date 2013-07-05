/*
time Duration utilities. Extends duration constants up to a year, and
provides methods for modulo durations

	d := MakeDuration(time.Minute * 72)
	hours := d.Hours()             // 1
	minutes := d.RemainingMinutes()  // 12
*/
package rye

import (
	"time"
)

type Duration struct {
	time.Duration
}

func MakeDuration(d time.Duration) Duration {
	return Duration{d}
}

const (
	Day   time.Duration = 24 * time.Hour
	Week                = 7 * Day
	Month               = 4 * Week
	Year                = 12 * Month
)

func (d Duration) Days() time.Duration {
	return d.Duration / Day
}

func (d Duration) Weeks() time.Duration {
	return d.Duration / Week
}

func (d Duration) Months() time.Duration {
	return d.Duration / Month
}

func (d Duration) Years() time.Duration {
	return d.Duration / Year
}

func (d Duration) RemainingSeconds() time.Duration {
	return d.Duration % time.Minute / time.Second
}

func (d Duration) RemainingMinutes() time.Duration {
	return d.Duration % time.Hour / time.Minute
}

func (d Duration) RemainingHours() time.Duration {
	return d.Duration % Day / time.Hour
}

func (d Duration) RemainingDays() time.Duration {
	return d.Duration % Week / Day
}

func (d Duration) RemainingWeeks() time.Duration {
	return d.Duration % Month / Week
}

func (d Duration) RemainingMonths() time.Duration {
	return d.Duration % Year / Month
}

