package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeUtilImpl_Timestamp2Time(t *testing.T) {
	assert.Equal(t, TimeUtilImpl{}.Timestamp2Time(1645173808).String(), "2022-02-18 16:43:28 +0800 CST")
}

func TestTimeUtilImpl_Timestamp2Str(t *testing.T) {
	assert.Equal(t, TimeUtilImpl{}.Timestamp2Str(1645173808), "2022-02-18")
}

func TestTimeUtilImpl_IsSameDay(t *testing.T) {
	testTable := []struct {
		t1     time.Time
		now    time.Time
		hour   int
		result bool
	}{
		{
			t1:     time.Date(2022, 2, 21, 4, 13, 2, 0, time.Local),
			now:    time.Date(2022, 2, 21, 14, 46, 21, 0, time.Local),
			hour:   5,
			result: false,
		},
		{
			t1:     time.Date(2022, 2, 21, 5, 13, 2, 0, time.Local),
			now:    time.Date(2022, 2, 21, 14, 46, 21, 0, time.Local),
			hour:   5,
			result: true,
		},
	}

	for _, e := range testTable {
		ret := TimeUtilImpl{}.IsSameDay(e.t1, e.now, e.hour)
		assert.Equal(t, ret, e.result)
	}
}

func TestTimeUtilImpl_IsSameWeek(t *testing.T) {
	testTable := []struct {
		t1      time.Time
		now     time.Time
		weekDay time.Weekday
		hour    int
		result  bool
	}{
		{
			t1:      time.Date(2022, 2, 27, 5, 13, 2, 0, time.Local),
			now:     time.Date(2022, 2, 21, 14, 14, 46, 21, time.Local),
			weekDay: time.Monday,
			hour:    5,
			result:  true,
		},
		{
			t1:      time.Date(2022, 2, 28, 4, 13, 2, 0, time.Local),
			now:     time.Date(2022, 2, 21, 14, 14, 46, 21, time.Local),
			weekDay: time.Monday,
			hour:    5,
			result:  true,
		},
		{
			t1:      time.Date(2022, 2, 28, 4, 13, 2, 0, time.Local),
			now:     time.Date(2022, 2, 21, 14, 14, 46, 21, time.Local),
			weekDay: time.Monday,
			hour:    5,
			result:  true,
		},
		{
			t1:      time.Date(2022, 2, 21, 4, 13, 2, 0, time.Local),
			now:     time.Date(2022, 2, 21, 14, 14, 46, 21, time.Local),
			weekDay: time.Monday,
			hour:    5,
			result:  false,
		},
		{
			t1:      time.Date(2022, 2, 27, 4, 13, 2, 0, time.Local),
			now:     time.Date(2022, 2, 27, 14, 14, 46, 21, time.Local),
			weekDay: time.Sunday,
			hour:    5,
			result:  false,
		},
	}

	for _, e := range testTable {
		ret := TimeUtilImpl{}.IsSameWeek(e.t1, e.now, e.weekDay, e.hour)
		assert.Equal(t, ret, e.result)
	}
}

func TestTimeUtilImpl_CalcLastHourAndNextHour(t *testing.T) {
	now := time.Unix(1645425981, 0)
	last, next := TimeUtilImpl{}.CalcLastHourAndNextHour(now, 5)
	assert.Equal(t, last, time.Unix(1645390800, 0))
	assert.Equal(t, next, time.Unix(1645477200, 0))

	testTable := []struct {
		now  time.Time
		last time.Time
		next time.Time
		hour int
	}{
		{
			now:  time.Date(2022, 2, 21, 14, 46, 21, 0, time.Local),
			last: time.Date(2022, 2, 21, 5, 0, 0, 0, time.Local),
			next: time.Date(2022, 2, 22, 5, 0, 0, 0, time.Local),
			hour: 5,
		},
	}

	for _, e := range testTable {
		retLast, retNext := TimeUtilImpl{}.CalcLastHourAndNextHour(e.now, e.hour)
		assert.Equal(t, e.last, retLast)
		assert.Equal(t, e.next, retNext)
	}
}

