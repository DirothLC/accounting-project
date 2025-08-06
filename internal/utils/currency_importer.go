package utils

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"sync"
	"time"
)

type CurrencyRates struct {
	Pair string
	Rate decimal.Decimal
	Err  error
}

var (
	mu       sync.Mutex
	rateMap  = make(map[string]decimal.Decimal)
	wg       sync.WaitGroup
	resultCh = make(chan CurrencyRates)
)

var pairs = []string{"USD/KZT", "RUB/KZT", "USD/RUB", "EUR/KZT", "EUR/USD"}

func fetchRate(pair string) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	done := make(chan CurrencyRates, 1)
	go func() {
		sleepDuration := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(sleepDuration)

		rate, ok := Rates[pair]
		if !ok {
			done <- CurrencyRates{Pair: pair, Err: fmt.Errorf("Currency pair %s not found", pair)}
			return
		}
		done <- CurrencyRates{
			Pair: pair,
			Rate: decimal.NewFromFloat(rate),
			Err:  nil,
		}
	}()
	select {
	case <-ctx.Done():
		resultCh <- CurrencyRates{Pair: pair, Err: fmt.Errorf("Request timeout")}
	case res := <-done:
		resultCh <- res
	}
}
