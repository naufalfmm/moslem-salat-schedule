package salatOption

import (
	"time"

	"github.com/shopspring/decimal"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	sunZenithEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
	"gitlab.com/naufalfmm/moslem-salat-schedule/utils"
)

type SalatOption struct {
	Date time.Time

	julianDate                  decimal.Decimal
	julianDay                   decimal.Decimal
	solarMeanAnomaly            decimal.Decimal
	solarMeanLong               decimal.Decimal
	solarGeocentricEclipticLong decimal.Decimal
	sunEarthRadius              decimal.Decimal
	earthEclipticTilt           decimal.Decimal
	solarRightAscension         decimal.Decimal

	FajrZenith     angle.Angle
	IshaZenith     angle.Angle
	IshaZenithType sunZenithEnum.IshaZenithType

	SolarDeclination decimal.Decimal
	EquationOfTime   decimal.Decimal
}

func (s SalatOption) SetDate(date time.Time) SalatOption {
	s.Date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	s = s.calcJulianDay()

	s.julianDate = s.julianDay.Sub(decimal.NewFromFloat(2451545.0))
	// g = 357.529 + 0.98560028* d;
	// q = 280.459 + 0.98564736* d;
	// L = q + 1.915* sin(g) + 0.020* sin(2*g);
	s.solarMeanAnomaly = decimal.NewFromFloat(357.529).Add(decimal.NewFromFloat(0.98560028).Mul(s.julianDate))
	s.solarMeanLong = decimal.NewFromFloat(280.459).Add(decimal.NewFromFloat(0.98564736).Mul(s.julianDate))
	s.solarGeocentricEclipticLong = s.solarMeanLong.
		Add(decimal.NewFromFloat(1.915).Mul(s.solarMeanAnomaly.Sin())).
		Add(decimal.NewFromFloat(0.020).Mul(decimal.NewFromInt(2).Mul(s.solarMeanAnomaly.Sin())))
	// R = 1.00014 – 0.01671* cos(g) – 0.00014* cos(2*g);
	// e = 23.439 – 0.00000036* d;
	// RA = arctan2(cos(e)* sin(L), cos(L))/ 15;
	s.sunEarthRadius = decimal.NewFromFloat(1.00014).
		Sub(decimal.NewFromFloat(0.01671).Mul(s.solarMeanAnomaly.Cos())).
		Sub(decimal.NewFromFloat(0.00014).Mul(decimal.NewFromInt(2).Mul(s.solarMeanAnomaly.Cos())))
	s.earthEclipticTilt = decimal.NewFromFloat(23.439).Sub(decimal.NewFromFloat(0.00000036).Mul(s.julianDate))
	s.solarRightAscension = utils.Atan2(
		s.earthEclipticTilt.Cos().Mul(s.solarGeocentricEclipticLong.Sin()),
		s.solarGeocentricEclipticLong.Cos(),
	).Div(decimal.NewFromFloat(15.0))

	return s
}

func (s SalatOption) Now() SalatOption {
	return s.SetDate(time.Now())
}

func (s SalatOption) calcJulianDay() SalatOption {
	if s.Date.IsZero() {
		s = s.Now()
	}

	year := decimal.NewFromInt(int64(s.Date.Year()))
	month := decimal.NewFromInt(int64(s.Date.Month()))
	day := decimal.NewFromInt(int64(s.Date.Day()))

	if month.Equal(decimal.NewFromInt(1)) || month.Equal(decimal.NewFromInt(2)) {
		year = year.Sub(decimal.NewFromInt(1))
		month = month.Add(decimal.NewFromInt(12))
	}

	a := year.Div(decimal.NewFromInt(100)).Floor()
	b := decimal.NewFromInt(2).Sub(a).Add(a.Div(decimal.NewFromInt(4)).Floor())

	e := decimal.NewFromInt(36525).Mul(year.Add(decimal.NewFromInt(4716)).Div(decimal.NewFromInt(100))).Floor()
	f := decimal.NewFromInt(306).Mul(month.Add(decimal.NewFromInt(1)).Div(decimal.NewFromInt(10))).Floor().Add(b)

	s.julianDate = e.Add(f).Add(day).Sub(decimal.NewFromFloat(1524.5))

	return s
}
