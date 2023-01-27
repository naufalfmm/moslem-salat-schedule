package option

import (
	"time"

	"github.com/naufalfmm/angle"
	higherLatEnum "github.com/naufalfmm/moslem-salat-times/enum/higherLat"
	mazhabEnum "github.com/naufalfmm/moslem-salat-times/enum/mazhab"
	periodicalEnum "github.com/naufalfmm/moslem-salat-times/enum/periodical"
	roundingTimeOptionEnum "github.com/naufalfmm/moslem-salat-times/enum/roundingTimeOption"
	salatEnum "github.com/naufalfmm/moslem-salat-times/enum/salat"
	sunZenithEnum "github.com/naufalfmm/moslem-salat-times/enum/sunZenith"
	"github.com/naufalfmm/moslem-salat-times/utils/sunPositions"
)

type Option interface {
	SetDateRange(dateStart, dateEnd time.Time) Option
	SetNow() Option
	SetDatePeriodical(dateStart time.Time, periodical periodicalEnum.Periodical) Option
	SetPeriodical(periodical periodicalEnum.Periodical) Option
	SetLatitudeLongitude(latitude, longitude angle.Angle) Option
	SetElevation(elevation float64) Option
	SetMazhab(mazhab mazhabEnum.Mazhab) Option
	SetHigherLatitudeMethod(higherLatMethod higherLatEnum.HigherLat) Option
	SetRoundingTimeOption(roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption) Option

	SetTimezoneOffset(timezoneOffset float64) Option
	SetTimezone(timezone *time.Location) Option

	SetFajrIshaZenith(fajrZenith, ishaZenith angle.Angle) Option
	SetSunZenith(sunZenith sunZenithEnum.SunZenith) Option

	ValidateBySalat(salat salatEnum.Salat) error

	CalculateSunPositions() (Option, error)
	CalculateFajrHighAltitude(declination angle.Angle) angle.Angle
	CalculateSunriseSunsetHighAltitude(declination angle.Angle) angle.Angle
	CalculateAsrAngle(declination angle.Angle) angle.Angle
	CalculateIshaHighAltitude(declination angle.Angle) (angle.Angle, sunZenithEnum.IshaZenithType)

	RoundTime(t time.Time) time.Time

	GetSunPositions() sunPositions.SunPositions
	GetDateRange() (time.Time, time.Time)
}
