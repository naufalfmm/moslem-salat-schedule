package model

import (
	"time"

	salatEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/salat"
)

type (
	SalatTime struct {
		Date  time.Time       `json:"date"`
		Salat salatEnum.Salat `json:"salat"`
		Time  time.Time       `json:"time"`
	}

	SalatTimes []SalatTime

	FiveSalatTime struct {
		Date       time.Time  `json:"date"`
		SalatTimes SalatTimes `json:"salat_times"`
	}
)
