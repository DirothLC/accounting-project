package examples

/*
This file implements a visual demonstration of tools for counting weekly transactions.
Focusing on input values, you can understand the principle of work for further adaptation to your needs.
*/

import (
	"Accounting/internal/calculations"
	"fmt"
	"github.com/shopspring/decimal"
)

func CalculationExample() {
	transactions := []map[int]decimal.Decimal{
		{1: decimal.NewFromFloat(25000.00)},
		{1: decimal.NewFromFloat(20000.00)},
		{2: decimal.NewFromFloat(-9800.00)},
		{3: decimal.NewFromFloat(-1222.22)},
		{4: decimal.NewFromFloat(-1500.07)},
		{5: decimal.NewFromFloat(1201.37)},
		{6: decimal.NewFromFloat(-100.32)},
		{7: decimal.NewFromFloat(-523.33)},
	}
	total := calculations.WeeklyCalculator(transactions)
	fmt.Println("Weekly result:\n", total, "\n")
}
