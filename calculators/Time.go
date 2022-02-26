package calculators

import "time"

func GetTimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func IsPartOfTimeRange(t time.Time, fromHour int, fromMin int, toHour int, toMin int) bool {
	fromTime := time.Date(t.Year(), t.Month(), t.Day(), fromHour, fromMin, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), toHour, toMin, 0, 0, t.Location())

	if toHour == 0 {
		if t.Hour() != 0 {
			endTime = endTime.AddDate(0, 0, 1)
		} else {
			// if it is 12 at night then the range is from the day before
			fromTime = time.Date(t.Year(), t.Month(), t.Day()-1, fromHour, fromMin, 0, 0, t.Location())
		}
	}

	return (t.After(fromTime) || t.Equal(fromTime)) && (t.Before(endTime) || t.Equal(endTime))
}

func GetTimeDifference(t time.Time, hour int, min int, secs int) time.Duration {
	toTime := time.Date(t.Year(), t.Month(), t.Day(), hour, min, secs, 0, t.Location())
	if hour == 0 {
		toTime = toTime.AddDate(0, 0, 1)
	}
	return toTime.Sub(t)
}
