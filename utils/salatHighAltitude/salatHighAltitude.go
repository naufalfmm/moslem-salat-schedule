package salatHighAltitude

import (
	"math"

	"github.com/naufalfmm/angle"
	"github.com/naufalfmm/angle/trig"
)

func CalcSalatHighAltitude(angleFactor, lat, dec angle.Angle, elev float64) angle.Angle {
	return trig.Acos((trig.Sin(angleFactor.Neg().SubScalar(0.0347*math.Sqrt(elev))) - trig.Sin(lat)*trig.Sin(dec)) / (trig.Cos(lat) * trig.Cos(dec))).Div(15.)
}
