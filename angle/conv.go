package angle

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
)

func (d Angle) ToMinuteSecond() Angle {
	if d.angType == angleType.DegreeMinuteSecond {
		return d
	}

	decDegree := Angle{
		degree:  math.Trunc(d.degree),
		neg:     d.neg,
		angType: angleType.DegreeMinuteSecond,
	}

	restDegree := (d.degree - decDegree.degree) * consts.TimeFormatConverter
	if restDegree > consts.DecimalZero {
		decDegree.minute = math.Trunc(restDegree)
		restDegree = (restDegree - decDegree.minute) * consts.TimeFormatConverter
	}

	if restDegree > consts.DecimalZero {
		decDegree.second = restDegree
	}

	return decDegree
}

func (d Angle) ToDecimal() Angle {
	if d.angType == angleType.Decimal {
		return d
	}

	return Angle{
		degree:  d.degree + (d.minute / consts.TimeFormatConverter) + (d.second / (consts.TimeFormatConverter * consts.TimeFormatConverter)),
		neg:     d.neg,
		angType: angleType.Decimal,
	}
}

func (d Angle) ToOtherType() Angle {
	if d.angType == angleType.Decimal {
		return d.ToMinuteSecond()
	}

	return d.ToDecimal()
}

func (d Angle) ToSpecificType(angType angleType.AngleType) Angle {
	if d.angType == angType {
		return d
	}

	if angType == angleType.Decimal {
		return d.ToDecimal()
	}

	return d.ToMinuteSecond()
}

func (d Angle) Abs() Angle {
	d.neg = false
	return d
}

func (d Angle) Neg() Angle {
	d.neg = true
	return d
}

func (d Angle) Floor() Angle {
	d1 := d
	if d.angType != angleType.Decimal {
		d1 = d.ToDecimal()
	}

	d1 = Angle{
		degree:  math.Floor(d1.degree),
		neg:     d1.neg,
		angType: d1.angType,
	}

	return d1.ToSpecificType(d.angType)
}

func (d Angle) Ceil() Angle {
	d1 := d
	if d.angType != angleType.Decimal {
		d1 = d.ToDecimal()
	}

	d1 = Angle{
		degree:  math.Ceil(d1.degree),
		neg:     d1.neg,
		angType: d1.angType,
	}

	return d1.ToSpecificType(d.angType)
}
