package currency

import (
	"Accounting/internal/errors"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"math/rand"
	"sync"
	"time"
)

type MockRateProvider struct {
	Rates map[string]decimal.Decimal
}

func (m MockRateProvider) GetRate(ctx context.Context, pair string) (decimal.Decimal, error) {
	sleepDuration := time.Duration(rand.Intn(2000)) * time.Millisecond
	select {
	case <-time.After(sleepDuration):
		r, ok := m.Rates[pair]
		if !ok {
			err := errors.RateNotFoundError{Pair: pair}
			log.Printf("[ERROR] %v", err)
			return decimal.Zero, err
		}
		return r, nil
	case <-ctx.Done():
		err := fmt.Errorf("timeout fetching %s", pair)
		log.Printf("[ERROR] %v", err)
		return decimal.Zero, err
	}
}

func FetchAllRates(provider RateProvider, pairs []string, timeout time.Duration) (map[string]decimal.Decimal, map[string]error) {
	var (
		mu        sync.Mutex
		wg        sync.WaitGroup
		results   = make(map[string]decimal.Decimal)
		errorsMap = make(map[string]error)
	)
	resultCh := make(chan CurrencyRates)
	for _, pair := range pairs {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			rate, err := provider.GetRate(ctx, p)
			resultCh <- CurrencyRates{Pair: p, Rate: rate, Err: err}
		}(pair)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	for res := range resultCh {
		mu.Lock()
		if res.Err == nil {
			results[res.Pair] = res.Rate
		} else {
			errorsMap[res.Pair] = res.Err
		}
		mu.Unlock()
	}

	return results, errorsMap
}
