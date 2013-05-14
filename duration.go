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

func (d Duration) RelativeMonths() time.Duration {
	return d.Duration % Year / Month
}

func (d Duration) RelativeWeeks() time.Duration {
	return d.Duration % Month / Week
}

func (d Duration) RelativeDays() time.Duration {
	return d.Duration % Week / Day
}

func (d Duration) RelativeHours() time.Duration {
	return d.Duration % Day / time.Hour
}

func (d Duration) RelativeMinutes() time.Duration {
	return d.Duration % time.Hour / time.Minute
}

func (d Duration) RelativeSeconds() time.Duration {
	return d.Duration % time.Minute / time.Second
}
