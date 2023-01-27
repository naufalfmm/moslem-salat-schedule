package sunPositions

import (
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/angle/trig"
	"github.com/naufalfmm/moslem-salat-schedule/consts"
	"github.com/naufalfmm/moslem-salat-schedule/utils/julian"
)

type (
	SunPosition struct {
		Date time.Time

		JulianDay  float64
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

	SunPositions []SunPosition
)

func NewFromDateRange(dateStart, dateEnd time.Time, loc *time.Location, longitude angle.Angle) SunPositions {
	dateSunPoss := make(SunPositions, int(dateEnd.Sub(dateStart).Hours()/24.)+1)

	for i := 0; i < int(dateEnd.Sub(dateStart).Hours()/24.)+1; i++ {
		date := dateStart.AddDate(0, 0, i)

		dateSunPoss[i] = calSunPositionByDate(date, loc, longitude)
	}

	return dateSunPoss
}

func calSunPositionByDate(date time.Time, loc *time.Location, longitude angle.Angle) SunPosition {
	dateSunPos := SunPosition{}

	dateSunPos.Date = time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, loc)
	dateSunPos.JulianDay = julian.GregorianToJulianUTC(dateSunPos.Date)
	dateSunPos.JulianDate = dateSunPos.JulianDay - 2451545.

	dateSunPos.MeanAnomaly = angle.NewDegreeFromFloat(357.529 + 0.98560028*dateSunPos.JulianDate).FullRotate()
	dateSunPos.MeanLongSun = angle.NewDegreeFromFloat(280.459 + 0.98564736*dateSunPos.JulianDate).FullRotate()
	dateSunPos.EclipticLong = dateSunPos.MeanLongSun.AddScalar(1.915*trig.Sin(dateSunPos.MeanAnomaly) + 0.02*trig.Sin(dateSunPos.MeanAnomaly.Mul(2.))).FullRotate()
	dateSunPos.Obliquity = angle.NewDegreeFromFloat(23.439 - 0.00000036*dateSunPos.JulianDate).FullRotate()
	dateSunPos.RightAscension = trig.Atan2(trig.Cos(dateSunPos.Obliquity)*trig.Sin(dateSunPos.EclipticLong), trig.Cos(dateSunPos.EclipticLong))
	dateSunPos.Declination = trig.Asin(trig.Sin(dateSunPos.Obliquity) * trig.Sin(dateSunPos.EclipticLong))

	dateSunPos.EquationOfTime = dateSunPos.MeanLongSun.Sub(dateSunPos.RightAscension)
	if dateSunPos.EquationOfTime.GreatherThan(angle.NewDegreeFromFloat(50.)) {
		dateSunPos.EquationOfTime = dateSunPos.EquationOfTime.SubScalar(360.)
	}

	_, offset := time.Now().In(loc).Zone()

	dateSunPos.SunTransitTime = longitude.Div(15.).Neg().Sub(dateSunPos.EquationOfTime.Mul(4.).Div(60.)).AddScalar(12.).AddScalar(float64(offset) / consts.OffsetTimezone)

	return dateSunPos
}
