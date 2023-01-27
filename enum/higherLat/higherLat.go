package higherLatEnum

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/naufalfmm/moslem-salat-times/err"
)

type (
	// HigherLatClass .
	HigherLatClass struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// HigherLat .
	HigherLat int
)

const (
	// NightMiddle .
	NightMiddle HigherLat = iota + 1
	// OneSeventh .
	OneSeventh
	// AngleBased .
	AngleBased
	// None .
	None
)

var (
	higherLatConsts = []HigherLatClass{
		{"nightMiddle", "NightMiddle"},
		{"oneSeventh", "OneSeventh"},
		{"angleBased", "AngleBased"},
		{"none", "None"},
	}
)

// Code .
func (c HigherLat) Code() string {
	if c < 1 || int(c) > len(higherLatConsts) {
		return ""
	}
	return higherLatConsts[c-1].Code
}

// Name .
func (c HigherLat) Name() string {
	if c < 1 || int(c) > len(higherLatConsts) {
		return ""
	}
	return higherLatConsts[c-1].Name
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *HigherLat) UnmarshalParam(src string) error {
	index := findIndex(src, func(c HigherLatClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = HigherLat(index)
	return nil
}

// MarshalJSON presents value to the client
func (c HigherLat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *HigherLat) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal, func(c HigherLatClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = HigherLat(index)
	return nil
}

// Scan retrieves value from the DB
func (c *HigherLat) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findIndex(dbVal, func(c HigherLatClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = HigherLat(index)
	return nil
}

// Value encodes value to the DB
func (c HigherLat) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func findIndex(code string, selector func(c HigherLatClass) string) int {
	for i, v := range higherLatConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteConstants() []HigherLatClass {
	list := make([]HigherLatClass, len(higherLatConsts))
	copy(list, higherLatConsts)
	return list
}

func GetAll() []HigherLat {
	return []HigherLat{
		NightMiddle,
		OneSeventh,
		AngleBased,
	}
}
