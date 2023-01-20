package sunZenithEnum

import (
	"database/sql/driver"
	"encoding/json"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
	"gitlab.com/naufalfmm/moslem-salat-schedule/err"
)

type (
	IshaZenith struct {
		Angle angle.Angle
		Type  IshaZenithType
	}

	// SunZenithClass .
	SunZenithClass struct {
		Code string      `json:"code"`
		Name string      `json:"name"`
		Fajr angle.Angle `json:"fajr"`
		Isha IshaZenith  `json:"isha"`
	}

	// SunZenith .
	SunZenith int
)

const (
	// KEMENAG .
	KEMENAG SunZenith = iota + 1
	// ESA .
	ESA
	// ISNA .
	ISNA
	// MCW .
	MCW
	// MWL .
	MWL
	// UAU .
	UAU
	// UIS .
	UIS
	// JAKIM .
	JAKIM
	// MUIS .
	MUIS
	// DIYANET .
	DIYANET
	// UOIF .
	UOIF
)

var (
	sunZenithConsts = []SunZenithClass{
		{"KEMENAG", "Kementerian Agama Republik Indonesia", angle.NewDegreeFromFloat(20), IshaZenith{angle.NewDegreeFromFloat(18), Standard}},
		{"ESA", "Egyptian General Authority Survey", angle.NewDegreeFromFloat(19.5), IshaZenith{angle.NewDegreeFromFloat(17.5), Standard}},
		{"ISNA", "Islamic Society of North America", angle.NewDegreeFromFloat(15), IshaZenith{angle.NewDegreeFromFloat(15), Standard}},
		{"MCW", "Moonsighting Committee Worldwide", angle.NewDegreeFromFloat(18), IshaZenith{angle.NewDegreeFromFloat(18), Standard}},
		{"MWL", "Muslim World League", angle.NewDegreeFromFloat(18), IshaZenith{angle.NewDegreeFromFloat(17), Standard}},
		{"UAU", "Umm Al-Qura University", angle.NewDegreeFromFloat(18.5), IshaZenith{angle.NewFromDegreeMinuteSecond(1., 30., consts.DecimalZero), AfterMagrib}},
		{"UIS", "University of Islamic Sciences, Karachi", angle.NewDegreeFromFloat(18), IshaZenith{angle.NewDegreeFromFloat(18), Standard}},
		{"JAKIM", "Jabatan Kemajuan Islam Malaysia", angle.NewDegreeFromFloat(18), IshaZenith{angle.NewDegreeFromFloat(18), Standard}},
		{"MUIS", "Majlis Ugama Islam Singapura", angle.NewDegreeFromFloat(20), IshaZenith{angle.NewDegreeFromFloat(18), Standard}},
		{"DIYANET", "Directorate of Religious Affairs", angle.NewDegreeFromFloat(18), IshaZenith{angle.NewDegreeFromFloat(17), Standard}},
		{"UOIF", "Union of Islamic Organisations of France", angle.NewDegreeFromFloat(12), IshaZenith{angle.NewDegreeFromFloat(12), Standard}},
	}
)

// Code .
func (c SunZenith) Code() string {
	if c < 1 || int(c) > len(sunZenithConsts) {
		return ""
	}
	return sunZenithConsts[c-1].Code
}

// Name .
func (c SunZenith) Name() string {
	if c < 1 || int(c) > len(sunZenithConsts) {
		return ""
	}
	return sunZenithConsts[c-1].Name
}

// FajrZenith .
func (c SunZenith) FajrZenith() angle.Angle {
	if c < 1 || int(c) > len(sunZenithConsts) {
		return angle.Zero
	}
	return sunZenithConsts[c-1].Fajr
}

func (c SunZenith) IshaZenith() IshaZenith {
	if c < 1 || int(c) > len(sunZenithConsts) {
		return IshaZenith{}
	}
	return sunZenithConsts[c-1].Isha
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *SunZenith) UnmarshalParam(src string) error {
	index := findIndex(src, func(c SunZenithClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = SunZenith(index)
	return nil
}

// MarshalJSON presents value to the client
func (c SunZenith) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Code())
}

// UnmarshalJSON parses value from the client
func (c *SunZenith) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := findIndex(rawVal, func(c SunZenithClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = SunZenith(index)
	return nil
}

// Scan retrieves value from the DB
func (c *SunZenith) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := findIndex(dbVal, func(c SunZenithClass) string {
		return c.Code
	})

	if index == 0 {
		return err.ErrUnknownConstant
	}

	*c = SunZenith(index)
	return nil
}

// Value encodes value to the DB
func (c SunZenith) Value() (driver.Value, error) {
	return string(c.Code()), nil
}

func findIndex(code string, selector func(c SunZenithClass) string) int {
	for i, v := range sunZenithConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents constants as their complete object form
func AsCompleteConstants() []SunZenithClass {
	list := make([]SunZenithClass, len(sunZenithConsts))
	copy(list, sunZenithConsts)
	return list
}
