package moslemSalatSchedule

import (
	"time"

	"github.com/naufalfmm/moslem-salat-schedule/model"
	"github.com/naufalfmm/moslem-salat-schedule/option/salatOption"
)

type MoslemSalatSchedule interface {
	SetDate(date time.Time) MoslemSalatSchedule
	Now() MoslemSalatSchedule

	Midnight(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Fajr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Sunrise(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Dhuhr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Asr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Sunset(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Maghrib(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Isha(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)

	AllTimes(opts ...salatOption.ApplyingSalatOption) (model.AllSalatTimes, error)
}
