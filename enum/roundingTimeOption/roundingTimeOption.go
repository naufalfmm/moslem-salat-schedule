package roundingTimeOptionEnum

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/err"
)

type (
	// RoundingTimeOptionClass .
	RoundingTimeOptionClass struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// RoundingTimeOption .
	RoundingTimeOption int
)

const (
	// NoRounding .
	NoRounding RoundingTimeOption = iota + 1
	// MinuteCeil .
	MinuteCeil
	// MinuteRound .
	MinuteRound
	// MinuteFloor .
	MinuteFloor
	// HourCeil .
	HourCeil
	// HourRound .
	HourRound
	// HourFloor .
	HourFloor

	// Default .
	Default = NoRounding
)

var (
	roundingTimeOptionConsts = []RoundingTimeOptionClass{
		{"noRounding", "No Rounding"},
		{"minuteCeil", "Minute Ceil"},
		{"minuteRound", "Minute Round"},
		{"minuteFloor", "Minute Floor"},
		{"hourCeil", "Hour Ceil"},
		{"hourRound", "Hour Round"},
		{"hourFloor", "Hour Floor"},
	}
)

// Code .
func (c RoundingTimeOption) Code() string {
	if c < 1 || int(c) > len(roundingTimeOptionConsts) {
		return ""
	}
	return roundingTimeOptionConsts[c-1].Code
}

// Name .
func (c RoundingTimeOption) Name() string {
	if c < 1 || int(c) > len(roundingTimeOptionConsts) {
		return ""
	}
	return roundingTimeOptionConsts[c-1].Name
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *RoundingTimeOption) UnmarshalParam(src string) error {
	index := findRoundingTimeOptionIndex(src, func(c RoundingTimeOptionClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = RoundingTimeOption(index)
	return nil
}

// MarshalJSON presents value to the client
func (c RoundingTimeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *RoundingTimeOption) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findRoundingTimeOptionIndex(rawVal, func(c RoundingTimeOptionClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = RoundingTimeOption(index)
	return nil
}

// Scan retrieves value from the DB
func (c *RoundingTimeOption) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findRoundingTimeOptionIndex(dbVal, func(c RoundingTimeOptionClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = RoundingTimeOption(index)
	return nil
}

// Value encodes value to the DB
func (c RoundingTimeOption) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func (c RoundingTimeOption) roundTimeMinute(t time.Time) time.Time {
	if c == MinuteCeil {
		if t.Second() > 0 {
			return t.Add(1 * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
		}

		return t.Add(-time.Duration(t.Second()) * time.Second)
	}

	if c == MinuteRound {
		if t.Second() >= 30 {
			return t.Add(1 * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
		}

		return t.Add(-time.Duration(t.Second()) * time.Second)
	}

	return t.Add(-time.Duration(t.Second()) * time.Second)
}

func (c RoundingTimeOption) roundTimeHour(t time.Time) time.Time {
	if c == HourCeil {
		if t.Minute() > 0 {
			return t.Add(1 * time.Hour).Add(-time.Duration(t.Minute()) * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
		}

		return t.Add(-time.Duration(t.Minute()) * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
	}

	if c == HourRound {
		if t.Minute() >= 30 {
			return t.Add(1 * time.Hour).Add(-time.Duration(t.Minute()) * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
		}

		return t.Add(-time.Duration(t.Minute()) * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
	}

	return t.Add(1 * time.Hour).Add(-time.Duration(t.Minute()) * time.Minute).Add(-time.Duration(t.Second()) * time.Second)
}

func (c RoundingTimeOption) RoundTime(t time.Time) time.Time {
	if c >= MinuteCeil && c <= MinuteFloor {
		return c.roundTimeMinute(t)
	}

	if c >= HourCeil && c <= HourFloor {
		return c.roundTimeHour(t)
	}

	return t
}

func findRoundingTimeOptionIndex(code string, selector func(c RoundingTimeOptionClass) string) int {
	for i, v := range roundingTimeOptionConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}
