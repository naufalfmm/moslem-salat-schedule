package moslemSalatSchedule

import (
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/angle/trig"
	"github.com/naufalfmm/moslem-salat-schedule/consts"
	roundingTimeOptionEnum "github.com/naufalfmm/moslem-salat-schedule/enum/roundingTimeOption"
	salatEnum "github.com/naufalfmm/moslem-salat-schedule/enum/salat"
	sunZenithEnum "github.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
	"github.com/naufalfmm/moslem-salat-schedule/err"
	"github.com/naufalfmm/moslem-salat-schedule/model"
	"github.com/naufalfmm/moslem-salat-schedule/option"
	"github.com/naufalfmm/moslem-salat-schedule/option/salatOption"
	"github.com/naufalfmm/moslem-salat-schedule/utils/salatHighAltitude"
)

func checkSalatOption(opt salatOption.SalatOption, defaultOpt option.Option, salat salatEnum.Salat) (salatOption.SalatOption, error) {
	if opt.Date.IsZero() {
		if defaultOpt.Date.IsZero() {
			return salatOption.SalatOption{}, err.ErrDateMissing
		}

		opt.Date = defaultOpt.Date
	}

	if opt.Latitude.IsZero() {
		if defaultOpt.Latitude.IsZero() {
			return salatOption.SalatOption{}, err.ErrLatitudeMissing
		}

		opt.Latitude = defaultOpt.Latitude
	}

	if opt.Longitude.IsZero() {
		if defaultOpt.Longitude.IsZero() {
			return salatOption.SalatOption{}, err.ErrLongitudeMissing
		}

		opt.Longitude = defaultOpt.Longitude
	}

	if opt.RoundingTimeOption == 0 {
		if defaultOpt.RoundingTimeOption == 0 {
			opt.RoundingTimeOption = roundingTimeOptionEnum.Default
		}

		opt.RoundingTimeOption = defaultOpt.RoundingTimeOption
	}

	if salat == salatEnum.Fajr {
		if opt.FajrZenith.IsZero() {
			if defaultOpt.FajrZenith.IsZero() {
				return salatOption.SalatOption{}, err.ErrFajrZenithMissing
			}

			opt.FajrZenith = defaultOpt.FajrZenith
		}
	}

	if salat == salatEnum.Isha {
		if opt.IshaZenith.IsZero() {
			if defaultOpt.IshaZenith.IsZero() {
				return salatOption.SalatOption{}, err.ErrIshaZenithMissing
			}

			opt.IshaZenith = defaultOpt.IshaZenith
		}
	}

	if salat == salatEnum.Asr {
		if opt.AsrMazhab == 0 {
			if defaultOpt.AsrMazhab == 0 {
				return salatOption.SalatOption{}, err.ErrMazhabMissing
			}

			opt.AsrMazhab = defaultOpt.AsrMazhab
		}
	}

	return opt, nil
}

func checkSalatOptionForAllTimes(opt salatOption.SalatOption, defaultOpt option.Option) (salatOption.SalatOption, error) {
	var err error
	for _, salat := range salatEnum.GetAll() {
		opt, err = checkSalatOption(opt, defaultOpt, salat)
		if err != nil {
			return salatOption.SalatOption{}, err
		}
	}

	return opt, nil
}

func sunriseAngleTime(salatOption salatOption.SalatOption) angle.Angle {
	return salatOption.SunTransitTime.Sub(salatHighAltitude.CalcSalatHighAltitude(angle.NewDegreeFromFloat(consts.SunriseSunsetAngleFactor), salatOption.Latitude, salatOption.Declination, salatOption.Elevation))
}

func sunsetAngleTime(salatOption salatOption.SalatOption) angle.Angle {
	return salatOption.SunTransitTime.Add(salatHighAltitude.CalcSalatHighAltitude(angle.NewDegreeFromFloat(consts.SunriseSunsetAngleFactor), salatOption.Latitude, salatOption.Declination, salatOption.Elevation))
}

func (i *impl) Midnight(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOpt := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOpt)
	}

	salatOpt, err := checkSalatOption(salatOpt, i.option, salatEnum.Isha)
	if err != nil {
		return model.SalatTime{}, err
	}

	yesterdaySunsetOpt := salatOpt
	yesterdaySunsetOpt.Date = yesterdaySunsetOpt.Date.Add(-24 * time.Hour)
	yesterdaySunset := sunsetAngleTime(yesterdaySunsetOpt)

	todaySunrise := sunriseAngleTime(salatOpt)

	durRange := angle.NewFromDegreeMinuteSecond(24., 0., 0.).ToDegree().Sub(yesterdaySunset).Add(todaySunrise).Div(2.)

	return model.SalatTime{
		Date:  salatOpt.Date,
		Salat: salatEnum.Midnight,
		Time:  salatOpt.RoundingTimeOption.RoundTime(yesterdaySunset.Add(durRange).ToTime()),
	}, nil
}

