package angle

import "gitlab.com/naufalfmm/moslem-salat-schedule/angle/angleType"

func (d Angle) Add(d1 Angle) Angle {
	return d.addToAugendType(d1)
}

func (d Angle) AddToSpecificType(d1 Angle, angType angleType.AngleType) Angle {
	if d.angType != angType {
		return d.ToOtherType().addToAugendType(d1)
	}

	return d.addToAugendType(d1)
}

func (d Angle) Sub(d1 Angle) Angle {
	return d.subToMinuendType(d1)
}

func (d Angle) SubToSpecificType(d1 Angle, degType angleType.AngleType) Angle {
	if d.angType != degType {
		return d.ToOtherType().subToMinuendType(d1)
	}

	return d.subToMinuendType(d1)
}
