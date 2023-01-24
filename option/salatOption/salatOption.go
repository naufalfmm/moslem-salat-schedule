package salatOption

import (
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/moslem-salat-schedule/consts"
	higherLatEnum "github.com/naufalfmm/moslem-salat-schedule/enum/higherLat"
	mazhabEnum "github.com/naufalfmm/moslem-salat-schedule/enum/mazhab"
	roundingTimeOptionEnum "github.com/naufalfmm/moslem-salat-schedule/enum/roundingTimeOption"
	sunZenithEnum "github.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type SalatOption struct {
	Date time.Time

	Latitude       angle.Angle
	Longitude      angle.Angle
	Elevation      float64
	TimezoneOffset float64

	FajrZenith           angle.Angle
	IshaZenith           angle.Angle
	IshaZenithType       sunZenithEnum.IshaZenithType
	AsrMazhab            mazhabEnum.Mazhab
	HigherLatitudeMethod higherLatEnum.HigherLat

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

type withTimezoneOffset struct {
	timezoneOffset float64
}

func (w withTimezoneOffset) Apply(o *SalatOption) {
	o.TimezoneOffset = float64(w.timezoneOffset)
}

func WithTimezoneOffset(timezoneOffset float64) ApplyingSalatOption {
	return withTimezoneOffset{
		timezoneOffset: timezoneOffset,
	}
}

type withTimezone struct {
	timezone *time.Location
}

func (w withTimezone) Apply(o *SalatOption) {
	now := time.Now().In(w.timezone)
	_, offset := now.Zone()

	WithTimezoneOffset(float64(offset) / consts.OffsetTimezone).Apply(o)
}

func WithTimezone(timezone *time.Location) ApplyingSalatOption {
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

type withHigherLatitudeMethod struct {
	higherLatMethod higherLatEnum.HigherLat
}

func (w withHigherLatitudeMethod) Apply(o *SalatOption) {
	o.HigherLatitudeMethod = w.higherLatMethod
}

func WithHigherLatitudeMethod(higherLatMethod higherLatEnum.HigherLat) ApplyingSalatOption {
	return withHigherLatitudeMethod{
		higherLatMethod: higherLatMethod,
	}
}
