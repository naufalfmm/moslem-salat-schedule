package salatOption

import (
	"math"
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	sunZenithEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type SalatOption struct {
	Date time.Time

	julianDate                  float64
	julianDay                   float64
	solarMeanAnomaly            float64
	solarMeanLong               float64
	solarGeocentricEclipticLong float64
	sunEarthRadius              float64
	earthEclipticTilt           float64
	solarRightAscension         float64

	FajrZenith     angle.Angle
	IshaZenith     angle.Angle
	IshaZenithType sunZenithEnum.IshaZenithType

	SolarDeclination float64
	EquationOfTime   float64
}

func (s SalatOption) SetDate(date time.Time) SalatOption {
	s.Date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	s = s.calcJulianDay()

	s.julianDate = s.julianDay - 2451545.0
	// g = 357.529 + 0.98560028* d;
	// q = 280.459 + 0.98564736* d;
	// L = q + 1.915* sin(g) + 0.020* sin(2*g);
	s.solarMeanAnomaly = 357.529 + 0.98560028*s.julianDate
	s.solarMeanLong = 280.459 + 0.98564736*s.julianDate
	s.solarGeocentricEclipticLong = s.solarMeanLong +
		1.915*math.Sin(s.solarMeanAnomaly) +
		0.020*math.Sin(2.0*s.solarMeanAnomaly)
	// R = 1.00014 – 0.01671* cos(g) – 0.00014* cos(2*g);
	// e = 23.439 – 0.00000036* d;
	// RA = arctan2(cos(e)* sin(L), cos(L))/ 15;
	s.sunEarthRadius = 1.00014 -
		0.01671*math.Cos(s.solarMeanAnomaly) -
		0.00014*math.Cos(2.0*s.solarMeanAnomaly)
	s.earthEclipticTilt = 23.439 - 0.00000036*s.julianDate
	s.solarRightAscension = math.Atan2(
		math.Cos(s.earthEclipticTilt)*math.Sin(s.solarGeocentricEclipticLong),
		math.Cos(s.solarGeocentricEclipticLong),
	) / 15.0

	return s
}

func (s SalatOption) Now() SalatOption {
	return s.SetDate(time.Now())
}

func (s SalatOption) calcJulianDay() SalatOption {
	if s.Date.IsZero() {
		s = s.Now()
	}

	year := float64(s.Date.Year())
	month := float64(s.Date.Month())
	day := float64(s.Date.Day())

	if month == 1 || month == 2 {
		year = year - 1
		month = month + 12
	}

	a := math.Floor(year / 100.0)
	b := 2.0 - a + math.Floor(a/4.0)

	e := math.Floor(36525.0 * ((year + 4716.0) / 100.0))
	f := math.Floor(306.0*((month+1.0)/10.0)) + b

	s.julianDate = e + f + day - 1524.5

	return s
}
