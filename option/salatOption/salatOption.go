package salatOption

import (
	"time"

	"github.com/naufalfmm/angle"
	mazhabEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/mazhab"
	roundingTimeOptionEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/roundingTimeOption"
	sunZenithEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type SalatOption struct {
	Date time.Time

	Latitude  angle.Angle
	Longitude angle.Angle
	Elevation float64
	Timezone  float64

	FajrZenith     angle.Angle
	IshaZenith     angle.Angle
	IshaZenithType sunZenithEnum.IshaZenithType
	AsrMazhab      mazhabEnum.Mazhab

	RoundingTimeOption roundingTimeOptionEnum.RoundingTimeOption

	julianDay  float64
	julianDate float64

	meanAnomaly    angle.Angle
	meanLongSun    angle.Angle
	eclipticLong   angle.Angle
	obliquity      angle.Angle
	rightAscension angle.Angle
	equationOfTime angle.Angle

	Declination    angle.Angle
	SunTransitTime angle.Angle
}

type withDate struct {
	date time.Time
}

func (w withDate) Apply(o *SalatOption) {
	o.SetDate(w.date)
}

func WithDate(date time.Time) ApplyingSalatOption {
	return withDate{
		date: date,
	}
}

type withSunZenith struct {
	sunZenith sunZenithEnum.SunZenith
}

func (w withSunZenith) Apply(o *SalatOption) {
	o.FajrZenith = w.sunZenith.FajrZenith()
	o.IshaZenith = w.sunZenith.IshaZenith().Angle
	o.IshaZenithType = w.sunZenith.IshaZenith().Type
}

func WithSunZenith(sunZenith sunZenithEnum.SunZenith) ApplyingSalatOption {
	return withSunZenith{
		sunZenith: sunZenith,
	}
}

type withTimezone struct {
	timezone int64
}

func (w withTimezone) Apply(o *SalatOption) {
	o.Timezone = float64(w.timezone)
}

func WithTimezone(timezone int64) ApplyingSalatOption {
	return withTimezone{
		timezone: timezone,
	}
}

type withMazhab struct {
	mazhab mazhabEnum.Mazhab
}

func (w withMazhab) Apply(o *SalatOption) {
	o.AsrMazhab = w.mazhab
}

func WithMazhab(mazhab mazhabEnum.Mazhab) ApplyingSalatOption {
	return withMazhab{
		mazhab: mazhab,
	}
}

type withRoundingTimeOption struct {
	roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption
}

func (w withRoundingTimeOption) Apply(o *SalatOption) {
	o.RoundingTimeOption = w.roundingTimeOpt
}

func WithRoundingTimeOption(roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption) ApplyingSalatOption {
	return withRoundingTimeOption{
		roundingTimeOpt: roundingTimeOpt,
	}
}
