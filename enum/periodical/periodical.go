package periodicalEnum

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/naufalfmm/moslem-salat-schedule/err"
)

type (
	// PeriodicalClass .
	PeriodicalClass struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// Periodical .
	Periodical int
)

const (
	// Daily .
	Daily Periodical = iota + 1
	// Weekly .
	Weekly
	// Monthly .
	Monthly
	// CurrentMonthly .
	CurrentMonthly
	// Custom .
	Custom

	weeklyDaysRange = 6.
)

var (
	periodicalConsts = []PeriodicalClass{
		{"daily", "Daily"},
		{"weekly", "Weekly"},
		{"monthly", "Monthly"},
		{"currentMonthly", "Current Monthly"},
		{"custom", "Custom"},
	}
)

// Code .
func (c Periodical) Code() string {
	if c < 1 || int(c) > len(periodicalConsts) {
		return ""
	}
	return periodicalConsts[c-1].Code
}

// Name .
func (c Periodical) Name() string {
	if c < 1 || int(c) > len(periodicalConsts) {
		return ""
	}
	return periodicalConsts[c-1].Name
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *Periodical) UnmarshalParam(src string) error {
	index := findIndex(src, func(c PeriodicalClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Periodical(index)
	return nil
}

// MarshalJSON presents value to the client
func (c Periodical) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *Periodical) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal, func(c PeriodicalClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Periodical(index)
	return nil
}

// Scan retrieves value from the DB
func (c *Periodical) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findIndex(dbVal, func(c PeriodicalClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Periodical(index)
	return nil
}

// Value encodes value to the DB
func (c Periodical) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func (c Periodical) GetDateRange(date time.Time) (time.Time, time.Time) {
	if c == Weekly {
		return date, date.AddDate(0, 0, weeklyDaysRange)
	}

	if c == Monthly {
		return date, date.AddDate(0, 1, 0).Add(-24 * time.Hour)
	}

	if c == CurrentMonthly {
		return date.AddDate(0, 0, -date.Day()+1), date.AddDate(0, 1, -date.Day())
	}

	return date, date
}

func findIndex(code string, selector func(c PeriodicalClass) string) int {
	for i, v := range periodicalConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteConstants() []PeriodicalClass {
	list := make([]PeriodicalClass, len(periodicalConsts))
	copy(list, periodicalConsts)
	return list
}

func GetByDateRange(dateStart, dateEnd time.Time) Periodical {
	if int(dateEnd.Sub(dateStart).Hours()/24.) == weeklyDaysRange {
		return Weekly
	}

	if dateStart.AddDate(0, 1, 0).Add(-24 * time.Hour).Equal(dateEnd) {
		return Monthly
	}

	if dateEnd.Equal(dateStart) {
		return Daily
	}

	if dateStart.Equal(dateStart.AddDate(0, 0, -dateStart.Day()+1)) &&
		dateEnd.Equal(dateStart.AddDate(0, 1, -dateStart.Day())) {
		return CurrentMonthly
	}

	return Custom
}
