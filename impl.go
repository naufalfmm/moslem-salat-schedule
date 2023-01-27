package moslemSalatTimes

import (
	"github.com/naufalfmm/moslem-salat-times/schedule"
)

func New(applyOpts ...schedule.ApplyCommOpt) (MoslemSalatTimes, error) {
	opt := schedule.CommOpt{}

	for _, applyOpt := range applyOpts {
		applyOpt.Apply(&opt)
	}

	opt, err := opt.CalculateSunPositions()
	if err != nil {
		return nil, err
	}

	return &schedule.Schedule{
		Opt: opt,
	}, nil
}
