package salatOption

import (
	"math"
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	"gitlab.com/naufalfmm/moslem-salat-schedule/utils/sunPosition"
)

func (s *SalatOption) fillSunPosition(julianDay, timezone float64, longitude angle.Angle) {
	sunPos := sunPosition.CalSunPosition(julianDay, timezone, longitude)

	s.julianDate = sunPos.JulianDate

	s.meanAnomaly = sunPos.MeanAnomaly
	s.meanLongSun = sunPos.MeanLongSun
	s.eclipticLong = sunPos.EclipticLong
	s.obliquity = sunPos.Obliquity
	s.rightAscension = sunPos.RightAscension
	s.Declination = sunPos.Declination
	s.equationOfTime = sunPos.EquationOfTime

	s.SunTransitTime = sunPos.SunTransitTime
}

func (s *SalatOption) SetDate(date time.Time) {
	s.Date = time.Date(date.Year(), date.Month(), date.Day(), 12., 0, 0, 0, date.Location())

	s = s.calcJulianDay()

	s.fillSunPosition(s.julianDay, s.Timezone, s.Longitude)
}

func (s *SalatOption) Now() {
	s.SetDate(time.Now())
}

func (s *SalatOption) calcJulianDay() *SalatOption {
	if s.Date.IsZero() {
		s.Now()
	}

	year := float64(s.Date.Year())
	month := float64(s.Date.Month())
	date := float64(s.Date.Day())

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

	s.julianDay = 1720994.5 + math.Floor(365.25*year) + math.Floor(30.6001*(month+1)) + b + date + (float64(s.Date.Hour())-s.Timezone)/24

	return s
}
