package angle

import (
	"github.com/shopspring/decimal"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
)

func (d Angle) ToMinuteSecond() Angle {
	if d.angType == angleType.DegreeMinuteSecond {
		return d
	}

	decDegree := Angle{
		degree:  d.degree.Truncate(0),
		neg:     d.neg,
		angType: angleType.DegreeMinuteSecond,
	}

	restDegree := d.degree.Sub(decDegree.degree).Mul(decimal.NewFromInt(60))
	if restDegree.GreaterThan(decimal.Zero) {
		decDegree.minute = restDegree.Truncate(0)
		restDegree = restDegree.Sub(decDegree.minute).Mul(decimal.NewFromInt(60))
	}

	if restDegree.GreaterThan(decimal.Zero) {
		decDegree.second = restDegree
	}

	return decDegree
}

func (d Angle) ToDecimal() Angle {
	if d.angType == angleType.Decimal {
		return d
	}

	return Angle{
		degree:  d.degree.Add(d.minute.Div(decimal.NewFromInt(60))).Add(d.second.Div(decimal.NewFromInt(3600))),
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
		degree:  d1.degree.Floor(),
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
		degree:  d1.degree.Ceil(),
		neg:     d1.neg,
		angType: d1.angType,
	}

	return d1.ToSpecificType(d.angType)
}
