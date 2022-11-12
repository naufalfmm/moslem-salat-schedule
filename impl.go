package moslemSalatSchedule

import (
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
