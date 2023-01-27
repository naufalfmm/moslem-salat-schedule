package model

import (
	"time"

	salatEnum "github.com/naufalfmm/moslem-salat-times/enum/salat"
)

type (
	SalatTime struct {
		Date  time.Time       `json:"date"`
		Salat salatEnum.Salat `json:"salat"`
		Time  time.Time       `json:"time"`
	}

	PeriodicSalatTime []SalatTime

	AllSalatTime struct {
		Date       time.Time         `json:"date"`
		SalatTimes PeriodicSalatTime `json:"salat_times"`
	}

	PeriodicAllSalatTime []AllSalatTime
)
