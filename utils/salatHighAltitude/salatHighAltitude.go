package salatHighAltitude

import (
	"math"

	"github.com/naufalfmm/angle"
)

func CalcSalatHighAltitude(angleFactor, lat, dec angle.Angle, elev float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Acos((math.Sin(angleFactor.Neg().Sub(angle.NewDegreeFromFloat(0.0347).Mul(math.Sqrt(elev))).ToRadian().ToFloat()) - (math.Sin(lat.ToRadian().ToFloat()) * math.Sin(dec.ToFloat()))) / (math.Cos(lat.ToRadian().ToFloat()) * math.Cos(dec.ToFloat())))).ToDegree().Div(15.)
}
