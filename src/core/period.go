package core

import (
	"errors"
	"time"
)

type Period struct {
	Start time.Time
	End   time.Time
}

func NewPeriod(start, end time.Time) (*Period, error) {
	if !end.After(start) {
		return nil, errors.New("end date must be after start date")
	}
	return &Period{Start: start, End: end}, nil
}

// Day returns a Period for the given year, month and day.
// The start is given day, and the end is the next day (exclusive).
func Day(year int, month int, day int) (*Period, error) {
	start := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	nextDay := start.AddDate(0, 0, 1)

	return NewPeriod(start, nextDay)
}

// Month returns a Period for the given year and month.
// The start is the first day of the month, and the end is the first day of the next month (exclusive).
func Month(year int, month int) (*Period, error) {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	nextMonth := start.AddDate(0, 1, 0)

	return NewPeriod(start, nextMonth)
}

// Year returns a Period for the given year.
// The start is the first day of the year, and the end is the first day of the next year (exclusive).
func Year(year int) (*Period, error) {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	nextYear := start.AddDate(1, 0, 0)

	return NewPeriod(start, nextYear)
}

// Equal compares two periods.
func (p Period) Equal(other Period) bool {
	return p.Start.Equal(other.Start) && p.End.Equal(other.End)
}

// Duration returns duration of given period.
func (p *Period) Duration() time.Duration {
	return p.End.Sub(p.Start)
}

// Contains check if a periods contains another one.
func (p *Period) Contains(t time.Time) bool {
	return (t.After(p.Start) || t.Equal(p.Start)) && (t.Before(p.End) || t.Equal(p.End))
}

// ContainsPeriod checks if the current period fully contains another period.
func (p *Period) ContainsPeriod(other Period) bool {
	return (p.Start.Before(other.Start) || p.Start.Equal(other.Start)) &&
		(p.End.After(other.End) || p.End.Equal(other.End))
}

// Intersects checks if two periods overlap.
func (p *Period) Intersects(other Period) bool {
	return p.Start.Before(other.End) && p.End.After(other.Start)
}

// Split a period using given function.
func (p *Period) Split(f func(current time.Time) time.Time) <-chan Period {
	ch := make(chan Period)

	go func() {
		defer close(ch)

		current := p.Start
		for current.Before(p.End) {
			next := f(current)
			ch <- Period{Start: current, End: next}
			current = next
		}
	}()

	return ch
}

// SplitByDays returns periods for each days in given period.
func (p *Period) SplitByDays() <-chan Period {
	return p.Split(func(current time.Time) time.Time { return current.AddDate(0, 0, 1) })
}

// SplitByMonths returns periods for each months in given period.
func (p *Period) SplitByMonths() <-chan Period {
	return p.Split(func(current time.Time) time.Time { return current.AddDate(0, 1, 0) })
}