func TestTimeUtilImpl_CalcLastWeekAndNextWeek(t *testing.T) {
	testTable := []struct {
		now     time.Time
		last    time.Time
		next    time.Time
		weekDay time.Weekday
		hour    int
	}{
		{
			now:     time.Date(2022, 2, 21, 14, 46, 21, 0, time.Local),
			last:    time.Date(2022, 2, 21, 5, 0, 0, 0, time.Local),
			next:    time.Date(2022, 2, 28, 5, 0, 0, 0, time.Local),
			weekDay: time.Monday,
			hour:    5,
		},
		{
			now:     time.Date(2022, 2, 20, 14, 46, 21, 0, time.Local),
			last:    time.Date(2022, 2, 14, 5, 0, 0, 0, time.Local),
			next:    time.Date(2022, 2, 21, 5, 0, 0, 0, time.Local),
			weekDay: time.Monday,
			hour:    5,
		},
		{
			now:     time.Date(2022, 2, 21, 4, 46, 21, 0, time.Local),
			last:    time.Date(2022, 2, 14, 5, 0, 0, 0, time.Local),
			next:    time.Date(2022, 2, 21, 5, 0, 0, 0, time.Local),
			weekDay: time.Monday,
			hour:    5,
		},
		{
			now:     time.Date(2022, 2, 21, 4, 46, 21, 0, time.Local),
			last:    time.Date(2022, 2, 20, 5, 0, 0, 0, time.Local),
			next:    time.Date(2022, 2, 27, 5, 0, 0, 0, time.Local),
			weekDay: time.Sunday,
			hour:    5,
		},
	}

	for _, e := range testTable {
		retLast, retNext := TimeUtilImpl{}.CalcLastWeekAndNextWeek(e.now, e.weekDay, e.hour)
		assert.Equal(t, retLast, e.last)
		assert.Equal(t, retNext, e.next)
	}
}

func TestTimeUtilImpl_CalcLastMonthAndNextMonth(t *testing.T) {
	testTable := []struct {
		now  time.Time
		last time.Time
		next time.Time
		day  int
		hour int
	}{
		{
			now:  time.Date(2022, 2, 21, 14, 46, 21, 0, time.Local),
			last: time.Date(2022, 2, 1, 5, 0, 0, 0, time.Local),
			next: time.Date(2022, 3, 1, 5, 0, 0, 0, time.Local),
			day:  1,
			hour: 5,
		},
		{
			now:  time.Date(2022, 12, 01, 14, 46, 21, 0, time.Local),
			last: time.Date(2022, 12, 1, 5, 0, 0, 0, time.Local),
			next: time.Date(2023, 1, 1, 5, 0, 0, 0, time.Local),
			day:  1,
			hour: 5,
		},
	}

	for _, e := range testTable {
		retLast, retNext := TimeUtilImpl{}.CalcLastMonthAndNextMonth(e.now, e.day, e.hour)
		assert.Equal(t, retLast, e.last)
		assert.Equal(t, retNext, e.next)
	}
}

func TestTimeUtilImpl_CalcDailyManyHourTime(t *testing.T) {
	testTable := []struct {
		head   time.Time
		tail   time.Time
		hour   int
		points []time.Time
	}{
		{
			head: time.Date(2023, 1, 01, 5, 0, 0, 0, time.Local),
			tail: time.Date(2023, 1, 03, 23, 0, 0, 0, time.Local),
			hour: 9,
			points: []time.Time{
				time.Date(2023, 1, 1, 9, 0, 0, 0, time.Local),
				time.Date(2023, 1, 2, 9, 0, 0, 0, time.Local),
				time.Date(2023, 1, 3, 9, 0, 0, 0, time.Local),
			},
		},
	}

	for _, e := range testTable {
		assert.Equal(t, TimeUtilImpl{}.CalcDailyManyHourTime(e.head, e.tail, e.hour), e.points)
	}
}

func TestTimeUtilImpl_CalcWeeklyManyWeekDayHourTime(t *testing.T) {
	testTable := []struct {
		head    time.Time
		tail    time.Time
		hour    int
		weekDay []time.Weekday
		points  []time.Time
	}{
		{
			head: time.Date(2022, 2, 25, 4, 0, 0, 0, time.Local),
			tail: time.Date(2022, 3, 6, 4, 0, 0, 0, time.Local),
			hour: 5,
			weekDay: []time.Weekday{
				time.Tuesday,
				time.Sunday,
			},
			points: []time.Time{
				time.Date(2022, 2, 27, 5, 0, 0, 0, time.Local),
				time.Date(2022, 3, 1, 5, 0, 0, 0, time.Local),
			},
		},
	}

	for _, e := range testTable {
		assert.Equal(t, TimeUtilImpl{}.CalcWeeklyManyWeekDayHourTime(e.head, e.tail, e.hour, e.weekDay...), e.points)
	}
}
