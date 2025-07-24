package main

import (
	"Accounting/config"
	"Accounting/internal/database"
	"Accounting/internal/examples"
	"Accounting/internal/service"
)

func main() {
	config.InitConfig()
	database.Connect(config.Cfg.DbURL)
	if config.Cfg.AutoMigrate {
		database.Migrate()
	}
	var transaction = examples.CreateTxExample()
	service.ChangeTransactionStatus(transaction.ID, "AUTH")
	service.ChargeTransaction(transaction.ID)
	service.CancelTransaction(transaction.ID)
}
