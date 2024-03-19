package timeutil

import "time"

func GetCurrentTimeISO() string {
    return ConvertTimeToIso(time.Now())
}

func ConvertTimeToIso(inputTime time.Time) string {
    return inputTime.Format(time.RFC3339)
}

