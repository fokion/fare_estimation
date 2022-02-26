package calculators

import "time"

func GetTimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func IsPartOfTimeRange(t time.Time, fromHour int, fromMin int, toHour int, toMin int) bool {
	fromTime := time.Date(t.Year(), t.Month(), t.Day(), fromHour, fromMin, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), toHour, toMin, 0, 0, t.Location())
	if toHour == 0 {
		endTime = endTime.AddDate(0, 0, 1)
	}

	return (t.After(fromTime) || t.Equal(fromTime)) && (t.Before(endTime) || t.Equal(endTime))
}
