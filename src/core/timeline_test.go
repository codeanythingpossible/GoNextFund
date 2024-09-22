package core

import (
	"testing"
	"time"
)

func TestTimeline_Add_ShouldHave2Items(t *testing.T) {
	// Create a new empty Timeline for budget values
	budgetTimeline := NewTimeline[float64]()

	// Create periods for January and February
	january, _ := Month(2024, 1)
	february, _ := Month(2024, 2)

	// Add values to the timeline
	budgetTimeline.Add(*january, 1000.0)
	budgetTimeline.Add(*february, 1200.0)

	// Verify the number of items added to the timeline
	if len(budgetTimeline.GetAll()) != 2 {
		t.Errorf("Expected 2 items, got %d", len(budgetTimeline.GetAll()))
	}

	// Check the periods
	firstPeriod := budgetTimeline.Items[0].Period
	secondPeriod := budgetTimeline.Items[1].Period

	// Verify that the first period is January 2024
	expectedJanStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedJanEnd := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	if !firstPeriod.Start.Equal(expectedJanStart) {
		t.Errorf("Expected first period start to be %v, got %v", expectedJanStart, firstPeriod.Start)
	}
	if !firstPeriod.End.Equal(expectedJanEnd) {
		t.Errorf("Expected first period end to be %v, got %v", expectedJanEnd, firstPeriod.End)
	}

	// Verify that the second period is February 2024
	expectedFebStart := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	expectedFebEnd := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	if !secondPeriod.Start.Equal(expectedFebStart) {
		t.Errorf("Expected second period start to be %v, got %v", expectedFebStart, secondPeriod.Start)
	}
	if !secondPeriod.End.Equal(expectedFebEnd) {
		t.Errorf("Expected second period end to be %v, got %v", expectedFebEnd, secondPeriod.End)
	}

	// Verify the values added to the timeline
	if budgetTimeline.Items[0].Value != 1000.0 {
		t.Errorf("Expected first value to be 1000.0, got %v", budgetTimeline.Items[0].Value)
	}
	if budgetTimeline.Items[1].Value != 1200.0 {
		t.Errorf("Expected second value to be 1200.0, got %v", budgetTimeline.Items[1].Value)
	}
}

func TestTimeline_Add_ShouldHandleOverlappingPeriod(t *testing.T) {
	// Create a new empty Timeline for budget values
	budgetTimeline := NewTimeline[float64]()

	// Create periods for January and February
	january, _ := Month(2024, 1)
	february, _ := Month(2024, 2)

	// Add values to the timeline
	budgetTimeline.Add(*january, 1000.0)
	budgetTimeline.Add(*february, 1200.0)

	// Add a value for the period from 15 January to 5 February
	periodJan15ToFeb5 := Period{
		Start: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
	}
	budgetTimeline.Add(periodJan15ToFeb5, 900.0)

	// Verify the number of items added to the timeline
	if len(budgetTimeline.GetAll()) != 3 {
		t.Errorf("Expected 3 items, got %d", len(budgetTimeline.GetAll()))
	}

	// Check the periods
	p0 := budgetTimeline.Items[0].Period
	p1 := budgetTimeline.Items[1].Period
	p2 := budgetTimeline.Items[2].Period

	// Expected periods
	expectedP0Start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedP0End := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	expectedP1Start := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	expectedP1End := time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC)

	expectedP2Start := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	expectedP2End := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	// Verifying periods
	if !p0.Start.Equal(expectedP0Start) || !p0.End.Equal(expectedP0End) {
		t.Errorf("Expected period 0 to be %v - %v, got %v - %v", expectedP0Start, expectedP0End, p0.Start, p0.End)
	}

	if !p1.Start.Equal(expectedP1Start) || !p1.End.Equal(expectedP1End) {
		t.Errorf("Expected period 1 to be %v - %v, got %v - %v", expectedP1Start, expectedP1End, p1.Start, p1.End)
	}

	if !p2.Start.Equal(expectedP2Start) || !p2.End.Equal(expectedP2End) {
		t.Errorf("Expected period 2 to be %v - %v, got %v - %v", expectedP2Start, expectedP2End, p2.Start, p2.End)
	}

	// Verifying values
	if budgetTimeline.Items[0].Value != 1000.0 {
		t.Errorf("Expected first value to be 1000.0, got %v", budgetTimeline.Items[0].Value)
	}
	if budgetTimeline.Items[1].Value != 900.0 {
		t.Errorf("Expected second value to be 900.0, got %v", budgetTimeline.Items[1].Value)
	}
	if budgetTimeline.Items[2].Value != 1200.0 {
		t.Errorf("Expected third value to be 1200.0, got %v", budgetTimeline.Items[2].Value)
	}
}

func TestTimeline_FindIntersects_ShouldReturnOnePeriod(t *testing.T) {
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	apr2024, _ := Month(2024, 4)
	timeline := Timeline[int]{
		Items: []PeriodValue[int]{
			{
				Period: *jan2024,
				Value:  123,
			},
			{
				Period: *feb2024,
				Value:  456,
			},
			{
				Period: *mar2024,
				Value:  69,
			},
			{
				Period: *apr2024,
				Value:  987,
			},
		},
	}

	p, _ := Day(2024, 2, 5)
	result := timeline.FindIntersects(*p)
	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if !result[0].Period.Equal(*feb2024) {
		t.Errorf("Expected feb2024 to be %v, got %v", *feb2024, result[0].Period)
	}
}

