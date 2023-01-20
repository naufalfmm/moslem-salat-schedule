package angle

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
)

func (a Angle) addToDecimalType(ang Angle) Angle {
	if ang.angType != angleType.Decimal {
		ang = ang.ToDecimal()
	}

	if ang.angUnit != a.angUnit {
		ang = ang.ToSpecificUnit(a.angUnit)
	}

	return Angle{
		degree:  a.degree + ang.degree,
		angType: angleType.Decimal,
		angUnit: a.angUnit,
	}
}

func (a Angle) addToMinuteSecondType(ang Angle) Angle {
	if ang.angType != angleType.DegreeMinuteSecond {
		ang = ang.ToMinuteSecond()
	}

	if ang.angUnit != a.angUnit {
		ang = ang.ToSpecificUnit(a.angUnit)
	}

	second := a.second + ang.second
	minute := a.minute + ang.minute
	degree := a.degree + ang.degree

	return Angle{
		degree:  degree,
		minute:  minute,
		second:  second,
		angType: angleType.DegreeMinuteSecond,
		angUnit: a.angUnit,
	}.prepareConvertMinuteSecond()
}

func (a Angle) addToAugendTypeUnit(a1 Angle) Angle {
	if a.neg && a1.neg {
		return a.Abs().addToAugendTypeUnit(a1.Abs()).Neg()
	}

	if a.neg {
		return a1.ToSpecificType(a.angType).Sub(a.Abs())
	}

	if a1.neg {
		return a.Sub(a1.Abs())
	}

	if a.angType == angleType.Decimal {
		return a.addToDecimalType(a1)
	}

	return a.addToMinuteSecondType(a1)
}

func (a Angle) subToDecimalType(a1 Angle) Angle {
	if a1.angType != angleType.Decimal {
		a1 = a1.ToDecimal()
	}

	if a1.angUnit != a.angUnit {
		a1 = a1.ToSpecificUnit(a.angUnit)
	}

	if a1.GreatherThan(a) {
		return a1.subToDecimalType(a).Neg()
	}

	return Angle{
		degree:  math.Abs(a.degree - a1.degree),
		angType: angleType.Decimal,
		angUnit: a.angUnit,
	}
}

func takeForSub(value, upperValue float64) (float64, float64) {
	if upperValue == consts.DecimalZero {
		return value, upperValue
	}

	return value + consts.TimeFormatConverter, upperValue - consts.DecimalOne
}

func (a Angle) prepareMinuend(a1 Angle) Angle {
	second := a.second
	minute := a.minute
	degree := a.degree

	if second < a1.second {
		second, minute = takeForSub(second, minute)
		if second == a.second {
			minute, degree = takeForSub(minute, degree)
		}
	}

	if minute < a1.minute {
		minute, degree = takeForSub(minute, degree)
	}

	return Angle{
		degree:  degree,
		minute:  minute,
		second:  second,
		angType: angleType.DegreeMinuteSecond,
	}
}

func (a Angle) subToMinuteSecondType(a1 Angle) Angle {
	if a1.angType != angleType.DegreeMinuteSecond {
		a1 = a1.ToMinuteSecond()
	}

	if a1.angUnit != a.angUnit {
		a1 = a1.ToSpecificUnit(a.angUnit)
	}

	a = a.prepareMinuend(a1)

	if a1.GreatherThan(a) {
		return a1.subToMinuteSecondType(a).Neg()
	}

	return Angle{
		degree:  math.Abs(a.degree - a1.degree),
		minute:  math.Abs(a.minute - a1.minute),
		second:  math.Abs(a.second - a1.second),
		angType: angleType.DegreeMinuteSecond,
		angUnit: a.angUnit,
	}.prepareConvertMinuteSecond()
}

func (a Angle) subToMinuendTypeUnit(a1 Angle) Angle {
	if a.neg && a1.neg {
		return a1.Abs().ToSpecificType(a.angType).subToMinuendTypeUnit(a.Abs())
	}

	if a.neg {
		return a.Abs().addToAugendTypeUnit(a1.Abs()).Neg()
	}

	if a1.neg {
		return a.Abs().addToAugendTypeUnit(a1.Abs())
	}

	if a1.angType == angleType.Decimal {
		return a.subToDecimalType(a1)
	}

	return a.subToMinuteSecondType(a1)
}

func (a Angle) divToDividendTypeUnit(d float64) Angle {
	angType := a.angType

	if a.angType != angleType.Decimal {
		a = a.ToDecimal()
	}

	a.degree = a.degree / d

	return a.ToSpecificType(angType)
}

func (a Angle) mulToMultiplierTypeUnit(d float64) Angle {
	angType := a.angType

	if a.angType != angleType.Decimal {
		a = a.ToDecimal()
	}

	a.degree *= d

	return a.ToSpecificType(angType)
}
