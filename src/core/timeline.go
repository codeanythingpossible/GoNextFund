package core

import (
	"sort"
	"time"
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

// Helper function to find the minimum of two times
func minTime(t1, t2 time.Time) time.Time {
	if t1.Before(t2) {
		return t1
	}
	return t2
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
