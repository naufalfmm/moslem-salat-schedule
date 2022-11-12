package err

import "errors"

var (
	ErrUnknownConstant   = errors.New("unknown constant")
	ErrConstantParsing   = errors.New("expected string for the constant")
	ErrDateMissing       = errors.New("date missing")
	ErrFajrZenithMissing = errors.New("fajr zenith angle missing")
	ErrIshaZenithMissing = errors.New("isha zenith angle missing")
	ErrTimezoneMissing   = errors.New("timezone missing")
	ErrLatitudeMissing   = errors.New("latitude missing")
	ErrLongitudeMissing  = errors.New("longitude missing")
)
