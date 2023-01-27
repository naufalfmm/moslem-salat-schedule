package schedule

import (
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/moslem-salat-times/consts"
	salatEnum "github.com/naufalfmm/moslem-salat-times/enum/salat"
	sunZenithEnum "github.com/naufalfmm/moslem-salat-times/enum/sunZenith"
	"github.com/naufalfmm/moslem-salat-times/model"
	"github.com/naufalfmm/moslem-salat-times/option"
	"github.com/naufalfmm/moslem-salat-times/utils/sunPositions"
)

func sunriseAngleTime(opt option.Option, sunPos sunPositions.SunPosition) angle.Angle {
	return sunPos.SunTransitTime.Sub(opt.CalculateSunriseSunsetHighAltitude(sunPos.Declination))
}

func sunsetAngleTime(opt option.Option, sunPos sunPositions.SunPosition) angle.Angle {
	return sunPos.SunTransitTime.Add(opt.CalculateSunriseSunsetHighAltitude(sunPos.Declination))
}

func maghribAngleTime(opt option.Option, sunPos sunPositions.SunPosition) angle.Angle {
	return sunsetAngleTime(opt, sunPos).Add(angle.NewDegreeFromFloat(consts.MaghribSlightMarginMinute / 60.))
}

func (s *Schedule) Midnight(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Isha); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		dateStart, _ := opt.GetDateRange()
		dateStart = dateStart.Add(-24 * time.Hour)

		yestSundayOpt := opt
		yestSundayOpt, err := yestSundayOpt.SetDateRange(dateStart, dateStart).CalculateSunPositions()
		if err != nil {
			return nil, err
		}

		yestSunset := sunsetAngleTime(yestSundayOpt, yestSundayOpt.GetSunPositions()[0])
		todaySunrise := sunriseAngleTime(opt, sunPosition)

		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Midnight,
			Time:  opt.RoundTime(yestSunset.Add(angle.NewFromDegreeMinuteSecond(24., 0., 0.).ToDegree().Sub(yestSunset).Add(todaySunrise).Div(2.)).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Fajr(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Fajr); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Fajr,
			Time:  opt.RoundTime(sunPosition.SunTransitTime.Sub(opt.CalculateFajrHighAltitude(sunPosition.Declination)).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Sunrise(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Fajr); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Sunrise,
			Time:  opt.RoundTime(sunriseAngleTime(opt, sunPosition).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Dhuhr(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Dhuhr); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Dhuhr,
			Time:  opt.RoundTime(sunPosition.SunTransitTime.AddScalar(consts.DhuhrSlightMarginMinute / 60.).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Asr(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Asr); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Asr,
			Time:  opt.RoundTime(sunPosition.SunTransitTime.Add(opt.CalculateAsrAngle(sunPosition.Declination)).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Sunset(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Maghrib); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Sunset,
			Time:  opt.RoundTime(sunsetAngleTime(opt, sunPosition).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Maghrib(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Maghrib); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Maghrib,
			Time:  opt.RoundTime(maghribAngleTime(opt, sunPosition).ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) Isha(opt option.Option) (model.PeriodicSalatTime, error) {
	if err := opt.ValidateBySalat(salatEnum.Isha); err != nil {
		return model.PeriodicSalatTime{}, err
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return model.PeriodicSalatTime{}, err
	}

	periodicSalatTimes := make(model.PeriodicSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		ishaHighAlt, ishaType := opt.CalculateIshaHighAltitude(sunPosition.Declination)

		angTime := angle.Angle{}
		if ishaType == sunZenithEnum.Standard {
			angTime = sunPosition.SunTransitTime.Add(ishaHighAlt)
		}

		if ishaType == sunZenithEnum.AfterMagrib {
			angTime = maghribAngleTime(opt, sunPosition).Add(ishaHighAlt)
		}

		periodicSalatTimes[i] = model.SalatTime{
			Date:  sunPosition.Date,
			Salat: salatEnum.Isha,
			Time:  opt.RoundTime(angTime.ToTime()),
		}
	}

	return periodicSalatTimes, nil
}

func (s *Schedule) AllTimes(opt option.Option) (model.PeriodicAllSalatTime, error) {
	if err := opt.ValidateBySalat(0); err != nil {
		return model.PeriodicAllSalatTime{}, err
	}

	periodicAllSalatTimes := make(model.PeriodicAllSalatTime, len(opt.GetSunPositions()))
	for i, sunPosition := range opt.GetSunPositions() {
		dateOpt := opt
		dateOpt, err := dateOpt.SetDateRange(sunPosition.Date, sunPosition.Date).CalculateSunPositions()
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		midnight, err := s.Midnight(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		fajr, err := s.Fajr(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		sunrise, err := s.Sunrise(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		dhuhr, err := s.Dhuhr(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		asr, err := s.Asr(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		sunset, err := s.Sunset(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		maghrib, err := s.Maghrib(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		isha, err := s.Isha(dateOpt)
		if err != nil {
			return model.PeriodicAllSalatTime{}, err
		}

		periodicAllSalatTimes[i] = model.AllSalatTime{
			Date: sunPosition.Date,
			SalatTimes: []model.SalatTime{
				midnight[0],
				fajr[0],
				sunrise[0],
				dhuhr[0],
				asr[0],
				sunset[0],
				maghrib[0],
				isha[0],
			},
		}
	}

	return periodicAllSalatTimes, nil
}
