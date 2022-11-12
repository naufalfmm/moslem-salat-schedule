package option

import (
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
)

type LocOpt struct {
	Latitude  angle.Angle
	Longitude angle.Angle
	Elevation float64
	Timezone  float64
}
