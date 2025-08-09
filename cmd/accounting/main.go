package main

import (
	"Accounting/config"
	"Accounting/internal/currency"
	"Accounting/internal/database"
)

func main() {
	config.InitConfig()
	database.Connect(config.Cfg.DbURL)
	if config.Cfg.AutoMigrate {
		database.Migrate()
	}
	/*var transaction = examples.CreateTxExample()
	service.ChangeTransactionStatus(transaction.ID, "AUTH")
	service.ChargeTransaction(transaction.ID)
	service.RefundTransaction(transaction.ID)*/

	currency.ImportCurrencyRates()
}
