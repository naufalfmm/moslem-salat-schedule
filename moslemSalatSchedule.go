package moslemSalatSchedule

import (
	"github.com/naufalfmm/moslem-salat-schedule/model"
	"github.com/naufalfmm/moslem-salat-schedule/option"
)

type MoslemSalatSchedule interface {
	Midnight(opt option.Option) (model.PeriodicSalatTime, error)
	Fajr(opt option.Option) (model.PeriodicSalatTime, error)
	Sunrise(opt option.Option) (model.PeriodicSalatTime, error)
	Dhuhr(opt option.Option) (model.PeriodicSalatTime, error)
	Asr(opt option.Option) (model.PeriodicSalatTime, error)
	Sunset(opt option.Option) (model.PeriodicSalatTime, error)
	Maghrib(opt option.Option) (model.PeriodicSalatTime, error)
	Isha(opt option.Option) (model.PeriodicSalatTime, error)

	AllTimes(opt option.Option) (model.PeriodicAllSalatTime, error)

	GetOption() option.Option
}
