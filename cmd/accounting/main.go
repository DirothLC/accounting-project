package main

import (
	"Accounting/config"
	"Accounting/internal/database"
)

func main() {
	config.InitConfig()
	database.Connect(config.Cfg.DbURL)
	if config.Cfg.AutoMigrate {
		database.Migrate()
	}

}
