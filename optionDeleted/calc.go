package optionDeleted

import (
	"math"
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/moslem-salat-schedule/utils/sunPosition"
)

func (o *Option) fillSunPosition(julianDay, timezone float64, longitude angle.Angle) {
	sunPos := sunPosition.CalSunPosition(julianDay, timezone, longitude)

	o.julianDate = sunPos.JulianDate

	o.meanAnomaly = sunPos.MeanAnomaly
	o.meanLongSun = sunPos.MeanLongSun
	o.eclipticLong = sunPos.EclipticLong
	o.obliquity = sunPos.Obliquity
	o.rightAscension = sunPos.RightAscension
	o.Declination = sunPos.Declination
	o.equationOfTime = sunPos.EquationOfTime

	o.SunTransitTime = sunPos.SunTransitTime
}

func (o *Option) SetDate(date time.Time) {
	o.DateStart = time.Date(date.Year(), date.Month(), date.Day(), 12., 0, 0, 0, date.Location())

	o = o.calcJulianDay()

	o.fillSunPosition(o.julianDay, o.TimezoneOffset, o.Longitude)
}

func (o *Option) Now() {
	o.SetDate(time.Now())
}

func (o *Option) calcJulianDay() *Option {
	if o.DateStart.IsZero() {
		o.Now()
	}

	year := float64(o.DateStart.Year())
	month := float64(o.DateStart.Month())
	date := float64(o.DateStart.Day())

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

	o.julianDay = 1720994.5 + math.Floor(365.25*year) + math.Floor(30.6001*(month+1)) + b + date + (12-o.TimezoneOffset)/24

	return o
}
