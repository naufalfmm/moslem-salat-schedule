package moslemSalatSchedule

import (
	roundingTimeOptionEnum "github.com/naufalfmm/moslem-salat-schedule/enum/roundingTimeOption"
	"github.com/naufalfmm/moslem-salat-schedule/option"
)

type impl struct {
	option option.Option
}

func New(opts ...option.ApplyingOption) (MoslemSalatSchedule, error) {
	option := option.Option{
		RoundingTimeOption: roundingTimeOptionEnum.Default,
	}

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
