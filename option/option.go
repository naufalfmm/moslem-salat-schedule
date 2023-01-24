package option

import (
	"errors"
	"time"

	"github.com/naufalfmm/angle"
	higherLatEnum "github.com/naufalfmm/moslem-salat-schedule/enum/higherLat"
	mazhabEnum "github.com/naufalfmm/moslem-salat-schedule/enum/mazhab"
	roundingTimeOptionEnum "github.com/naufalfmm/moslem-salat-schedule/enum/roundingTimeOption"
	sunZenithEnum "github.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type Option struct {
	Title string

	Date time.Time

	Latitude  angle.Angle
	Longitude angle.Angle
	Elevation float64
	Timezone  float64

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

func (opt Option) Validate() error {
	if opt.Latitude.AngleType() != opt.Longitude.AngleType() {
		return errors.New("latitude and longitude should have same degree type")
	}

	if opt.Elevation == 0 {
		return errors.New("elevation should be exist")
	}

	if opt.HigherLatitudeMethod == 0 {
		opt.HigherLatitudeMethod = higherLatEnum.None
	}

	return nil
}

type withTitle string

func (w withTitle) Apply(o *Option) {
	o.Title = string(w)
}

func WithTitle(title string) ApplyingOption {
	return withTitle(title)
}

type withLatitudeLongitude struct {
	latitude  angle.Angle
	longitude angle.Angle
}

func (w withLatitudeLongitude) Apply(o *Option) {
	o.Latitude = w.latitude
	o.Longitude = w.longitude
}

func WithLatitudeLongitude(lat, long angle.Angle) ApplyingOption {
	return withLatitudeLongitude{
		latitude:  lat,
		longitude: long,
	}
}

type withTimezone struct {
	timezone int64
}

func (w withTimezone) Apply(o *Option) {
	o.Timezone = float64(w.timezone)
}

func WithTimezone(timezone int64) ApplyingOption {
	return withTimezone{
		timezone: timezone,
	}
}

type withElevation struct {
	elevation float64
}

func (w withElevation) Apply(o *Option) {
	o.Elevation = w.elevation
}

func WithElevation(elevation float64) ApplyingOption {
	return withElevation{
		elevation: elevation,
	}
}

type withFajrIshaZenith struct {
	fajrZenith angle.Angle
	ishaZenith angle.Angle
}

func (w withFajrIshaZenith) Apply(o *Option) {
	o.FajrZenith = w.fajrZenith
	o.IshaZenith = w.ishaZenith
	o.IshaZenithType = sunZenithEnum.Standard
}

func WithFajrIshaZenith(fajrZenith, ishaZenith angle.Angle) ApplyingOption {
	return withFajrIshaZenith{
		fajrZenith: fajrZenith,
		ishaZenith: ishaZenith,
	}
}

type withSunZenith struct {
	sunZenith sunZenithEnum.SunZenith
}

func (w withSunZenith) Apply(o *Option) {
	o.FajrZenith = w.sunZenith.FajrZenith()
	o.IshaZenith = w.sunZenith.IshaZenith().Angle
	o.IshaZenithType = w.sunZenith.IshaZenith().Type
}

func WithSunZenith(sunZenith sunZenithEnum.SunZenith) ApplyingOption {
	return withSunZenith{
		sunZenith: sunZenith,
	}
}

type withMazhab struct {
	mazhab mazhabEnum.Mazhab
}

func (w withMazhab) Apply(o *Option) {
	o.AsrMazhab = w.mazhab
}

func WithMazhab(mazhab mazhabEnum.Mazhab) ApplyingOption {
	return withMazhab{
		mazhab: mazhab,
	}
}

type withRoundingTimeOption struct {
	roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption
}

func (w withRoundingTimeOption) Apply(o *Option) {
	o.RoundingTimeOption = w.roundingTimeOpt
}

func WithRoundingTimeOption(roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption) ApplyingOption {
	return withRoundingTimeOption{
		roundingTimeOpt: roundingTimeOpt,
	}
}

type withHigherLatitudeMethod struct {
	higherLatMethod higherLatEnum.HigherLat
}

func (w withHigherLatitudeMethod) Apply(o *Option) {
	o.HigherLatitudeMethod = w.higherLatMethod
}

func WithHigherLatitudeMethod(higherLatMethod higherLatEnum.HigherLat) ApplyingOption {
	return withHigherLatitudeMethod{
		higherLatMethod: higherLatMethod,
	}
}
