package angle

import (
	"github.com/shopspring/decimal"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
)

var Zero = NewFromDecimal(decimal.Zero)

func NewFromDecimal(dec decimal.Decimal) Angle {
	return Angle{
		degree:  dec.Abs(),
		neg:     dec.IsNegative(),
		angType: angleType.Decimal,
	}
}

func NewFromFloat64(val float64) Angle {
	return NewFromDecimal(decimal.NewFromFloat(val))
}

func NewFromString(str string) (Angle, error) {
	var deg Angle

	if err := deg.scanByString(str); err != nil {
		return Angle{}, err
	}

	return deg, nil
}

func NewFromDegreeMinuteSecond(degree, minute, second decimal.Decimal) Angle {
	neg := false

	if degree.IsNegative() {
		neg = true
		degree = degree.Abs()
	}

	if minute.IsNegative() {
		neg = true
		minute = minute.Abs()
	}

	if second.IsNegative() {
		neg = true
		second = second.Abs()
	}

	return Angle{
		degree: degree,
		minute: minute,
		second: second,

		neg:     neg,
		angType: angleType.DegreeMinuteSecond,
	}
}
