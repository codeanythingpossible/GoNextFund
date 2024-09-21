package core

import (
	"sort"
)

// Timeline represents a list of PeriodValue objects
type Timeline[T any] struct {
	Items []PeriodValue[T]
}

// NewTimeline creates and returns an empty Timeline.
func NewTimeline[T any]() Timeline[T] {
	return Timeline[T]{
		Items: []PeriodValue[T]{}, // Initialize an empty slice
	}
}

// SortTimelineByStart sorts the Timeline items by the Start date of their Periods
func (t *Timeline[T]) SortTimelineByStart() {
	sort.Slice(t.Items, func(i, j int) bool {
		return t.Items[i].Period.Start.Before(t.Items[j].Period.Start)
	})
}

func (t *Timeline[T]) FindIntersects(period Period) []PeriodValue[T] {
	var items []PeriodValue[T]

	for _, current := range t.Items {

		// items are ordered, so if we are after period then we finished scan
		if current.Period.Start.After(period.End) {
			break
		}

		if current.Period.Intersects(period) {
			items = append(items, current)
		}
	}

	return items
}

// Add allows adding a new PeriodValue to the Timeline
func (t *Timeline[T]) Add(newPeriod Period, newValue T) {
	// Update the Timeline items with the new list
	t.Items = append(t.Items, PeriodValue[T]{
		Period: newPeriod,
		Value:  newValue,
	})

	// Ensure the items are sorted by the Start date after adding
	t.SortTimelineByStart()
}

// GetAll returns all PeriodValues in the Timeline
func (t *Timeline[T]) GetAll() []PeriodValue[T] {
	return t.Items
}

// Aggregate returns another Timeline having all values with same period aggregated, slicing them if necessary.
func (t *Timeline[T]) Aggregate(f func(p Period, a T, b T) T) Timeline[T] {
	var items []PeriodValue[T]
	var buffer PeriodValue[T]

	for i, current := range t.Items {
		if i == 0 {
			buffer = current
			continue
		}

		// Note that periods are ordered
		if current.Period.Intersects(buffer.Period) {
			parts := buffer.Period.SplitFromPeriod(current.Period)
			for p := range parts {

				var value T
				if p.Intersects(current.Period) {
					value = f(p, current.Value, buffer.Value)
				} else {
					value = buffer.Value
				}
				pv := NewPeriodValue(p, value)
				items = append(items, pv)
			}
		}

		// if current is after buffer, then we can finalize buffer and compute next period
		if current.Period.Start.Compare(buffer.Period.End) >= 0 {
			items = append(items, current)
			buffer = current
			continue
		}
	}

	return Timeline[T]{Items: items}
}
