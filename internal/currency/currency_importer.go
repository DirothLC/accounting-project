package currency

import (
	"github.com/shopspring/decimal"
	"log"
	"time"
)

func ImportCurrencyRates() map[string]decimal.Decimal {
	mockProvider := MockRateProvider{
		Rates: map[string]decimal.Decimal{
			"USD/KZT": decimal.NewFromFloat(540.42),
			"RUB/KZT": decimal.NewFromFloat(6.71),
			"USD/RUB": decimal.NewFromFloat(80.26),
			"EUR/KZT": decimal.NewFromFloat(610.93),
			"EUR/USD": decimal.NewFromFloat(1.15),
		},
	}

	pairs := []string{"USD/KZT", "RUB/KZT", "USD/RUB", "EUR/KZT", "EUR/USD"}

	log.Println("Imported currency rates:")

	results, errorsMap := FetchAllRates(mockProvider, pairs, 1500*time.Millisecond)

	for _, p := range pairs {
		if rate, ok := results[p]; ok {
			log.Printf("%s: %s", p, rate.StringFixed(2))
		} else if err, ok := errorsMap[p]; ok {
			log.Printf("%s: %v", p, err)
		} else {
			log.Printf("%s: unknown error", p)
		}
	}

	return results
}