func TestTimeline_FindIntersects_ShouldReturnTwoPeriods(t *testing.T) {
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	apr2024, _ := Month(2024, 4)
	timeline := Timeline[int]{
		Items: []PeriodValue[int]{
			{
				Period: *jan2024,
				Value:  123,
			},
			{
				Period: *feb2024,
				Value:  456,
			},
			{
				Period: *mar2024,
				Value:  69,
			},
			{
				Period: *apr2024,
				Value:  987,
			},
		},
	}

	p, _ := NewPeriod(time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC), time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC))
	result := timeline.FindIntersects(*p)
	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if !result[0].Period.Equal(*feb2024) {
		t.Errorf("Expected feb2024 to be %v, got %v", *feb2024, result[0].Period)
	}

	if !result[1].Period.Equal(*mar2024) {
		t.Errorf("Expected mar2024 to be %v, got %v", *mar2024, result[1].Period)
	}
}

func TestTimeline_FindIntersects_ShouldReturnZeroPeriod(t *testing.T) {
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	apr2024, _ := Month(2024, 4)
	timeline := Timeline[int]{
		Items: []PeriodValue[int]{
			{
				Period: *jan2024,
				Value:  123,
			},
			{
				Period: *feb2024,
				Value:  456,
			},
			{
				Period: *mar2024,
				Value:  69,
			},
			{
				Period: *apr2024,
				Value:  987,
			},
		},
	}

	p, _ := Day(2024, 7, 14)
	result := timeline.FindIntersects(*p)
	if len(result) != 0 {
		t.Errorf("Expected 0 items, got %d", len(result))
	}
}

func TestTimeline_Aggregate_ShouldReturn5Periods(t *testing.T) {
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	day15jan2024, _ := Day(2024, 1, 15)
	timeline := Timeline[int]{
		Items: []PeriodValue[int]{
			{
				Period: *jan2024,
				Value:  100,
			},
			{
				Period: *feb2024,
				Value:  200,
			},
			{
				Period: *mar2024,
				Value:  300,
			},
			{
				Period: *day15jan2024,
				Value:  80,
			},
		},
	}
	timeline.SortTimelineByPeriodStart()

	result, _ := timeline.Aggregate(func(p Period, a int, b int) int {
		return a + b
	})

	if len(result.Items) != 5 {
		t.Errorf("Expected 5 items, got %d", len(result.Items))
	}

	expectedPv1, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 1), DateOnly(2024, 1, 15), 100)
	expectedPv2, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 15), DateOnly(2024, 1, 16), 180)
	expectedPv3, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 16), DateOnly(2024, 2, 1), 100)
	expectedPv4, _ := NewPeriodValueFromTimes(DateOnly(2024, 2, 1), DateOnly(2024, 3, 1), 200)
	expectedPv5, _ := NewPeriodValueFromTimes(DateOnly(2024, 3, 1), DateOnly(2024, 4, 1), 300)

	if !result.Items[0].Period.Equal(expectedPv1.Period) || result.Items[0].Value != expectedPv1.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv1, result.Items[0])
	}

	if !result.Items[1].Period.Equal(expectedPv2.Period) || result.Items[1].Value != expectedPv2.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv2, result.Items[1])
	}

	if !result.Items[2].Period.Equal(expectedPv3.Period) || result.Items[2].Value != expectedPv3.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv3, result.Items[2])
	}

	if !result.Items[3].Period.Equal(expectedPv4.Period) || result.Items[3].Value != expectedPv4.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv4, result.Items[3])
	}

	if !result.Items[4].Period.Equal(expectedPv5.Period) || result.Items[4].Value != expectedPv5.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv5, result.Items[4])
	}

}

func TestTimeline_Aggregate_ShouldNotAggregateContiguousPeriodsAndReturnSamePeriods(t *testing.T) {
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	timeline := Timeline[int]{
		Items: []PeriodValue[int]{
			{
				Period: *jan2024,
				Value:  100,
			},
			{
				Period: *feb2024,
				Value:  200,
			},
			{
				Period: *mar2024,
				Value:  300,
			},
		},
	}
	timeline.SortTimelineByPeriodStart()

	result, _ := timeline.Aggregate(func(p Period, a int, b int) int {
		return a + b
	})

	if len(result.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(result.Items))
	}

	expectedPv1, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 1), DateOnly(2024, 2, 1), 100)
	expectedPv2, _ := NewPeriodValueFromTimes(DateOnly(2024, 2, 1), DateOnly(2024, 3, 1), 200)
	expectedPv3, _ := NewPeriodValueFromTimes(DateOnly(2024, 3, 1), DateOnly(2024, 4, 1), 300)

	if !result.Items[0].Period.Equal(expectedPv1.Period) || result.Items[0].Value != expectedPv1.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv1, result.Items[0])
	}

	if !result.Items[1].Period.Equal(expectedPv2.Period) || result.Items[1].Value != expectedPv2.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv2, result.Items[1])
	}

	if !result.Items[2].Period.Equal(expectedPv3.Period) || result.Items[2].Value != expectedPv3.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv3, result.Items[2])
	}

}

