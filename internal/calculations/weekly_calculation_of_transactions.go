package calculations

import (
	"fmt"
	"github.com/shopspring/decimal"
)

var day = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
}

func WeeklyCalculator(transactions []map[int]decimal.Decimal) decimal.Decimal {
	var total decimal.Decimal
	for _, transaction := range transactions {
		for txDay, amount := range transaction {
			dayReporter(amount, txDay)
			total = total.Add(amount)
		}
	}
	return total
}
func dayReporter(amount decimal.Decimal, txDay int) {
	dayName, exists := day[txDay]
	if !exists {
		dayName = "Unknown day"
	}
	fmt.Println(dayName + ":")
	if amount.GreaterThan(decimal.Zero) {
		fmt.Println("Positive amount:", amount, "\n")
	} else {
		fmt.Println("Negative amount:", amount, "\n")
	}

}
