package core

import (
	"errors"
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

// SortTimelineByPeriodStart sorts the Timeline items by the Start date of their Periods
func (t *Timeline[T]) SortTimelineByPeriodStart() {
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
	t.SortTimelineByPeriodStart()
}

// GetAll returns all PeriodValues in the Timeline
func (t *Timeline[T]) GetAll() []PeriodValue[T] {
	return t.Items
}

// Aggregate returns another Timeline having all values with same period aggregated, slicing them if necessary.
func (t *Timeline[T]) Aggregate(f func(p Period, a T, b T) T) (Timeline[T], error) {
	var items []PeriodValue[T]
	var buffer []PeriodValue[T]
	var currentPeriod Period

	for i, next := range t.Items {
		if i == 0 {
			currentPeriod = next.Period
			buffer = append(buffer, next)
			continue
		}

		// We assume that periods are chronologically sorted
		if next.Period.Before(currentPeriod) {
			return Timeline[T]{}, errors.New("timeline should have sorted periods")
		}

		if next.Period.After(currentPeriod) {

			periods := resolvePeriods(buffer)

			for _, period := range periods {
				var currentValue T

				for _, candidate := range buffer {
					if candidate.Period.Intersects(period) {
						currentValue = f(period, candidate.Value, currentValue)
					}
				}

				items = append(items, NewPeriodValue(period, currentValue))
			}

			currentPeriod = next.Period

			buffer = clampPeriods(buffer, currentPeriod)
			buffer = append(buffer, next)

			continue
		}

		period, err := NewPeriod(currentPeriod.Start, maxTime(next.Period.Start, currentPeriod.End))
		if err != nil {
			return Timeline[T]{}, err
		}
		currentPeriod = *period
		buffer = append(buffer, next)
	}

	for _, pv := range buffer {
		items = append(items, pv)
	}

	return Timeline[T]{Items: items}, nil
}

func clampPeriods[T any](periodValues []PeriodValue[T], limit Period) []PeriodValue[T] {
	var results []PeriodValue[T]

	for _, pv := range periodValues {
		clamp, err := pv.Clamp(limit)
		if err == nil && !clamp.Period.IsEmpty() {
			results = append(results, clamp)
		}
	}

	return results
}

func resolvePeriods[T any](periodValues []PeriodValue[T]) []Period {
	var result []Period
	timeMap := make(map[time.Time]struct{}, len(periodValues)*2)

	for _, pv := range periodValues {
		timeMap[pv.Period.Start] = struct{}{}
		timeMap[pv.Period.End] = struct{}{}
	}

	// Extract keys (times) and ensure order
	allTimes := make([]time.Time, 0, len(timeMap))
	for t := range timeMap {
		allTimes = append(allTimes, t)
	}
	sort.Slice(allTimes, func(i, j int) bool {
		return allTimes[i].Before(allTimes[j])
	})

	for i := 1; i < len(allTimes); i++ {
		result = append(result, Period{
			Start: allTimes[i-1],
			End:   allTimes[i],
		})
	}

	return result
}
