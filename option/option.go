package option

import (
	"errors"

	"github.com/shopspring/decimal"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	sunZenithEnum "gitlab.com/naufalfmm/moslem-salat-schedule/enum/sunZenith"
)

type Option struct {
	Title string

	LocOpt  LocOpt
	CalcOpt CalcOpt
}

func (opt Option) Validate() error {
	if opt.LocOpt.Latitude.AngleType() != opt.LocOpt.Longitude.AngleType() {
		return errors.New("latitude and longitude should have same degree type")
	}

	if opt.LocOpt.Elevation.IsZero() {
		return errors.New("elevation should be exist")
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
	o.LocOpt.Latitude = w.latitude
	o.LocOpt.Longitude = w.longitude
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
	o.LocOpt.Timezone = decimal.NewFromInt(w.timezone)
}

func WithTimezone(timezone int64) ApplyingOption {
	return withTimezone{
		timezone: timezone,
	}
}

type withFajrIshaZenith struct {
	fajrZenith angle.Angle
	ishaZenith angle.Angle
}

func (w withFajrIshaZenith) Apply(o *Option) {
	o.CalcOpt.FajrZenith = w.fajrZenith
	o.CalcOpt.IshaZenith = w.ishaZenith
	o.CalcOpt.IshaZenithType = sunZenithEnum.Standard
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
	o.CalcOpt.FajrZenith = w.sunZenith.FajrZenith()
	o.CalcOpt.IshaZenith = w.sunZenith.IshaZenith().Angle
	o.CalcOpt.IshaZenithType = w.sunZenith.IshaZenith().Type
}

func WithSunZenith(sunZenith sunZenithEnum.SunZenith) ApplyingOption {
	return withSunZenith{
		sunZenith: sunZenith,
	}
}
