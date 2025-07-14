package logic

import (
	"Accounting/pkg/enums"
	"Accounting/pkg/utils"
	"fmt"
	"github.com/shopspring/decimal"
)

func ConverterExamples() {
	fmt.Println("Select conversion:\n 1. KZT to...\n 2. ... to KZT")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	switch choice {
	case 1:
		convertAmount(true)
	case 2:
		convertAmount(false)
	default:
		fmt.Println("Invalid choice")
	}
}

func convertAmount(fromKZT bool) {
	var direction string
	if fromKZT {
		direction = "KZT to ..."
	} else {
		direction = "... to KZT"
	}
	currency, err := chooseCurrency(direction)
	if err != nil {
		fmt.Println(err)
		return
	}

	var amount decimal.Decimal
	if fromKZT {
		amount, err = readAmount("Enter amount in KZT")
	} else {
		amount, err = readAmount("Enter amount in selected currency")
	}
	if err != nil {
		fmt.Println("Invalid amount:", err)
		return
	}

	var result decimal.Decimal
	if fromKZT {
		result = utils.KztToConverter(amount, currency)
	} else {
		result = utils.ToKztConverter(amount, currency)
	}

	fmt.Printf("Converted amount: %s\n", result)
}

func readAmount(prompt string) (decimal.Decimal, error) {
	fmt.Println(prompt)
	var scan float64
	_, err := fmt.Scanln(&scan)
	amount := decimal.NewFromFloat(scan)
	return amount, err
}
func chooseCurrency(prompt string) (enums.CurrencyType, error) {
	fmt.Println(prompt)
	fmt.Println("1. USD\n2. EUR\n3. RUB")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return 0, err
	}

	switch choice {
	case 1:
		return enums.USD, nil
	case 2:
		return enums.EUR, nil
	case 3:
		return enums.RUB, nil
	default:
		return 0, fmt.Errorf("invalid choice")
	}
}
