package core

import "time"

// PeriodValue associate a value to given period
type PeriodValue[T any] struct {
	Period Period
	Value  T
}

// NewPeriodValue create new PeriodValue
func NewPeriodValue[T any](period Period, value T) PeriodValue[T] {
	return PeriodValue[T]{
		Period: period,
		Value:  value,
	}
}

// NewPeriodValueFromTimes create new PeriodValue from times
func NewPeriodValueFromTimes[T any](start time.Time, end time.Time, value T) (*PeriodValue[T], error) {
	period, err := NewPeriod(start, end)
	if err != nil {
		return nil, err
	}

	pv := PeriodValue[T]{
		Period: *period,
		Value:  value,
	}
	return &pv, nil
}

func (p *PeriodValue[T]) IsEmpty() bool {
	return p.Period.IsEmpty()
}
