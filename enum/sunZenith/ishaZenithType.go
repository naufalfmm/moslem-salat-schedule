package sunZenithEnum

import (
	"database/sql/driver"
	"encoding/json"

	"gitlab.com/naufalfmm/moslem-salat-schedule/err"
)

type (
	// IshaZenithTypeClass .
	IshaZenithTypeClass struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// IshaZenithType .
	IshaZenithType int
)

const (
	// Standard .
	Standard IshaZenithType = iota + 1
	// AfterMagrib .
	AfterMagrib
)

var (
	ishaZenithTypeConsts = []IshaZenithTypeClass{
		{"standard", "Standard"},
		{"afterMagrib", "After Magrib"},
	}
)

// Code .
func (c IshaZenithType) Code() string {
	if c < 1 || int(c) > len(ishaZenithTypeConsts) {
		return ""
	}
	return ishaZenithTypeConsts[c-1].Code
}

// Name .
func (c IshaZenithType) Name() string {
	if c < 1 || int(c) > len(ishaZenithTypeConsts) {
		return ""
	}
	return ishaZenithTypeConsts[c-1].Name
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *IshaZenithType) UnmarshalParam(src string) error {
	index := findIshaZenithTypeIndex(src, func(c IshaZenithTypeClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = IshaZenithType(index)
	return nil
}

// MarshalJSON presents value to the client
func (c IshaZenithType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *IshaZenithType) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIshaZenithTypeIndex(rawVal, func(c IshaZenithTypeClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = IshaZenithType(index)
	return nil
}

// Scan retrieves value from the DB
func (c *IshaZenithType) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findIshaZenithTypeIndex(dbVal, func(c IshaZenithTypeClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = IshaZenithType(index)
	return nil
}

// Value encodes value to the DB
func (c IshaZenithType) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func findIshaZenithTypeIndex(code string, selector func(c IshaZenithTypeClass) string) int {
	for i, v := range ishaZenithTypeConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteIshaZenithTypeConstants() []IshaZenithTypeClass {
	list := make([]IshaZenithTypeClass, len(ishaZenithTypeConsts))
	copy(list, ishaZenithTypeConsts)
	return list
}
