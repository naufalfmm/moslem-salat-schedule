package sunPosition

import (
	"math"

	"github.com/naufalfmm/angle"
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
	eclipticLong := angle.NewDegreeFromFloat(1.915).
		Mul(math.Sin(meanAnomaly.ToRadian().ToFloat())).
		Add(angle.NewDegreeFromFloat(0.08).Mul(math.Sin(meanAnomaly.ToRadian().Mul(2).ToFloat()))).
		Add(meanLongSun).FullRotate()
	obliquity := angle.NewDegreeFromFloat(23.439 + 0.00000036*julianDate).FullRotate()
	rightAscension := angle.NewRadianFromFloat(math.Atan2(math.Cos(obliquity.ToRadian().ToFloat())*math.Sin(eclipticLong.ToRadian().ToFloat()), math.Cos(eclipticLong.ToRadian().ToFloat())))
	declination := angle.NewRadianFromFloat(math.Asin(math.Sin(obliquity.ToRadian().ToFloat()) * math.Sin(eclipticLong.ToRadian().ToFloat())))

	equationOfTime := meanLongSun.ToDegree().Sub(rightAscension.ToDegree())
	if equationOfTime.GreatherThan(angle.NewDegreeFromFloat(50.)) {
		equationOfTime = equationOfTime.Sub(angle.NewDegreeFromFloat(360))
	}

	SunTransitTime := angle.NewDegreeFromFloat(12 + timezone).Sub(longitude.Div(15.)).Sub(equationOfTime.Mul(4).Div(60))

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
