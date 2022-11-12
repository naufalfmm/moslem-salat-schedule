package option

import (
	"math"
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	sunZenithEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type CalcOpt struct {
	Date time.Time

	FajrZenith     angle.Angle
	IshaZenith     angle.Angle
	IshaZenithType sunZenithEnum.IshaZenithType

	SolarDeclination float64
	EquationOfTime   float64

	julianDate                  float64
	julianDay                   float64
	solarMeanAnomaly            float64
	solarMeanLong               float64
	solarGeocentricEclipticLong float64
	sunEarthRadius              float64
	earthEclipticTilt           float64
	solarRightAscension         float64
}

func (c CalcOpt) SetDate(date time.Time) CalcOpt {
	c.Date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	c = c.calcJulianDay()

	c.julianDate = c.julianDay - 2451545.0
	// g = 357.529 + 0.98560028* d;
	// q = 280.459 + 0.98564736* d;
	// L = q + 1.915* sin(g) + 0.020* sin(2*g);
	c.solarMeanAnomaly = 357.529 + 0.98560028*c.julianDate
	c.solarMeanLong = 280.459 + 0.98564736*c.julianDate
	c.solarGeocentricEclipticLong = c.solarMeanLong +
		1.915*math.Sin(c.solarMeanAnomaly) +
		0.020*math.Sin(2.0*c.solarMeanAnomaly)
	// R = 1.00014 – 0.01671* cos(g) – 0.00014* cos(2*g);
	// e = 23.439 – 0.00000036* d;
	// RA = arctan2(cos(e)* sin(L), cos(L))/ 15;
	c.sunEarthRadius = 1.00014 -
		0.01671*math.Cos(c.solarMeanAnomaly) -
		0.00014*math.Cos(2.0*c.solarMeanAnomaly)
	c.earthEclipticTilt = 23.439 - 0.00000036*c.julianDate
	c.solarRightAscension = math.Atan2(
		math.Cos(c.earthEclipticTilt)*math.Sin(c.solarGeocentricEclipticLong),
		math.Cos(c.solarGeocentricEclipticLong),
	) / 15.0

	return c
}

func (c CalcOpt) Now() CalcOpt {
	return c.SetDate(time.Now())
}

func (c CalcOpt) calcJulianDay() CalcOpt {
	if c.Date.IsZero() {
		c = c.Now()
	}

	year := float64(c.Date.Year())
	month := float64(c.Date.Month())
	day := float64(c.Date.Day())

	if month == 1 || month == 2 {
		year = year - 1
		month = month + 12
	}

	a := math.Floor(year / 100.0)
	b := 2.0 - a + math.Floor(a/4.0)

	e := math.Floor(36525.0 * ((year + 4716.0) / 100.0))
	f := math.Floor(306.0*((month+1.0)/10.0)) + b

	c.julianDate = e + f + day - 1524.5

	return c
}
