package task

import "time"

var NEW_YEARS_DAY = NewFixedDateHoliday(time.January, 1)
var MLK_DAY = NewOffsetDateHoliday(3, time.Monday, time.January)
var PRESIDENTS_DAY = NewOffsetDateHoliday(3, time.Monday, time.February)
var MEMORIAL_DAY = NewOffsetDateHoliday(4, time.Monday, time.May)
var INDEPENDENCE_DAY = NewFixedDateHoliday(time.July, 4)
var LABOR_DAY = NewOffsetDateHoliday(1, time.Monday, time.September)
var INDIGENOUS_PEOPLES_DAY = NewOffsetDateHoliday(2, time.Monday, time.October)
var VETERANS_DAY = NewFixedDateHoliday(time.November, 11)
var THANKSGIVING_DAY = NewOffsetDateHoliday(4, time.Thursday, time.November)
var CHRISTMAS_DAY = NewFixedDateHoliday(time.December, 25)

var US_FEDERAL_HOLIDAYS = []Holiday{
	NEW_YEARS_DAY,
	MLK_DAY,
	PRESIDENTS_DAY,
	MEMORIAL_DAY,
	INDEPENDENCE_DAY,
	LABOR_DAY,
	INDIGENOUS_PEOPLES_DAY,
	VETERANS_DAY,
	THANKSGIVING_DAY,
	CHRISTMAS_DAY,
}
