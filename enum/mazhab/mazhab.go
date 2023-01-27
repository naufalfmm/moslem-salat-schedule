package mazhabEnum

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/naufalfmm/moslem-salat-times/err"
)

type (
	// MazhabClass .
	MazhabClass struct {
		Code            string  `json:"code"`
		Name            string  `json:"name"`
		AsrShadowLength float64 `json:"asrShadowLength"`
	}

	// Mazhab .
	Mazhab int
)

const (
	// Standard .
	Standard Mazhab = iota + 1
	// Hanafi .
	Hanafi
)

var (
	mazhabConsts = []MazhabClass{
		{"standard", "Standard", 1},
		{"hanafi", "Hanafi", 2},
	}
)

// Code .
func (c Mazhab) Code() string {
	if c < 1 || int(c) > len(mazhabConsts) {
		return ""
	}
	return mazhabConsts[c-1].Code
}

// Name .
func (c Mazhab) Name() string {
	if c < 1 || int(c) > len(mazhabConsts) {
		return ""
	}
	return mazhabConsts[c-1].Name
}

// AsrShadowLength .
func (c Mazhab) AsrShadowLength() float64 {
	if c < 1 || int(c) > len(mazhabConsts) {
		return 0
	}
	return mazhabConsts[c-1].AsrShadowLength
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *Mazhab) UnmarshalParam(src string) error {
	index := findIndex(src, func(c MazhabClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Mazhab(index)
	return nil
}

// MarshalJSON presents value to the client
func (c Mazhab) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *Mazhab) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal, func(c MazhabClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Mazhab(index)
	return nil
}

// Scan retrieves value from the DB
func (c *Mazhab) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findIndex(dbVal, func(c MazhabClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = Mazhab(index)
	return nil
}

// Value encodes value to the DB
func (c Mazhab) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func findIndex(code string, selector func(c MazhabClass) string) int {
	for i, v := range mazhabConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteConstants() []MazhabClass {
	list := make([]MazhabClass, len(mazhabConsts))
	copy(list, mazhabConsts)
	return list
}
