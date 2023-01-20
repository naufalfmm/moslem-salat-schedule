package angle

import "gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"

func (d Angle) IsNegative() bool {
	return d.neg
}

func (d Angle) IsZero() bool {
	return d.degree == consts.DecimalZero && d.minute == consts.DecimalZero && d.second == consts.DecimalZero
}

func (d Angle) IsPositive() bool {
	return !d.neg
}

func (d Angle) Equal(d1 Angle) bool {
	return d.neg && d1.neg &&
		d.degree == d1.degree &&
		d.minute == d1.minute &&
		d.second == d1.second
}

func (d Angle) GreatherThan(d1 Angle) bool {
	if d1.neg && !d.neg {
		return true
	}

	if d.degree > d1.degree {
		return true
	}

	if d.minute > d1.minute {
		return true
	}

	if d.second > d1.second {
		return true
	}

	return false
}

func (d Angle) GreaterThanOrEqual(d1 Angle) bool {
	return d.GreatherThan(d1) || d.Equal(d1)
}

func (d Angle) LessThan(d1 Angle) bool {
	return !d.GreatherThan(d1)
}

func (d Angle) LessThanOrEqual(d1 Angle) bool {
	return d.LessThan(d1) || d.Equal(d1)
}
