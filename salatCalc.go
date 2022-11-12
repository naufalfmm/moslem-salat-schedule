package moslemSalatSchedule

import (
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
	salatOption := salatOption.SalatOption{}

	for _, opt := range opts {
		opt.Apply(&salatOption)
	}

	salatOption, err := checkSalatOption(salatOption, i.option.CalcOpt, salatEnum.Dhuhr)
	if err != nil {
		return model.SalatTime{}, err
	}

	return model.SalatTime{
		Salat: salatEnum.Dhuhr,
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