func TestTimeline_AggregateWithMultipleIntersects_ShouldReturn5Periods(t *testing.T) {
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	pv1, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 10), DateOnly(2024, 1, 17), 80)
	pv2, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 12), DateOnly(2024, 1, 15), 50)

	timeline := Timeline[int]{
		Items: []PeriodValue[int]{
			{
				Period: *jan2024,
				Value:  100,
			},
			{
				Period: *feb2024,
				Value:  200,
			},
			{
				Period: *mar2024,
				Value:  300,
			},
			*pv1,
			*pv2,
		},
	}
	timeline.SortTimelineByPeriodStart()

	result, _ := timeline.Aggregate(func(p Period, a int, b int) int {
		return a + b
	})

	if len(result.Items) != 5 {
		t.Errorf("Expected 5 items, got %d", len(result.Items))
	}

	expectedPv1, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 1), DateOnly(2024, 1, 15), 100)
	expectedPv2, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 15), DateOnly(2024, 1, 16), 180)
	expectedPv3, _ := NewPeriodValueFromTimes(DateOnly(2024, 1, 16), DateOnly(2024, 2, 1), 100)
	expectedPv4, _ := NewPeriodValueFromTimes(DateOnly(2024, 2, 1), DateOnly(2024, 3, 1), 200)
	expectedPv5, _ := NewPeriodValueFromTimes(DateOnly(2024, 3, 1), DateOnly(2024, 4, 1), 300)

	if !result.Items[0].Period.Equal(expectedPv1.Period) || result.Items[0].Value != expectedPv1.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv1, result.Items[0])
	}

	if !result.Items[1].Period.Equal(expectedPv2.Period) || result.Items[1].Value != expectedPv2.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv2, result.Items[1])
	}

	if !result.Items[2].Period.Equal(expectedPv3.Period) || result.Items[2].Value != expectedPv3.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv3, result.Items[2])
	}

	if !result.Items[3].Period.Equal(expectedPv4.Period) || result.Items[3].Value != expectedPv4.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv4, result.Items[3])
	}

	if !result.Items[4].Period.Equal(expectedPv5.Period) || result.Items[4].Value != expectedPv5.Value {
		t.Errorf("Expected period to be %v, got %v", *expectedPv5, result.Items[4])
	}

}

func TestResolvePeriods(t *testing.T) {
	// Créer des périodes de test.
	jan2024, _ := Month(2024, 1)
	feb2024, _ := Month(2024, 2)
	mar2024, _ := Month(2024, 3)
	pv1, _ := NewPeriodValueFromTimes[int](DateOnly(2024, 1, 10), DateOnly(2024, 1, 17), 80)
	pv2, _ := NewPeriodValueFromTimes[int](DateOnly(2024, 1, 12), DateOnly(2024, 1, 15), 50)

	// Créer la slice de PeriodValue.
	periodValues := []PeriodValue[int]{
		{
			Period: *jan2024,
			Value:  100,
		},
		{
			Period: *feb2024,
			Value:  200,
		},
		{
			Period: *mar2024,
			Value:  300,
		},
		*pv1,
		*pv2,
	}

	// Appeler resolvePeriods pour obtenir les périodes résolues.
	resolved := resolvePeriods(periodValues)

	// Définir les périodes attendues.
	expected := []Period{
		{
			Start: DateOnly(2024, 1, 1),
			End:   DateOnly(2024, 1, 10),
		},
		{
			Start: DateOnly(2024, 1, 10),
			End:   DateOnly(2024, 1, 12),
		},
		{
			Start: DateOnly(2024, 1, 12),
			End:   DateOnly(2024, 1, 15),
		},
		{
			Start: DateOnly(2024, 1, 15),
			End:   DateOnly(2024, 1, 17),
		},
		{
			Start: DateOnly(2024, 1, 17),
			End:   DateOnly(2024, 2, 1),
		},
		{
			Start: DateOnly(2024, 2, 1),
			End:   DateOnly(2024, 3, 1),
		},
		{
			Start: DateOnly(2024, 3, 1),
			End:   DateOnly(2024, 4, 1),
		},
	}

	// Vérifier la longueur des résultats.
	if len(resolved) != len(expected) {
		t.Fatalf("Expected %d periods, got %d", len(expected), len(resolved))
	}

	// Vérifier chaque période.
	for i, p := range expected {
		if !resolved[i].Start.Equal(p.Start) || !resolved[i].End.Equal(p.End) {
			t.Errorf("Period %d mismatch. Expected: %v - %v, Got: %v - %v",
				i, p.Start.Format("2006-01-02"), p.End.Format("2006-01-02"),
				resolved[i].Start.Format("2006-01-02"), resolved[i].End.Format("2006-01-02"))
		}
	}
}
