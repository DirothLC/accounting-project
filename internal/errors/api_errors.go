package errors

import "fmt"

type RateNotFoundError struct {
	Pair string
}

func (e RateNotFoundError) Error() string {
	return fmt.Sprintf("currency pair %s not found", e.Pair)
}
