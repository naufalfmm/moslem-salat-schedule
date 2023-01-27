package moslemSalatTimes

import (
	"github.com/naufalfmm/moslem-salat-times/model"
	"github.com/naufalfmm/moslem-salat-times/option"
)

type MoslemSalatTimes interface {
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