func (i *impl) Fajr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Fajr)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Fajr,
		Time:  salatOption.RoundingTimeOption.RoundTime(salatOption.SunTransitTime.Sub(salatHighAltitude.CalcSalatHighAltitude(salatOption.FajrZenith, salatOption.Latitude, salatOption.Declination, salatOption.Elevation)).ToTime()),
	}, nil
}

func (i *impl) Sunrise(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Fajr)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Sunrise,
		Time:  salatOption.RoundingTimeOption.RoundTime(sunriseAngleTime(salatOption).ToTime()),
	}, nil
}

func (i *impl) Dhuhr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Dhuhr)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Dhuhr,
		Time:  salatOption.RoundingTimeOption.RoundTime(salatOption.SunTransitTime.AddScalar(consts.DhuhrSlightMarginMinute / 60.).ToTime()),
	}, nil
}

func (i *impl) Asr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,
		AsrMazhab:      i.option.AsrMazhab,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Asr)
	if err != nil {
		return model.SalatTime{}, err
	}

	asrFactorAng := trig.Acos((trig.Sin(trig.Acot(salatOption.AsrMazhab.AsrShadowLength()+trig.Tan(salatOption.Latitude.Sub(salatOption.Declination).Abs()))) - (trig.Sin(salatOption.Latitude) * trig.Sin(salatOption.Declination))) / (trig.Cos(salatOption.Latitude) * trig.Cos(salatOption.Declination))).Div(15.)

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Asr,
		Time:  salatOption.RoundingTimeOption.RoundTime(salatOption.SunTransitTime.Add(asrFactorAng).ToTime()),
	}, nil
}

func (i *impl) Sunset(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Maghrib)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Sunset,
		Time:  salatOption.RoundingTimeOption.RoundTime(sunsetAngleTime(salatOption).ToTime()),
	}, nil
}

func (i *impl) Maghrib(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Maghrib)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Maghrib,
		Time:  salatOption.RoundingTimeOption.RoundTime(sunsetAngleTime(salatOption).Add(angle.NewDegreeFromFloat(consts.MaghribSlightMarginMinute / 60.)).ToTime()),
	}, nil
}

func (i *impl) Isha(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		Declination:    i.option.Declination,
		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option, salatEnum.Isha)
	if err != nil {
		return model.SalatTime{}, err
	}

	angTime := angle.Angle{}
	if salatOption.IshaZenithType == sunZenithEnum.Standard {
		angTime = salatOption.SunTransitTime.Add(salatHighAltitude.CalcSalatHighAltitude(salatOption.IshaZenith, salatOption.Latitude, salatOption.Declination, salatOption.Elevation))
	}

	if salatOption.IshaZenithType == sunZenithEnum.AfterMagrib {
		angTime = sunsetAngleTime(salatOption).Add(salatOption.IshaZenith)
	}

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Isha,
		Time:  salatOption.RoundingTimeOption.RoundTime(angTime.ToTime()),
	}, nil
}

func (i *impl) AllTimes(opts ...salatOption.ApplyingSalatOption) (model.AllSalatTimes, error) {
	salatOpt := salatOption.SalatOption{
		Date:           i.option.Date,
		FajrZenith:     i.option.FajrZenith,
		IshaZenith:     i.option.IshaZenith,
		IshaZenithType: i.option.IshaZenithType,

		Latitude:       i.option.Latitude,
		Longitude:      i.option.Longitude,
		Elevation:      i.option.Elevation,
		TimezoneOffset: i.option.TimezoneOffset,

		RoundingTimeOption: i.option.RoundingTimeOption,

		SunTransitTime: i.option.SunTransitTime,
	}

	for _, opt := range opts {
		opt.Apply(&salatOpt)
	}

	salatOpt, err := checkSalatOptionForAllTimes(salatOpt, i.option)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	midnight, err := i.Midnight(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	fajr, err := i.Fajr(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	sunrise, err := i.Sunrise(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	dhuhr, err := i.Dhuhr(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	asr, err := i.Asr(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	sunset, err := i.Sunset(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	maghrib, err := i.Maghrib(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	isha, err := i.Isha(opts...)
	if err != nil {
		return model.AllSalatTimes{}, err
	}

	return model.AllSalatTimes{
		Date: salatOpt.Date,
		SalatTimes: []model.SalatTime{
			midnight,
			fajr,
			sunrise,
			dhuhr,
			asr,
			sunset,
			maghrib,
			isha,
		},
	}, nil
}
