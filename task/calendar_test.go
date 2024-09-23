package task

import (
	"testing"
	"time"
)

func TestOffsetHolidayMLK(t *testing.T) {
	testcases := map[int]time.Time{
		2021: time.Date(2021, time.January, 18, 0, 0, 0, 0, time.Local),
		2022: time.Date(2022, time.January, 17, 0, 0, 0, 0, time.Local),
		2023: time.Date(2023, time.January, 16, 0, 0, 0, 0, time.Local),
		2024: time.Date(2024, time.January, 15, 0, 0, 0, 0, time.Local),
		2025: time.Date(2025, time.January, 20, 0, 0, 0, 0, time.Local),
		2026: time.Date(2026, time.January, 19, 0, 0, 0, 0, time.Local),
		2027: time.Date(2027, time.January, 18, 0, 0, 0, 0, time.Local),
	}
	for year, expected := range testcases {
		date := MLK_DAY.DateIn(year)
		if date.Year() != expected.Year() || date.Month() != expected.Month() || date.Day() != expected.Day() {
			t.Errorf("expected %v, got %v", expected, date)
		}
	}
}

func TestOffsetHolidayPresidentsDay(t *testing.T) {
	holiday := PRESIDENTS_DAY
	testcases := map[int]time.Time{
		2021: time.Date(2021, holiday.Month(), 15, 0, 0, 0, 0, time.Local),
		2022: time.Date(2022, holiday.Month(), 21, 0, 0, 0, 0, time.Local),
		2023: time.Date(2023, holiday.Month(), 20, 0, 0, 0, 0, time.Local),
		2024: time.Date(2024, holiday.Month(), 19, 0, 0, 0, 0, time.Local),
		2025: time.Date(2025, holiday.Month(), 17, 0, 0, 0, 0, time.Local),
		2026: time.Date(2026, holiday.Month(), 16, 0, 0, 0, 0, time.Local),
		2027: time.Date(2027, holiday.Month(), 15, 0, 0, 0, 0, time.Local),
		2028: time.Date(2028, holiday.Month(), 21, 0, 0, 0, 0, time.Local),
		2029: time.Date(2029, holiday.Month(), 19, 0, 0, 0, 0, time.Local),
		2030: time.Date(2030, holiday.Month(), 18, 0, 0, 0, 0, time.Local),
	}
	for year, expected := range testcases {
		date := holiday.DateIn(year)
		if date.Year() != expected.Year() || date.Month() != expected.Month() || date.Day() != expected.Day() {
			t.Errorf("expected %v, got %v", expected, date)
		}
	}
}

func TestFixedDateHolidayNYE(t *testing.T) {
	holiday := NEW_YEARS_DAY
	for year := 2000; year <= 2030; year++ {
		expected := time.Date(year, holiday.Month(), 1, 0, 0, 0, 0, time.Local)
		date := holiday.DateIn(year)
		if date.Year() != expected.Year() || date.Month() != expected.Month() || date.Day() != expected.Day() {
			t.Errorf("expected %v, got %v", expected, date)
		}
	}
}

func TestCalendarDateIsHoliday(t *testing.T) {
	holiday := NewFixedDateHoliday(time.January, 1)
	for year := 2000; year <= 2030; year++ {
		expected := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
		date := holiday.DateIn(year)
		if date.Year() != expected.Year() || date.Month() != expected.Month() || date.Day() != expected.Day() {
			t.Errorf("expected %v, got %v", expected, date)
		}
	}
}

func TestCalendarIsHoliday(t *testing.T) {
	cal := &Calendar{}
	cal.Add(NEW_YEARS_DAY)
	cal.Add(MLK_DAY)
	if !cal.IsHoliday(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local)) {
		t.Fail()
	}
	if cal.IsHoliday(time.Date(2021, time.January, 2, 0, 0, 0, 0, time.Local)) {
		t.Fail()
	}
}
