package utils

import (
	"math"

	"github.com/shopspring/decimal"
)

func Atan2(y, x decimal.Decimal) decimal.Decimal {
	if y.IsZero() {
		if x.IsNegative() {
			return decimal.NewFromFloat(math.Pi)
		}
	}

	if x.IsZero() {
		if y.IsPositive() {
			return decimal.NewFromFloat(math.Pi).Div(decimal.NewFromFloat(2.0))
		}

		if y.IsNegative() {
			return decimal.NewFromFloat(math.Pi).Div(decimal.NewFromFloat(2.0)).Neg()
		}
	}

	return y.Div(x).Atan()
}
