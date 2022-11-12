package consts

import "github.com/shopspring/decimal"

const (
	DegreeSymbolRune   = '°'
	MinuteSymbolRune   = '\''
	SecondSymbolRune   = '"'
	NegativeSymbolRune = '-'
)

var (
	TimeFormatConverter = decimal.NewFromInt(60)
	DecimalOne          = decimal.NewFromInt(1)
)
