package currency

import (
	"context"
	"github.com/shopspring/decimal"
)

type RateProvider interface {
	GetRate(ctx context.Context, pair string) (decimal.Decimal, error)
}

type CurrencyRates struct {
	Pair string
	Rate decimal.Decimal
	Err  error
}
