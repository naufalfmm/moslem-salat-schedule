package angle

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleUnit"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle/err"
)

func (a Angle) fullRotateRadian() Angle {
	if a.angType != angleType.Decimal {
		panic(err.ErrShouldBeDecimal)
	}

	a = a.ToDegree()

	return Angle{
		degree:  math.Mod(a.ToFloat(), 360),
		neg:     false,
		angType: angleType.Decimal,
		angUnit: angleUnit.Degree,
	}.ToRadian()
}

func (a Angle) fullRotateDegree() Angle {
	ang := Angle{
		degree:  math.Mod(a.ToDecimal().ToFloat(), 360),
		neg:     false,
		angType: angleType.Decimal,
		angUnit: angleUnit.Degree,
	}

	if a.angType != angleType.DegreeMinuteSecond {
		return ang
	}

	return ang.ToMinuteSecond()
}

func (a Angle) FullRotate() Angle {
	if a.angUnit == angleUnit.Radian {
		return a.fullRotateRadian()
	}

	return a.fullRotateDegree()
}
