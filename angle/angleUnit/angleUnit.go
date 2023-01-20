package angleUnit

type (
	AngleUnitClass struct {
		Code string
		Name string
		Unit string
	}

	AngleUnit int
)

const (
	Degree AngleUnit = iota + 1
	Radian
)

var (
	angleUnitConsts = []AngleUnitClass{
		{"degree", "Degree", "°"},
		{"radian", "Radian", "rad"},
	}
)

// Code .
func (c AngleUnit) Code() string {
	if c < 1 || int(c) > len(angleUnitConsts) {
		return ""
	}
	return angleUnitConsts[c-1].Code
}

// Name .
func (c AngleUnit) Name() string {
	if c < 1 || int(c) > len(angleUnitConsts) {
		return ""
	}
	return angleUnitConsts[c-1].Name
}

// Unit .
func (c AngleUnit) Unit() string {
	if c < 1 || int(c) > len(angleUnitConsts) {
		return ""
	}
	return angleUnitConsts[c-1].Unit
}
