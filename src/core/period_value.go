package core

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
