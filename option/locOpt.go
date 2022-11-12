package option

import (
	"github.com/shopspring/decimal"
	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
)

type LocOpt struct {
	Latitude  angle.Angle
	Longitude angle.Angle
	Elevation decimal.Decimal
	Timezone  decimal.Decimal
}
