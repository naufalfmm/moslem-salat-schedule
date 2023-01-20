package angle

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleUnit"
)

func (a Angle) Add(a1 Angle) Angle {
	return a.addToAugendTypeUnit(a1)
}

func (a Angle) AddToSpecificType(a1 Angle, angType angleType.AngleType) Angle {
	a1 = a1.ToSpecificType(angType)

	if a.angType != angType {
		return a.ToSpecificType(angType).addToAugendTypeUnit(a1)
	}

	return a.addToAugendTypeUnit(a1)
}

func (a Angle) AddToSpecificUnit(a1 Angle, angUnit angleUnit.AngleUnit) Angle {
	a1 = a1.ToSpecificUnit(angUnit)

	if a.angUnit != angUnit {
		return a.ToSpecificUnit(angUnit).addToAugendTypeUnit(a1)
	}

	return a.addToAugendTypeUnit(a1)
}

func (a Angle) Sub(a1 Angle) Angle {
	return a.subToMinuendTypeUnit(a1)
}

func (a Angle) SubToSpecificType(a1 Angle, angType angleType.AngleType) Angle {
	a1 = a1.ToSpecificType(angType)

	if a.angType != angType {
		return a.ToSpecificType(angType).subToMinuendTypeUnit(a1)
	}

	return a.subToMinuendTypeUnit(a1)
}

func (a Angle) SubToSpecificUnit(a1 Angle, angUnit angleUnit.AngleUnit) Angle {
	a1 = a1.ToSpecificUnit(angUnit)

	if a.angUnit != angUnit {
		return a.ToSpecificUnit(angUnit).subToMinuendTypeUnit(a1)
	}

	return a.subToMinuendTypeUnit(a1)
}

func (a Angle) Div(d float64) Angle {
	return a.divToDividendTypeUnit(d)
}

func (a Angle) DivToSpecificType(d float64, angType angleType.AngleType) Angle {
	if a.angType != angType {
		return a.ToSpecificType(angType).divToDividendTypeUnit(d)
	}

	return a.divToDividendTypeUnit(d)
}

func (a Angle) DivToSpecificUnit(d float64, angUnit angleUnit.AngleUnit) Angle {
	if a.angUnit != angUnit {
		return a.ToSpecificUnit(angUnit).divToDividendTypeUnit(d)
	}

	return a.divToDividendTypeUnit(d)
}

func (a Angle) Mul(d float64) Angle {
	return a.mulToMultiplierTypeUnit(d)
}

func (a Angle) MulToSpecificType(m float64, angType angleType.AngleType) Angle {
	if a.angType != angType {
		return a.ToSpecificType(angType).mulToMultiplierTypeUnit(m)
	}

	return a.mulToMultiplierTypeUnit(m)
}

func (a Angle) MulToSpecificUnit(m float64, angUnit angleUnit.AngleUnit) Angle {
	if a.angUnit != angUnit {
		return a.ToSpecificUnit(angUnit).mulToMultiplierTypeUnit(m)
	}

	return a.mulToMultiplierTypeUnit(m)
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
		angUnit: d1.angUnit,
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
		angUnit: d1.angUnit,
	}

	return d1.ToSpecificType(d.angType)
}
