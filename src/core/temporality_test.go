package core

import (
	"testing"
	"time"
)

func TestNewPeriod(t *testing.T) {
	start := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 8, 31, 23, 59, 59, 0, time.UTC)

	period, err := NewPeriod(start, end)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if period.Start != start || period.End != end {
		t.Errorf("expected start %v and end %v, got start %v and end %v", start, end, period.Start, period.End)
	}
}

func TestNewPeriod_Invalid(t *testing.T) {
	start := time.Date(2024, 8, 31, 23, 59, 59, 0, time.UTC)
	end := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)

	_, err := NewPeriod(start, end)
	if err == nil {
		t.Error("expected an error when end date is before start date")
	}
}

func TestPeriod_Contains(t *testing.T) {
	start := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 8, 31, 23, 59, 59, 0, time.UTC)
	period, _ := NewPeriod(start, end)

	testDate := time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)
	if !period.Contains(testDate) {
		t.Errorf("expected date %v to be contained in the period", testDate)
	}
}

func TestPeriod_Duration(t *testing.T) {
	start := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 8, 31, 23, 59, 59, 0, time.UTC)
	period, _ := NewPeriod(start, end)

	expectedDuration := end.Sub(start)
	if period.Duration() != expectedDuration {
		t.Errorf("expected duration %v, got %v", expectedDuration, period.Duration())
	}
}

func TestMonth(t *testing.T) {
	period, err := Month(2024, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedEnd := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	if !period.Start.Equal(expectedStart) {
		t.Errorf("Expected start %v, got %v", expectedStart, period.Start)
	}

	if !period.End.Equal(expectedEnd) {
		t.Errorf("Expected end %v, got %v", expectedEnd, period.End)
	}
}

func TestYear(t *testing.T) {
	period, err := Year(2024)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedEnd := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	if !period.Start.Equal(expectedStart) {
		t.Errorf("Expected start %v, got %v", expectedStart, period.Start)
	}

	if !period.End.Equal(expectedEnd) {
		t.Errorf("Expected end %v, got %v", expectedEnd, period.End)
	}
}

func TestPeriod_SplitJanuaryByDays(t *testing.T) {
	period, err := Month(2024, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	days := period.SplitByDays()
	count := 0
	current := *period

	for d := range days {
		count++
		current = d
	}

	if !current.Start.Equal(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("expected start %v, got %v", time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC), current.Start)
	}

	if count != 31 {
		t.Errorf("expected 31 days, got %v", count)
	}
}

func TestPeriod_SplitFebruaryByDays(t *testing.T) {
	period, err := Month(2024, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	days := period.SplitByDays()
	count := 0
	current := *period

	for d := range days {
		count++
		current = d
	}

	if !current.Start.Equal(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("expected start %v, got %v", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), current.Start)
	}

	if count != 29 {
		t.Errorf("expected 29 days, got %v", count)
	}
}

func TestPeriod_SplitYear2024ByDays(t *testing.T) {
	period, err := Year(2024)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	days := period.SplitByDays()
	count := 0
	current := *period

	for d := range days {
		count++
		current = d
	}

	if !current.Start.Equal(time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("expected start %v, got %v", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), current.Start)
	}

	if count != 366 {
		t.Errorf("expected 31 days, got %v", count)
	}
}

func TestPeriod_SplitYear2024ByMonths(t *testing.T) {
	period, err := Year(2024)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	days := period.SplitByMonths()
	count := 0
	var current Period

	for d := range days {
		count++
		current = d
	}

	if !current.Start.Equal(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("expected start %v, got %v", time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC), current.Start)
	}

	if !current.End.Equal(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("expected end %v, got %v", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), current.End)
	}

	if count != 12 {
		t.Errorf("expected 12 months, got %v", count)
	}
}
