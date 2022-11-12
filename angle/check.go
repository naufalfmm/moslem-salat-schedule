package angle

func (d Angle) IsNegative() bool {
	return d.neg
}

func (d Angle) IsZero() bool {
	return d.degree.IsZero() || d.minute.IsZero() || d.second.IsZero()
}

func (d Angle) IsPositive() bool {
	return !d.neg
}

func (d Angle) Equal(d1 Angle) bool {
	return d.neg && d1.neg &&
		d.degree.Equal(d1.degree) &&
		d.minute.Equal(d1.minute) &&
		d.second.Equal(d1.second)
}

func (d Angle) GreatherThan(d1 Angle) bool {
	if d1.neg && !d.neg {
		return true
	}

	if d.degree.GreaterThan(d1.degree) {
		return true
	}

	if d.minute.GreaterThan(d1.minute) {
		return true
	}

	if d.second.GreaterThan(d1.second) {
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
