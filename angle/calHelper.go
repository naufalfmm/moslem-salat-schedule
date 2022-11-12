package angle

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
)

func (d Angle) addToDecimalType(deg Angle) Angle {
	if deg.angType != angleType.Decimal {
		deg = deg.ToDecimal()
	}

	return Angle{
		degree:  d.degree + deg.degree,
		angType: angleType.Decimal,
	}
}

func (d Angle) addToMinuteSecondType(deg Angle) Angle {
	if deg.angType != angleType.DegreeMinuteSecond {
		deg = deg.ToMinuteSecond()
	}

	second := d.second + deg.second
	minute := d.minute + deg.minute
	degree := d.degree + deg.degree

	return Angle{
		degree:  degree,
		minute:  minute,
		second:  second,
		angType: angleType.DegreeMinuteSecond,
	}.prepareConvertMinuteSecond()
}

func (d Angle) addToAugendType(d1 Angle) Angle {
	if d.neg && d1.neg {
		return d.Abs().addToAugendType(d1.Abs()).Neg()
	}

	if d.neg {
		return d1.ToSpecificType(d.angType).Sub(d.Abs())
	}

	if d1.neg {
		return d.Sub(d1.Abs())
	}

	if d.angType == angleType.Decimal {
		return d.addToDecimalType(d1)
	}

	return d.addToMinuteSecondType(d1)
}

func (d Angle) subToDecimalType(d1 Angle) Angle {
	if d1.angType != angleType.Decimal {
		d1 = d1.ToDecimal()
	}

	if d1.GreatherThan(d) {
		return d1.subToDecimalType(d).Neg()
	}

	return Angle{
		degree:  math.Abs(d.degree - d1.degree),
		angType: angleType.Decimal,
	}
}

func takeForSub(value, upperValue float64) (float64, float64) {
	if upperValue == consts.DecimalZero {
		return value, upperValue
	}

	return value + consts.TimeFormatConverter, upperValue - consts.DecimalOne
}

func (d Angle) prepareMinuend(d1 Angle) Angle {
	second := d.second
	minute := d.minute
	degree := d.degree

	if second < d1.second {
		second, minute = takeForSub(second, minute)
		if second == d.second {
			minute, degree = takeForSub(minute, degree)
		}
	}

	if minute < d1.minute {
		minute, degree = takeForSub(minute, degree)
	}

	return Angle{
		degree:  degree,
		minute:  minute,
		second:  second,
		angType: angleType.DegreeMinuteSecond,
	}
}

func (d Angle) subToMinuteSecondType(d1 Angle) Angle {
	if d1.angType != angleType.DegreeMinuteSecond {
		d1 = d1.ToMinuteSecond()
	}

	d = d.prepareMinuend(d1)

	if d1.GreatherThan(d) {
		return d1.subToMinuteSecondType(d).Neg()
	}

	return Angle{
		degree:  math.Abs(d.degree - d1.degree),
		minute:  math.Abs(d.minute - d1.minute),
		second:  math.Abs(d.second - d1.second),
		angType: angleType.DegreeMinuteSecond,
	}.prepareConvertMinuteSecond()
}

func (d Angle) subToMinuendType(d1 Angle) Angle {
	if d.neg && d1.neg {
		return d1.Abs().ToSpecificType(d.angType).subToMinuendType(d1.Abs())
	}

	if d.neg {
		return d.Abs().addToAugendType(d1.Abs()).Neg()
	}

	if d1.neg {
		return d.Abs().addToAugendType(d1.Abs())
	}

	if d1.angType == angleType.Decimal {
		return d.subToDecimalType(d1)
	}

	return d.subToMinuteSecondType(d1)
}
