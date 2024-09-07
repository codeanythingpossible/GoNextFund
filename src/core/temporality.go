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
	if end.Before(start) {
		return nil, errors.New("la date de fin doit être après la date de début")
	}
	return &Period{Start: start, End: end}, nil
}

func (p *Period) Duration() time.Duration {
	return p.End.Sub(p.Start)
}

func (p *Period) Contains(t time.Time) bool {
	return (t.After(p.Start) || t.Equal(p.Start)) && (t.Before(p.End) || t.Equal(p.End))
}
