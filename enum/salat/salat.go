package salatEnum

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/naufalfmm/moslem-salat-times/err"
)

type (
	// SalatClass .
	SalatClass struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// Salat .
	Salat int
)

const (
	// Fajr .
	Fajr Salat = iota + 1
	// Sunrise .
	Sunrise
	// Dhuhr .
	Dhuhr
	// Asr .
	Asr
	// Sunset .
	Sunset
	// Maghrib .
	Maghrib
	// Isha .
	Isha
	// Midnight .
	Midnight
)

var (
	salatConsts = []SalatClass{
		{"fajr", "Fajr"},
		{"sunrise", "Sunrise"},
		{"dhuhr", "Dhuhr"},
		{"asr", "Asr"},
		{"sunset", "Sunset"},
		{"maghrib", "Maghrib"},
		{"isha", "Isha"},
		{"midnight", "Midnight"},
	}
)

// Code .
func (c Salat) Code() string {
	if c < 1 || int(c) > len(salatConsts) {
		return ""
	}
	return salatConsts[c-1].Code
}

// Name .
func (c Salat) Name() string {
	if c < 1 || int(c) > len(salatConsts) {
		return ""
	}
	return salatConsts[c-1].Name
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *Salat) UnmarshalParam(src string) error {
	index := findIndex(src, func(c SalatClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Salat(index)
	return nil
}

// MarshalJSON presents value to the client
func (c Salat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *Salat) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal, func(c SalatClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Salat(index)
	return nil
}

// Scan retrieves value from the DB
func (c *Salat) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findIndex(dbVal, func(c SalatClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Salat(index)
	return nil
}

// Value encodes value to the DB
func (c Salat) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func findIndex(code string, selector func(c SalatClass) string) int {
	for i, v := range salatConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteConstants() []SalatClass {
	list := make([]SalatClass, len(salatConsts))
	copy(list, salatConsts)
	return list
}

func GetAll() []Salat {
	return []Salat{
		Fajr,
		Dhuhr,
		Asr,
		Maghrib,
		Isha,
	}
}
