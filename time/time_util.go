package main

import (
	"time"
)

const (
	TimeFormat = "2006-01-02"
)

// 获取本周指定星期指定小时的时间
func getWeekTime(timePram time.Time, weekDay time.Weekday, hour int) time.Time {

	var ret time.Time
	if timePram.Hour() < hour {
		temp := time.Date(timePram.Year(), timePram.Month(), timePram.Day(), hour, 0, 0, 0, time.Local).Add(time.Hour * time.Duration(-24))
		timePram = temp
	} else {
		temp := time.Date(timePram.Year(), timePram.Month(), timePram.Day(), hour, 0, 0, 0, time.Local)
		timePram = temp
	}
	mondayTime := timePram.Add(time.Hour * time.Duration(24*(time.Monday-timePram.Weekday())))
	if timePram.Weekday() == time.Sunday {
		mondayTime = timePram.Add(time.Hour * 24 * -6)
	}
	if weekDay != time.Sunday {
		temp := mondayTime.Add(time.Hour * 24 * time.Duration(weekDay-1))
		ret = temp
	} else {
		temp := mondayTime.Add(time.Hour * 24 * 6)
		ret = temp
	}
	return ret
}

// 获得当天该小时的时间戳
func getHourTime(now time.Time, hour int) time.Time {
	timeHour := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local)
	// 当前时间在今日该小时之前，则返回昨日该小时
	if now.Before(timeHour) {
		timeHour = timeHour.Add(time.Hour * time.Duration(-24))
	}
	return timeHour
}

type TimeUtil interface {
	// IsSameDay 1. 按照 n 点刷新,(和当前时间)是否在同一天(若不是同一天则返回当天的这个小时)
	IsSameDay(t1 time.Time, t2 time.Time, hour int) (isSameDay bool)
	// IsSameWeek 2. 按照 n 点刷新,(和当前时间)是否在同一周(若不是同一周则返回本周的这个小时)
	IsSameWeek(t1 time.Time, t2 time.Time, weekDay time.Weekday, hour int) (isSameWeek bool)
	// CalcLastHourAndNextHour  3. 按照每天刷新, [%H] 点刷新,下一次时间点,上一次时间点
	CalcLastHourAndNextHour(timeParam time.Time, hour int) (last time.Time, next time.Time)
	// CalcLastWeekAndNextWeek  4. 按照每周刷新,周 [%D], [%H]点刷新,下一次时间点,上一次时间点
	CalcLastWeekAndNextWeek(timeParam time.Time, weekDay time.Weekday, hour int) (last time.Time, next time.Time)
	// CalcLastMonthAndNextMonth  5. 按照每月刷新,[%D-%H-%M, %D-%H-%M] 号刷新,[%D-%H-%M, %D-%H-%M] 点刷新,下一次时间点,上一次时间点
	CalcLastMonthAndNextMonth(timeParam time.Time, day int, hour int) (last time.Time, next time.Time)
	// CalcDailyManyHourTime 6. 按照每天,[Y1, Y2] 点刷新,两个时间点之间,有多少次,分别是什么时间点
	CalcDailyManyHourTime(head time.Time, tail time.Time, hour ...int) (timePoints []time.Time)
	// CalcWeeklyManyWeekDayHourTime  7. 按照每周,周[X1, X2],点刷新,两个时间点之间有多少次,分别是什么时间点
	CalcWeeklyManyWeekDayHourTime(head time.Time, tail time.Time, hour int, weekDay ...time.Weekday) (timePoints []time.Time)
	// CalcEachMonthManyMonthDayManyHourTime  8. 按照每月 [X1, X2] 号, [Y1, Y2] 点刷新,两个时间点之间有多少次,分别是什么时间点
	CalcEachMonthManyMonthDayManyHourTime(head time.Time, tail time.Time, hour int, monthDay ...int) (timePoints []time.Time)
	// Timestamp2Time 时间戳转时间
	Timestamp2Time(timestamp int64) time.Time
	// ParseStr 日期字符串转时间
	ParseStr(str string) (*time.Time, error)
	// Timestamp2Str 时间戳转字符串
	Timestamp2Str(timestamp int64) string
}

type TimeUtilImpl struct {
}

func (t TimeUtilImpl) IsSameDay(t1 time.Time, t2 time.Time, hour int) (isSameDay bool) {
	if getHourTime(t1, hour) != getHourTime(t2, hour) {
		return false
	}
	return true
}

func (t TimeUtilImpl) IsSameWeek(t1 time.Time, t2 time.Time, weekDay time.Weekday, hour int) (isSameWeek bool) {
	if getWeekTime(t1, weekDay, hour) != getWeekTime(t2, weekDay, hour) {
		return false
	}
	return true
}

func (t TimeUtilImpl) CalcLastHourAndNextHour(timeParam time.Time, hour int) (last time.Time, next time.Time) {
	last = getHourTime(timeParam, hour)
	temp := last.Add(time.Hour * 24)
	next = temp
	return last, next
}

