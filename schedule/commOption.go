package schedule

import (
	"fmt"
	"time"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/moslem-salat-times/consts"
	higherLatEnum "github.com/naufalfmm/moslem-salat-times/enum/higherLat"
	mazhabEnum "github.com/naufalfmm/moslem-salat-times/enum/mazhab"
	periodicalEnum "github.com/naufalfmm/moslem-salat-times/enum/periodical"
	roundingTimeOptionEnum "github.com/naufalfmm/moslem-salat-times/enum/roundingTimeOption"
	sunZenithEnum "github.com/naufalfmm/moslem-salat-times/enum/sunZenith"
	"github.com/naufalfmm/moslem-salat-times/utils/sunPositions"
)

type CommOpt struct {
	dateStart  time.Time
	dateEnd    time.Time
	periodical periodicalEnum.Periodical

	latitude    angle.Angle
	longitude   angle.Angle
	elevation   float64
	timezoneLoc *time.Location

	fajrZenith     angle.Angle
	ishaZenith     angle.Angle
	ishaZenithType sunZenithEnum.IshaZenithType

	mazhab               mazhabEnum.Mazhab
	higherLatitudeMethod higherLatEnum.HigherLat

	roundingTimeOption roundingTimeOptionEnum.RoundingTimeOption

	sunPositions sunPositions.SunPositions
}

func (c CommOpt) ToOption() Option {
	return Option(c)
}

func (c *CommOpt) CalculateSunPositions() (CommOpt, error) {
	if len(c.sunPositions) > 0 {
		return *c, nil
	}

	if c.dateStart.IsZero() {
		SetNow().Apply(c)
	}

	if c.dateEnd.IsZero() {
		c.dateEnd = c.dateStart
		c.periodical = periodicalEnum.Custom
	}

	if c.timezoneLoc == nil {
		c.timezoneLoc = c.dateStart.Location()
	}

	c.sunPositions = sunPositions.NewFromDateRange(c.dateStart, c.dateEnd, c.timezoneLoc, c.longitude)
	return *c, nil
}

type ApplyCommOpt interface {
	Apply(o *CommOpt)
}

type setNow struct{}

func (s setNow) Apply(o *CommOpt) {
	o.dateStart = time.Now()
	o.dateEnd = o.dateStart
	o.periodical = periodicalEnum.Custom
}

func SetNow() ApplyCommOpt {
	return setNow{}
}

type withDateRange struct {
	dateStart, dateEnd time.Time
}

func (w withDateRange) Apply(o *CommOpt) {
	o.dateStart = w.dateStart
	o.dateEnd = w.dateEnd
	o.periodical = periodicalEnum.Custom
}

type withPeriodical struct {
	periodical periodicalEnum.Periodical
}

func (w withPeriodical) Apply(o *CommOpt) {
	date := o.dateStart
	if date.IsZero() {
		date = time.Now()

		if o.timezoneLoc != nil {
			date = date.In(o.timezoneLoc)
		}
	}

	o.dateStart, o.dateEnd = w.periodical.GetDateRange(date)
	o.periodical = w.periodical
}

func WithPeriodical(periodical periodicalEnum.Periodical) ApplyCommOpt {
	return withPeriodical{
		periodical: periodical,
	}
}

type withLatitudeLongitude struct {
	latitude  angle.Angle
	longitude angle.Angle
}

func (w withLatitudeLongitude) Apply(o *CommOpt) {
	o.latitude = w.latitude
	o.longitude = w.longitude
}

func WithLatitudeLongitude(lat, long angle.Angle) ApplyCommOpt {
	return withLatitudeLongitude{
		latitude:  lat,
		longitude: long,
	}
}

type withTimezoneOffset struct {
	timezoneOffset float64
}

func (w withTimezoneOffset) Apply(o *CommOpt) {
	angTime := angle.NewDegreeFromFloat(w.timezoneOffset)

	negStr := ""
	if angTime.IsNegative() {
		negStr = "-"
	}

	o.timezoneLoc = time.FixedZone(fmt.Sprintf("%s%s", negStr, angTime.Abs().ToTime().Format("0304")), int(w.timezoneOffset*consts.OffsetTimezone))
}

func WithTimezoneOffset(timezoneOffset float64) ApplyCommOpt {
	return withTimezoneOffset{
		timezoneOffset: timezoneOffset,
	}
}

type withTimezone struct {
	timezone *time.Location
}

func (w withTimezone) Apply(o *CommOpt) {
	o.timezoneLoc = w.timezone
}

func WithTimezone(timezone *time.Location) ApplyCommOpt {
	return withTimezone{
		timezone: timezone,
	}
}

type withElevation struct {
	elevation float64
}

func (w withElevation) Apply(o *CommOpt) {
	o.elevation = w.elevation
}

func WithElevation(elevation float64) ApplyCommOpt {
	return withElevation{
		elevation: elevation,
	}
}

type withFajrIshaZenith struct {
	fajrZenith angle.Angle
	ishaZenith angle.Angle
}

func (w withFajrIshaZenith) Apply(o *CommOpt) {
	o.fajrZenith = w.fajrZenith
	o.ishaZenith = w.ishaZenith
	o.ishaZenithType = sunZenithEnum.Standard
}

func WithFajrIshaZenith(fajrZenith, ishaZenith angle.Angle) ApplyCommOpt {
	return withFajrIshaZenith{
		fajrZenith: fajrZenith,
		ishaZenith: ishaZenith,
	}
}

type withSunZenith struct {
	sunZenith sunZenithEnum.SunZenith
}

func (w withSunZenith) Apply(o *CommOpt) {
	o.fajrZenith = w.sunZenith.FajrZenith()
	o.ishaZenith = w.sunZenith.IshaZenith().Angle
	o.ishaZenithType = w.sunZenith.IshaZenith().Type
}

func WithSunZenith(sunZenith sunZenithEnum.SunZenith) ApplyCommOpt {
	return withSunZenith{
		sunZenith: sunZenith,
	}
}

type withMazhab struct {
	mazhab mazhabEnum.Mazhab
}

func (w withMazhab) Apply(o *CommOpt) {
	o.mazhab = w.mazhab
}

func WithMazhab(mazhab mazhabEnum.Mazhab) ApplyCommOpt {
	return withMazhab{
		mazhab: mazhab,
	}
}

type withRoundingTimeOption struct {
	roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption
}

func (w withRoundingTimeOption) Apply(o *CommOpt) {
	o.roundingTimeOption = w.roundingTimeOpt
}

func WithRoundingTimeOption(roundingTimeOpt roundingTimeOptionEnum.RoundingTimeOption) ApplyCommOpt {
	return withRoundingTimeOption{
		roundingTimeOpt: roundingTimeOpt,
	}
}

type withHigherLatitudeMethod struct {
	higherLatMethod higherLatEnum.HigherLat
}

func (w withHigherLatitudeMethod) Apply(o *CommOpt) {
	o.higherLatitudeMethod = w.higherLatMethod
}

func WithHigherLatitudeMethod(higherLatMethod higherLatEnum.HigherLat) ApplyCommOpt {
	return withHigherLatitudeMethod{
		higherLatMethod: higherLatMethod,
	}
}
