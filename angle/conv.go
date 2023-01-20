package angle

import (
	"math"
	"time"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleUnit"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/err"
)

func (d Angle) ToMinuteSecond() Angle {
	if d.angUnit != angleUnit.Degree {
		panic(err.ErrShouldBeDegree)
	}

	if d.angType == angleType.DegreeMinuteSecond {
		return d
	}

	decDegree := Angle{
		degree:  math.Trunc(d.degree),
		neg:     d.neg,
		angType: angleType.DegreeMinuteSecond,
		angUnit: d.angUnit,
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
		angUnit: d.angUnit,
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

func (a Angle) ToRadian() Angle {
	if a.angUnit == angleUnit.Radian {
		return a
	}

	if a.angType != angleType.Decimal {
		panic(err.ErrShouldBeDecimal)
	}

	return Angle{
		degree:  a.degree * math.Pi / 180.0,
		neg:     a.neg,
		angType: a.angType,
		angUnit: angleUnit.Radian,
	}
}

func (a Angle) ToDegree() Angle {
	if a.angUnit == angleUnit.Degree {
		return a
	}

	if a.angType != angleType.Decimal {
		panic(err.ErrShouldBeDecimal)
	}

	return Angle{
		degree:  a.degree * 180.0 / math.Pi,
		neg:     a.neg,
		angType: a.angType,
		angUnit: angleUnit.Degree,
	}
}

func (a Angle) ToOtherUnit() Angle {
	if a.angUnit == angleUnit.Radian {
		return a.ToDegree()
	}

	return a.ToRadian()
}

func (a Angle) ToSpecificUnit(unit angleUnit.AngleUnit) Angle {
	if a.angUnit == unit {
		return a
	}

	if unit == angleUnit.Radian {
		return a.ToRadian()
	}

	return a.ToDegree()
}

func (a Angle) ToTime() time.Time {
	if a.angType != angleType.DegreeMinuteSecond {
		a = a.ToSpecificType(angleType.DegreeMinuteSecond)
	}

	now := time.Now()

	return time.Date(now.Year(), now.Month(), now.Day(), int(a.degree), int(a.minute), int(a.second), 0, now.Location())
}

func (a Angle) ToFloat() float64 {
	fact := 1.
	if a.neg {
		fact = -1.
	}

	return a.ToDecimal().degree * fact
}
