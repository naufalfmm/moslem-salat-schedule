package moslemSalatSchedule

import (
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/option"
)

type impl struct {
	option option.Option
}

func New(opts ...option.ApplyingOption) (MoslemSalatSchedule, error) {
	option := option.Option{}

	for _, opt := range opts {
		opt.Apply(&option)
	}

	if err := option.Validate(); err != nil {
		return nil, err
	}

	return &impl{
		option: option,
	}, nil
}

func (i *impl) SetDate(date time.Time) MoslemSalatSchedule {
	i.option.CalcOpt = i.option.CalcOpt.SetDate(date)

	return i
}

func (i *impl) Now() MoslemSalatSchedule {
	i.option.CalcOpt = i.option.CalcOpt.Now()

	return i
}
