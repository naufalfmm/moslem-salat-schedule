package moslemSalatSchedule

import (
	"github.com/naufalfmm/moslem-salat-schedule/schedule"
)

// type impl struct {
// 	option option.Option
// }

// func New(opts ...option.ApplyingOption) (MoslemSalatSchedule, error) {
// 	option := option.Option{
// 		RoundingTimeOption: roundingTimeOptionEnum.Default,
// 	}

// 	for _, opt := range opts {
// 		opt.Apply(&option)
// 	}

// 	if err := option.Validate(); err != nil {
// 		return nil, err
// 	}

// 	return &schedule.Schedule{
// 		option: option,
// 	}, nil
// }

func New(applyOpts ...schedule.ApplyCommOpt) (MoslemSalatSchedule, error) {
	opt := schedule.CommOpt{}

	for _, applyOpt := range applyOpts {
		applyOpt.Apply(&opt)
	}

	// if err := opt.Validate(); err != nil {
	// 	return nil, err
	// }

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return nil, err
	}

	return &schedule.Schedule{
		Opt: opt,
	}, nil
}
