package angle

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
)

var Zero = NewFromFloat64(consts.DecimalZero)

// func NewFromDecimal(dec float64) Angle {
// 	return Angle{
// 		degree:  dec.Abs(),
// 		neg:     dec.IsNegative(),
// 		angType: angleType.Decimal,
// 	}
// }

func NewFromFloat64(val float64) Angle {
	return Angle{
		degree:  math.Abs(val),
		neg:     val < 0,
		angType: angleType.Decimal,
	}
}

func NewFromString(str string) (Angle, error) {
	var deg Angle

	if err := deg.scanByString(str); err != nil {
		return Angle{}, err
	}

	return deg, nil
}

func NewFromDegreeMinuteSecond(degree, minute, second float64) Angle {
	neg := false

	if degree < 0 {
		neg = true
		degree = math.Abs(degree)
	}

	if minute < 0 {
		neg = true
		minute = math.Abs(minute)
	}

	if second < 0 {
		neg = true
		second = math.Abs(second)
	}

	return Angle{
		degree: degree,
		minute: minute,
		second: second,

		neg:     neg,
		angType: angleType.DegreeMinuteSecond,
	}
}
