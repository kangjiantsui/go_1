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
	now := time.Unix(1645425981, 0)

	t1 := time.Unix(1645387982, 0)
	assert.Equal(t, TimeUtilImpl{}.IsSameDay(&t1, &now, 5), false)

	t2 := time.Unix(1645391582, 0)
	assert.Equal(t, TimeUtilImpl{}.IsSameDay(&t2, &now, 5), true)
}

func TestTimeUtilImpl_IsSameWeek(t *testing.T) {
	now := time.Unix(1645425981, 0)

	t1 := time.Unix(1645909982, 0)
	assert.Equal(t, TimeUtilImpl{}.IsSameWeek(&t1, &now, time.Monday, 5), true)

	t2 := time.Unix(1645996382, 0)
	assert.Equal(t, TimeUtilImpl{}.IsSameWeek(&t2, &now, time.Monday, 5), false)

	t3 := time.Unix(1645992782, 0)
	assert.Equal(t, TimeUtilImpl{}.IsSameWeek(&t3, &now, time.Monday, 5), true)

	t4 := time.Unix(1645387982, 0)
	assert.Equal(t, TimeUtilImpl{}.IsSameWeek(&t4, &now, time.Monday, 5), false)
}

func TestTimeUtilImpl_CalcLastHourAndNextHour(t *testing.T) {
	now := time.Unix(1645425981, 0)
	last, next := TimeUtilImpl{}.CalcLastHourAndNextHour(&now, 5)
	assert.Equal(t, *last, time.Unix(1645390800, 0))
	assert.Equal(t, *next, time.Unix(1645477200, 0))
}

func TestTimeUtilImpl_CalcLastWeekAndNextWeek(t *testing.T) {
	now := time.Unix(1645425981, 0)
	last, next := TimeUtilImpl{}.CalcLastWeekAndNextWeek(&now, time.Monday, 5)
	assert.Equal(t, *last, time.Unix(1645390800, 0))
	assert.Equal(t, *next, time.Unix(1645995600, 0))

	now = time.Unix(1645304400, 0)
	last, next = TimeUtilImpl{}.CalcLastWeekAndNextWeek(&now, time.Monday, 5)
	assert.Equal(t, *last, time.Unix(1644786000, 0))
	assert.Equal(t, *next, time.Unix(1645390800, 0))

	now = time.Unix(1645387200, 0)
	last, next = TimeUtilImpl{}.CalcLastWeekAndNextWeek(&now, time.Monday, 5)
	assert.Equal(t, *last, time.Unix(1644786000, 0))
	assert.Equal(t, *next, time.Unix(1645390800, 0))
}

func TestTimeUtilImpl_CalcLastMonthAndNextMonth(t *testing.T) {
	now := time.Unix(1645425981, 0)
	last, next := TimeUtilImpl{}.CalcLastMonthAndNextMonth(&now, 1, 5)
	assert.Equal(t, *last, time.Unix(1643662800, 0))
	assert.Equal(t, *next, time.Unix(1646082000, 0))

	now = time.Unix(1671051600, 0)
	last, next = TimeUtilImpl{}.CalcLastMonthAndNextMonth(&now, 1, 5)
	assert.Equal(t, *last, time.Unix(1669842000, 0))
	assert.Equal(t, *next, time.Unix(1672520400, 0))
}

func TestTimeUtilImpl_CalcDailyManyHourTime(t *testing.T) {
	head := time.Unix(1672520400, 0)
	tail := time.Unix(1672758000, 0)
	p1 := time.Unix(1672534800, 0)
	p2 := time.Unix(1672621200, 0)
	p3 := time.Unix(1672707600, 0)
	assert.Equal(t, TimeUtilImpl{}.CalcDailyManyHourTime(&head, &tail, 9), []*time.Time{&p1, &p2, &p3})
}

func TestTimeUtilImpl_CalcWeeklyManyWeekDayHourTime(t *testing.T) {
	head := time.Unix(1645732800, 0)
	tail := time.Unix(1646510400, 0)
	p1 := time.Unix(1645909200, 0)
	p2 := time.Unix(1646082000, 0)
	t.Log(TimeUtilImpl{}.CalcWeeklyManyWeekDayHourTime(&head, &tail, 5, time.Tuesday, time.Sunday))
	assert.Equal(t, TimeUtilImpl{}.CalcWeeklyManyWeekDayHourTime(&head, &tail, 5, time.Tuesday, time.Sunday), []*time.Time{&p1, &p2})
}
