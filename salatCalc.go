package moslemSalatSchedule

import (
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	salatEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/salat"
	"gitlab.com/naufalfmm/moslem-salat-schedule/err"
	"gitlab.com/naufalfmm/moslem-salat-schedule/model"
	"gitlab.com/naufalfmm/moslem-salat-schedule/option"
	"gitlab.com/naufalfmm/moslem-salat-schedule/option/salatOption"
)

func checkSalatOption(opt salatOption.SalatOption, defaultOpt option.CalcOpt, salat salatEnum.Salat) (salatOption.SalatOption, error) {
	if opt.Date.IsZero() {
		if defaultOpt.Date.IsZero() {
			return salatOption.SalatOption{}, err.ErrDateMissing
		}

		opt.Date = defaultOpt.Date
	}

	if opt.Timezone == 0 {
		return salatOption.SalatOption{}, err.ErrTimezoneMissing
	}

	if opt.Latitude.IsZero() {
		return salatOption.SalatOption{}, err.ErrLatitudeMissing
	}

	if opt.Longitude.IsZero() {
		return salatOption.SalatOption{}, err.ErrLongitudeMissing
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

	return opt, nil
}

func checkSalatOptionForFiveTimes(opt salatOption.SalatOption, defaultOpt option.CalcOpt) (salatOption.SalatOption, error) {
	var err error
	for _, salat := range salatEnum.GetAll() {
		opt, err = checkSalatOption(opt, defaultOpt, salat)
		if err != nil {
			return salatOption.SalatOption{}, err
		}
	}

	return opt, nil
}

func (i *impl) Fajr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option.CalcOpt, salatEnum.Fajr)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Salat: salatEnum.Fajr,
	}, nil
}

func (i *impl) Dhuhr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{
		Date:             i.option.CalcOpt.Date,
		FajrZenith:       i.option.CalcOpt.FajrZenith,
		IshaZenith:       i.option.CalcOpt.IshaZenith,
		IshaZenithType:   i.option.CalcOpt.IshaZenithType,
		SolarDeclination: i.option.CalcOpt.SolarDeclination,
		EquationOfTime:   i.option.CalcOpt.EquationOfTime,

		Latitude:  i.option.LocOpt.Latitude,
		Longitude: i.option.LocOpt.Longitude,
		Elevation: i.option.LocOpt.Elevation,
		Timezone:  i.option.LocOpt.Timezone,
	}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option.CalcOpt, salatEnum.Dhuhr)
	if err != nil {
		return model.SalatTime{}, err
	}

	angTime := salatOption.Longitude.Div(15).Neg().Add(angle.NewFromFloat64(12 + salatOption.Timezone)).Sub(angle.NewFromFloat64(salatOption.EquationOfTime))

	return model.SalatTime{
		Date:  salatOption.Date,
		Salat: salatEnum.Dhuhr,
		Time:  angTime.ToTime(),
	}, nil
}

func (i *impl) Asr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option.CalcOpt, salatEnum.Asr)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Salat: salatEnum.Asr,
	}, nil
}

func (i *impl) Maghrib(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option.CalcOpt, salatEnum.Maghrib)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Salat: salatEnum.Maghrib,
	}, nil
}

func (i *impl) Isha(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error) {
	salatOption := salatOption.SalatOption{}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option.CalcOpt, salatEnum.Isha)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Salat: salatEnum.Isha,
	}, nil
}

func (i *impl) AllFiveTimes(opts ...salatOption.ApplyingSalatOption) (model.FiveSalatTime, error) {
	salatOption := salatOption.SalatOption{}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOptionForFiveTimes(salatOption, i.option.CalcOpt)
	if err != nil {
		return model.FiveSalatTime{}, err
	}

	return model.FiveSalatTime{
		Date: salatOption.Date,
	}, nil
}
