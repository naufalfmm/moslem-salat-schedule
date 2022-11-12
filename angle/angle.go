package angle

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
	"gitlab.com/naufalfmm/moslem-salat-schedule/err"
)

type Angle struct {
	degree decimal.Decimal
	minute decimal.Decimal
	second decimal.Decimal

	neg     bool
	angType angleType.AngleType
}

func (d *Angle) fillBySymbol(src string, symbol rune) error {
	if symbol == consts.DegreeSymbolRune {
		if err := d.degree.Scan(src); err != nil {
			return err
		}
	}

	if symbol == consts.MinuteSymbolRune {
		if err := d.minute.Scan(src); err != nil {
			return err
		}
	}

	if err := d.second.Scan(src); err != nil {
		return err
	}

	return nil
}

func (d *Angle) decideFillBySymbol(src string, symbol rune) error {
	d.angType = angleType.Decimal
	if symbol == consts.MinuteSymbolRune ||
		symbol == consts.SecondSymbolRune {
		d.angType = angleType.DegreeMinuteSecond
	}

	return d.fillBySymbol(src, symbol)
}

func (d *Angle) scanByString(src string) error {
	var (
		buff bytes.Buffer
		s    rune
	)

	for _, s = range src {
		if s != consts.DegreeSymbolRune &&
			s != consts.MinuteSymbolRune &&
			s != consts.SecondSymbolRune {
			if _, err := buff.WriteRune(s); err != nil {
				return err
			}

			continue
		}

		if s == consts.NegativeSymbolRune {
			d.neg = true
			continue
		}

		if err := d.decideFillBySymbol(buff.String(), s); err != nil {
			return err
		}
	}

	return d.decideFillBySymbol(buff.String(), s)
}

func (d *Angle) UnmarshalParam(src string) error {
	return d.scanByString(src)
}

func (d *Angle) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	return d.scanByString(rawVal)
}

func (d *Angle) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	return d.scanByString(dbVal)
}

func (d Angle) String() string {
	var neg string
	if d.neg {
		neg = string(consts.NegativeSymbolRune)
	}

	if d.angType == angleType.Decimal {
		return fmt.Sprintf("%s%s", neg, d.degree.String()+string(consts.DegreeSymbolRune))
	}

	return fmt.Sprintf("%s%s", neg, d.degree.String()+string(consts.DegreeSymbolRune)+
		d.minute.String()+string(consts.MinuteSymbolRune)+
		d.second.String()+string(consts.SecondSymbolRune))
}

func (d Angle) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d Angle) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d Angle) AngleType() angleType.AngleType {
	return d.angType
}
