package julian

import (
	"math"
	"time"
)

func GregorianToJulianUTC(timeDate time.Time) float64 {
	timeDate = timeDate.In(time.UTC)

	year := float64(timeDate.Year())
	month := float64(timeDate.Month())
	date := float64(timeDate.Day())

	if month < 3 {
		year = year - 1
		month = month + 12
	}

	a := math.Floor(year / 100)
	b := 0.0

	if year == 1582 {
		if month == 10 {
			if date > 4 {
				b = 2.0 - a + math.Floor(a/4.0)
			}
		} else {
			b = 2.0 - a + math.Floor(a/4.0)
		}
	} else {
		b = 2.0 - a + math.Floor(a/4.0)
	}

	return 1720994.5 + math.Floor(365.25*year) + math.Floor(30.6001*(month+1)) + b + date + float64(timeDate.Hour())/24. + float64(timeDate.Minute())/24.*60.
}
