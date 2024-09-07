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
