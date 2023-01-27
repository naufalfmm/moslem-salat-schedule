package moslemSalatSchedule

import (
	"github.com/naufalfmm/moslem-salat-schedule/schedule"
)

func New(applyOpts ...schedule.ApplyCommOpt) (MoslemSalatSchedule, error) {
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
