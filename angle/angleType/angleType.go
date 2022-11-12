package angleType

type (
	AngleTypeClass struct {
		Code string
		Name string
	}

	AngleType int
)

const (
	Decimal AngleType = iota + 1
	DegreeMinuteSecond
)

var (
	degreeTypeConsts = []AngleTypeClass{
		{"decimal", "Decimal"},
		{"degreeMinuteSecond", "Degree Minute Second"},
	}
)

// Code .
func (c AngleType) Code() string {
	if c < 1 || int(c) > len(degreeTypeConsts) {
		return ""
	}
	return degreeTypeConsts[c-1].Code
}

// Name .
func (c AngleType) Name() string {
	if c < 1 || int(c) > len(degreeTypeConsts) {
		return ""
	}
	return degreeTypeConsts[c-1].Name
}
