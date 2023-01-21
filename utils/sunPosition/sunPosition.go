package sunPosition

import (
	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/angle/trig"
)

type SunPosition struct {
	JulianDate float64

	MeanAnomaly    angle.Angle
	MeanLongSun    angle.Angle
	EclipticLong   angle.Angle
	Obliquity      angle.Angle
	RightAscension angle.Angle
	Declination    angle.Angle
	EquationOfTime angle.Angle

	SunTransitTime angle.Angle
}

func CalSunPosition(julianDay, timezone float64, longitude angle.Angle) SunPosition {
	julianDate := julianDay - 2451545.

	meanAnomaly := angle.NewDegreeFromFloat(357.259 + 0.98560028*julianDate).FullRotate()
	meanLongSun := angle.NewDegreeFromFloat(280.459 + 0.98564736*julianDate).FullRotate()
	eclipticLong := meanLongSun.AddScalar(1.915*trig.Sin(meanAnomaly) + 0.02*trig.Sin(meanAnomaly.Mul(2.))).FullRotate()
	obliquity := angle.NewDegreeFromFloat(23.439 + 0.00000036*julianDate).FullRotate()
	rightAscension := trig.Atan2(trig.Cos(obliquity)*trig.Sin(eclipticLong), trig.Cos(eclipticLong))
	declination := trig.Asin(trig.Sin(obliquity) * trig.Sin(eclipticLong))

	equationOfTime := meanLongSun.Sub(rightAscension)
	if equationOfTime.GreatherThan(angle.NewDegreeFromFloat(50.)) {
		equationOfTime = equationOfTime.SubScalar(360.)
	}

	SunTransitTime := longitude.Div(15.).Neg().Sub(equationOfTime.Mul(4.).Div(60.)).AddScalar(12.).AddScalar(timezone)

	return SunPosition{
		JulianDate:     julianDate,
		MeanAnomaly:    meanAnomaly,
		MeanLongSun:    meanLongSun,
		EclipticLong:   eclipticLong,
		Obliquity:      obliquity,
		RightAscension: rightAscension,
		Declination:    declination,
		EquationOfTime: equationOfTime,
		SunTransitTime: SunTransitTime,
	}
}
