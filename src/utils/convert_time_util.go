package utils

import "time"

func ConvertFloat64ToInt64(data float64) int64 {
	return int64(data)
}

func ConvertMilliSecondToTime(data int64) time.Time {
	return time.UnixMilli(data)
}
