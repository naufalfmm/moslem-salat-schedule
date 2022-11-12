package moslemSalatSchedule

import (
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/model"
	"gitlab.com/naufalfmm/moslem-salat-schedule/option/salatOption"
)

type MoslemSalatSchedule interface {
	SetDate(date time.Time) MoslemSalatSchedule
	Now() MoslemSalatSchedule

	Fajr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Dhuhr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Asr(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Maghrib(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)
	Isha(opts ...salatOption.ApplyingSalatOption) (model.SalatTime, error)

	AllFiveTimes(opts ...salatOption.ApplyingSalatOption) (model.FiveSalatTime, error)
}
