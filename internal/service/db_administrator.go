package service

import (
	"Accounting/internal/database"
	"Accounting/internal/entities"
	"Accounting/internal/utils"
	"errors"
	"time"
)

var statusTransitions = map[string][]string{
	"NEW":    {"AUTH"},
	"AUTH":   {"CHARGE", "CANCEL"},
	"CHARGE": {"REFUND"},
}

func CreateTransaction(terminalID int64, orderID string, amount float64, code string, message *string) (*entities.Transaction, error) {
	tx := &entities.Transaction{
		ID:            utils.GenerateTransactionID(),
		TerminalID:    terminalID,
		OrderID:       orderID,
		Amount:        amount,
		Status:        "NEW",
		CreatedAt:     time.Now(),
		StatusChanged: time.Now(),
		Code:          code,
		Message:       message,
	}
	if err := database.DB.Create(tx).Error; err != nil {
		return nil, err
	}
	return tx, nil
}

func ChangeTransactionStatus(transactionID int64, newStatus string) error {
	var tx entities.Transaction
	if err := database.DB.First(&tx, "id = ?", transactionID).Error; err != nil {
		return err
	}
	validNext := statusTransitions[tx.Status]
	allowed := false
	for _, next := range validNext {
		if newStatus == next {
			allowed = true
			break
		}
	}
	if !allowed {
		return errors.New("invalid status transition")
	}
	tx.Status = newStatus
	tx.StatusChanged = time.Now()
	return database.DB.Save(&tx).Error
}

func CancelTransaction(id int64) error {
	return ChangeTransactionStatus(id, "CANCEL")
}

func ChargeTransaction(id int64) error {
	return ChangeTransactionStatus(id, "CHARGE")
}
func RefundTransaction(id int64) error {
	return ChangeTransactionStatus(id, "REFUND")
}