func (t TimeUtilImpl) CalcLastWeekAndNextWeek(timeParam time.Time, weekDay time.Weekday, hour int) (last time.Time, next time.Time) {
	timeParam = getHourTime(timeParam, hour)
	if timeParam.Weekday() == time.Sunday || timeParam.Weekday() >= weekDay {
		last = getWeekTime(timeParam, weekDay, hour)
		temp := last.In(time.Local).Add(time.Hour * 24 * 7).In(time.Local)
		next = temp
	} else {
		next = getWeekTime(timeParam, weekDay, hour)
		temp := next.Add(time.Hour*24 - 7)
		last = temp
	}
	return last, next
}

func (t TimeUtilImpl) CalcLastMonthAndNextMonth(timeParam time.Time, day int, hour int) (last time.Time, next time.Time) {
	timeParam = getHourTime(timeParam, hour)
	if timeParam.Day() >= day {
		temp := time.Date(timeParam.Year(), timeParam.Month(), day, hour, 0, 0, 0, time.Local)
		last = temp
		temp2 := time.Date(last.Year(), last.Month()+1, day, hour, 0, 0, 0, time.Local)
		next = temp2
		if last.Month() == 12 {
			temp1 := time.Date(last.Year()+1, 1, day, hour, 0, 0, 0, time.Local)
			next = temp1
		}
	} else {
		temp := time.Date(timeParam.Year(), timeParam.Month(), day, hour, 0, 0, 0, time.Local)
		next = temp
		temp2 := time.Date(last.Year(), last.Month()+1, day, hour, 0, 0, 0, time.Local)
		next = temp2
		if next.Month() == 1 {
			temp1 := time.Date(next.Year()-1, 12, day, hour, 0, 0, 0, time.Local)
			last = temp1
		}
	}
	return last, next
}

func (t TimeUtilImpl) CalcDailyManyHourTime(head time.Time, tail time.Time, hour ...int) (timePoints []time.Time) {
	if hour == nil {
		return nil
	}
	hours := map[int]interface{}{}
	for _, e := range hour {
		hours[e] = nil
	}
	temp := time.Date(head.Year(), head.Month(), head.Day(), head.Hour(), head.Minute(), head.Second(), 0, time.Local)
	begin := &temp
	for {
		if begin.After(tail) {
			break
		}
		if _, ok := hours[begin.Hour()]; ok {
			timePoint := time.Date(begin.Year(), begin.Month(), begin.Day(), begin.Hour(), begin.Minute(), begin.Second(), 0, time.Local)
			timePoints = append(timePoints, timePoint)
		}
		temp1 := begin.Add(time.Hour)
		begin = &temp1
	}
	return timePoints
}

func (t TimeUtilImpl) CalcWeeklyManyWeekDayHourTime(head time.Time, tail time.Time, hour int, weekDay ...time.Weekday) (timePoints []time.Time) {
	if weekDay == nil {
		return nil
	}
	weekDays := map[time.Weekday]interface{}{}
	for _, e := range weekDay {
		weekDays[e] = nil
	}
	temp := time.Date(head.Year(), head.Month(), head.Day(), head.Hour(), head.Minute(), head.Second(), 0, time.Local)
	begin := &temp
	for {
		if begin.After(tail) {
			break
		}
		if _, ok := weekDays[begin.Weekday()]; ok && begin.Hour() == hour {
			timePoint := time.Date(begin.Year(), begin.Month(), begin.Day(), begin.Hour(), begin.Minute(), begin.Second(), 0, time.Local)
			timePoints = append(timePoints, timePoint)
		}
		temp1 := begin.Add(time.Hour)
		begin = &temp1
	}
	return timePoints
}

func (t TimeUtilImpl) CalcEachMonthManyMonthDayManyHourTime(head time.Time, tail time.Time, hour int, monthDay ...int) (timePoints []time.Time) {
	if monthDay == nil {
		return nil
	}
	monthDays := map[int]interface{}{}
	for _, e := range monthDay {
		monthDays[e] = nil
	}
	temp := time.Date(head.Year(), head.Month(), head.Day(), head.Hour(), head.Minute(), head.Second(), 0, time.Local)
	begin := &temp
	for {
		if begin.After(tail) {
			break
		}
		if _, ok := monthDays[begin.Day()]; ok && begin.Hour() == hour {
			timePoint := time.Date(begin.Year(), begin.Month(), begin.Day(), begin.Hour(), begin.Minute(), begin.Second(), 0, time.Local)
			timePoints = append(timePoints, timePoint)
		}
		temp1 := begin.Add(time.Hour)
		begin = &temp1
	}
	return timePoints
}

func (t TimeUtilImpl) Timestamp2Time(timestamp int64) time.Time {
	ret := time.Unix(timestamp, 0)
	return ret
}

func (t TimeUtilImpl) ParseStr(str string) (*time.Time, error) {
	ret, err := time.ParseInLocation(TimeFormat, str, time.Local)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (t TimeUtilImpl) Timestamp2Str(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeFormat)
}
