package adapters

import "log"

type MockBank struct{}

func (m *MockBank) Auth(clientID, clientSecret string) bool {
	return true
}
func (m *MockBank) Charge(amount float64) error {
	log.Printf("Mock charge of amount: %.2f", amount)
	return nil
}

func (m *MockBank) Refund(transactionID int64) error {
	log.Printf("Mock refund of transaction: %d", transactionID)
	return nil
}

func init() {
	RegisterAdapter("MockBank", &MockBank{})
}
