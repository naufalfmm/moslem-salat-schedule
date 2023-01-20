package salatHighAltitude

import (
	"math"

	"gitlab.com/naufalfmm/moslem-salat-schedule/angle"
)

func CalcSalatHighAltitude(angleFact, lat, dec angle.Angle, elev float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Acos((math.Sin(angleFact.Neg().Sub(angle.NewDegreeFromFloat(0.0347).Mul(math.Sqrt(elev))).ToRadian().ToFloat()) - (math.Sin(lat.ToRadian().ToFloat()) * math.Sin(dec.ToFloat()))) / (math.Cos(lat.ToRadian().ToFloat()) * math.Cos(dec.ToFloat())))).ToDegree().Div(15.)
}
