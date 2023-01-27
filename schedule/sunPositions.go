package schedule

import (
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/angle/trig"
	"github.com/naufalfmm/moslem-salat-schedule/consts"
	"github.com/naufalfmm/moslem-salat-schedule/utils/julian"
)

type (
	dateSunPosition struct {
		date time.Time

		julianDay  float64
		julianDate float64

		meanAnomaly    angle.Angle
		meanLongSun    angle.Angle
		eclipticLong   angle.Angle
		obliquity      angle.Angle
		rightAscension angle.Angle
		declination    angle.Angle
		equationOfTime angle.Angle

		sunTransitTime angle.Angle
	}

	dateSunPositions []dateSunPosition
)

func newFromDateRange(dateStart, dateEnd time.Time, loc *time.Location, longitude angle.Angle) dateSunPositions {
	dateSunPoss := make(dateSunPositions, int(dateEnd.Sub(dateStart))+1)

	for i := 0; i <= int(dateEnd.Sub(dateStart)); i++ {
		date := dateStart.AddDate(0, 0, i)

		dateSunPoss[i] = calSunPositionByDate(date, loc, longitude)
	}

	return dateSunPoss
}

func calSunPositionByDate(date time.Time, loc *time.Location, longitude angle.Angle) dateSunPosition {
	dateSunPos := dateSunPosition{}

	dateSunPos.date = time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, loc)
	dateSunPos.julianDay = julian.GregorianToJulianUTC(dateSunPos.date)
	dateSunPos.julianDate = dateSunPos.julianDay - 2451545.

	dateSunPos.meanAnomaly = angle.NewDegreeFromFloat(357.529 + 0.98560028*dateSunPos.julianDate).FullRotate()
	dateSunPos.meanLongSun = angle.NewDegreeFromFloat(280.459 + 0.98564736*dateSunPos.julianDate).FullRotate()
	dateSunPos.eclipticLong = dateSunPos.meanLongSun.AddScalar(1.915*trig.Sin(dateSunPos.meanAnomaly) + 0.02*trig.Sin(dateSunPos.meanAnomaly.Mul(2.))).FullRotate()
	dateSunPos.obliquity = angle.NewDegreeFromFloat(23.439 - 0.00000036*dateSunPos.julianDate).FullRotate()
	dateSunPos.rightAscension = trig.Atan2(trig.Cos(dateSunPos.obliquity)*trig.Sin(dateSunPos.eclipticLong), trig.Cos(dateSunPos.eclipticLong))
	dateSunPos.declination = trig.Asin(trig.Sin(dateSunPos.obliquity) * trig.Sin(dateSunPos.eclipticLong))

	dateSunPos.equationOfTime = dateSunPos.meanLongSun.Sub(dateSunPos.rightAscension)
	if dateSunPos.equationOfTime.GreatherThan(angle.NewDegreeFromFloat(50.)) {
		dateSunPos.equationOfTime = dateSunPos.equationOfTime.SubScalar(360.)
	}

	_, offset := time.Now().In(loc).Zone()

	dateSunPos.sunTransitTime = longitude.Div(15.).Neg().Sub(dateSunPos.equationOfTime.Mul(4.).Div(60.)).AddScalar(12.).AddScalar(float64(offset) / consts.OffsetTimezone)

	return dateSunPos
}
