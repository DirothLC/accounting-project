package database

import (
	"Accounting/internal/entities"
	"errors"
	"gorm.io/gorm"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&entities.Terminal{}, &entities.TransactionStatus{}, &entities.Transaction{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	seedTransactionStatuses()
}

func seedTransactionStatuses() {
	statuses := []entities.TransactionStatus{
		{Code: "REFUND", Description: "Возврат суммы"},
		{Code: "AUTH", Description: "Блокировка суммы"},
		{Code: "CANCEL", Description: "Разблокировка"},
		{Code: "CHARGE", Description: "Списание"},
		{Code: "VERIFIED", Description: "Проверка карты"},
		{Code: "CANCEL_OLD", Description: "Истёк срок CHARGE"},
		{Code: "FAILED", Description: "Неуспешно"},
		{Code: "FINGERPRINT", Description: "Проверка перед 3D"},
		{Code: "3D", Description: "Ошибка 3D"},
		{Code: "NEW", Description: "Ожидание"},
		{Code: "REJECT", Description: "Отклонено"},
	}
	for _, code := range statuses {
		var existing entities.TransactionStatus
		err := DB.First(&existing, "code = ?", code.Code).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&code)
		}
	}
}
