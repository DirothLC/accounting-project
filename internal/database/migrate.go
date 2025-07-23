package database

import "Accounting/internal/entities"

func Migrate() {
	DB.AutoMigrate(&entities.Terminal{}, &entities.TransactionStatus{}, &entities.Transaction{})
}
