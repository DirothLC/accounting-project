package utils

import (
	"Accounting/internal/enums"
	"github.com/shopspring/decimal"
)

const (
	RUB = 6.70
	USD = 522.37
	EUR = 610.93
)

func KztToConverter(amount decimal.Decimal, currencyType enums.CurrencyType) decimal.Decimal {
	switch currencyType {
	case enums.USD:
		return amount.Div(decimal.NewFromFloat(USD))
	case enums.EUR:
		return amount.Div(decimal.NewFromFloat(EUR))
	case enums.RUB:
		return amount.Div(decimal.NewFromFloat(RUB))
	default:
		return amount
	}
}
func ToKztConverter(amount decimal.Decimal, currencyType enums.CurrencyType) decimal.Decimal {
	switch currencyType {
	case enums.USD:
		return amount.Mul(decimal.NewFromFloat(USD))
	case enums.EUR:
		return amount.Mul(decimal.NewFromFloat(EUR))
	case enums.RUB:
		return amount.Mul(decimal.NewFromFloat(RUB))
	default:
		return amount
	}

}
