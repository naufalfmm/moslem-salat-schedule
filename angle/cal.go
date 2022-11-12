package angle

import "gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"

func (a Angle) Add(a1 Angle) Angle {
	return a.addToAugendType(a1)
}

func (a Angle) AddToSpecificType(a1 Angle, angType angleType.AngleType) Angle {
	if a.angType != angType {
		return a.ToOtherType().addToAugendType(a1)
	}

	return a.addToAugendType(a1)
}

func (a Angle) Sub(a1 Angle) Angle {
	return a.subToMinuendType(a1)
}

func (a Angle) SubToSpecificType(a1 Angle, angType angleType.AngleType) Angle {
	if a.angType != angType {
		return a.ToOtherType().subToMinuendType(a1)
	}

	return a.subToMinuendType(a1)
}

func (a Angle) Div(d float64) Angle {
	return a.divToDividendType(d)
}

func (a Angle) DivToSpecificType(d float64, angType angleType.AngleType) Angle {
	if a.angType != angType {
		return a.ToSpecificType(angType).divToDividendType(d)
	}

	return a.divToDividendType(d)
}
