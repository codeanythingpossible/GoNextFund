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

func (p *Period) Duration() time.Duration {
	return p.End.Sub(p.Start)
}

func (p *Period) Contains(t time.Time) bool {
	return (t.After(p.Start) || t.Equal(p.Start)) && (t.Before(p.End) || t.Equal(p.End))
}

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

func (p *Period) SplitByDays() <-chan Period {
	return p.Split(func(current time.Time) time.Time { return current.AddDate(0, 0, 1) })
}

func (p *Period) SplitByMonths() <-chan Period {
	return p.Split(func(current time.Time) time.Time { return current.AddDate(0, 1, 0) })
}
