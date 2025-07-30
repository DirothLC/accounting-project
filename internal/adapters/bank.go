package adapters

type BankAdapter interface {
	Auth(clientID, clientSecret string) bool
	Charge(amount float64) error
	Refund(transactionID int64) error
}

var Adapters = make(map[string]BankAdapter)

func RegisterAdapter(name string, adapter BankAdapter) {
	Adapters[name] = adapter
}
