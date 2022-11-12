package option

import (
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	sunZenithEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type CalcOpt struct {
	Date time.Time

	FajrZenith     angle.Angle
	IshaZenith     angle.Angle
	IshaZenithType sunZenithEnum.IshaZenithType
}

func (c CalcOpt) SetDate(date time.Time) CalcOpt {
	c.Date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	return c
}

func (c CalcOpt) Now() CalcOpt {
	return c.SetDate(time.Now())
}
