package task

import (
	"time"
)

type Calendar struct {
	Holidays    map[time.Month][]Holiday
	WeekendDays map[time.Weekday]struct{}
}

func (c *Calendar) Add(holiday Holiday) {
	if c.Holidays == nil {
		c.Holidays = make(map[time.Month][]Holiday)
	}
	if holidays_in_month, ok := c.Holidays[holiday.Month()]; ok {
		c.Holidays[holiday.Month()] = append(holidays_in_month, holiday)
	} else {
		c.Holidays[holiday.Month()] = []Holiday{holiday}

	}
}

func (c *Calendar) IsHoliday(date time.Time) bool {
	for _, holiday := range c.Holidays[date.Month()] {
		if holiday.IsOccurence(date) {
			return true
		}
	}
	return false
}

type Holiday interface {
	Month() time.Month
	DateIn(year int) time.Time
	IsOccurence(date time.Time) bool
	SetTimezone(timezone *time.Location)
}
type fixedDateHoliday struct {
	month    time.Month
	day      int
	timezone *time.Location
}

func NewOffsetDateHoliday(offset int, weekday time.Weekday, month time.Month) Holiday {
	if offset <= 0 || offset > 4 {
		panic("offset cannot be must be 1, 2, 3, or 4")
	}
	return &offsetDateHoliday{
		offset:   offset,
		weekday:  weekday,
		month:    month,
		timezone: time.Local,
	}

}

type offsetDateHoliday struct {
	offset   int // 1,2,3,4,-1
	weekday  time.Weekday
	month    time.Month
	timezone *time.Location
}

func (h *offsetDateHoliday) IsOccurence(date time.Time) bool {
	date_in_year := h.DateIn(date.Year())
	return date_in_year.Month() == date.Month() && date_in_year.Day() == date.Day()
}

func (h *offsetDateHoliday) Month() time.Month {
	return h.month
}
func (h *offsetDateHoliday) SetTimezone(timezone *time.Location) {
	h.timezone = timezone
}

func (h offsetDateHoliday) DateIn(year int) time.Time {
	if h.timezone == nil {
		h.timezone = time.Local
	}
	date := time.Date(year, h.month, 1, 0, 0, 0, 0, h.timezone)
	for date.Weekday() != h.weekday {
		date = date.AddDate(0, 0, 1)
	}
	date = date.AddDate(0, 0, 7*(h.offset-1))
	return date
}

// fixed date
// offset
func NewFixedDateHoliday(month time.Month, day int) Holiday {
	return &fixedDateHoliday{
		month: month,
		day:   day,
	}
}

func (h *fixedDateHoliday) IsOccurence(date time.Time) bool {
	date_in_year := h.DateIn(date.Year())
	return date_in_year.Month() == date.Month() && date_in_year.Day() == date.Day()
}

func (h *fixedDateHoliday) Month() time.Month {
	return h.month
}
func (h *fixedDateHoliday) SetTimezone(timezone *time.Location) {
	h.timezone = timezone
}
func (h fixedDateHoliday) DateIn(year int) time.Time {
	if h.timezone == nil {
		h.timezone = time.Local
	}
	return time.Date(year, h.month, h.day, 0, 0, 0, 0, h.timezone)
}

func (c *Calendar) BusinessDayOffset(days int) (next_business_day time.Time) {
	now := time.Now()
	delta := 1
	if days < 0 {
		delta = -1
	}
	for days != 0 {
		next_business_day = now.AddDate(0, 0, delta)
		if next_business_day.Weekday() == time.Saturday {
			next_business_day = next_business_day.AddDate(0, 0, 2)
		}
	}
	return next_business_day
}
