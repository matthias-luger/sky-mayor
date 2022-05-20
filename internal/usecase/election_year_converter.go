package usecase

import (
	"time"
)

func GetTimeSpanForSkyblockYear(year int) (time.Time, time.Time) {
	var year88 = time.Unix(1599664500000/1000, 0)
	var yearSpanHours = 124
	var start time.Time
	var end time.Time
	var numberOfYears int

	if year > 88 {
		numberOfYears = year - 88
		start = year88.Add(time.Hour * time.Duration(numberOfYears*yearSpanHours))
		end = start.Add(time.Hour * time.Duration(yearSpanHours))
	} else {
		numberOfYears = 88 - year
		start = year88.Add(time.Hour * time.Duration(numberOfYears*yearSpanHours) * -1)
		end = start.Add(time.Hour * time.Duration(yearSpanHours))
	}

	return start, end
}
