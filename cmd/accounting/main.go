package main

import (
	"Accounting/config"
	"Accounting/internal/database"
	"Accounting/internal/examples"
)

func main() {
	config.InitConfig()
	database.Connect(config.Cfg.DbURL)
	if config.Cfg.AutoMigrate {
		database.Migrate()
	}
	examples.CreateTxExample()

}
