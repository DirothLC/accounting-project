package service

import (
	"Accounting/internal/database"
	"Accounting/internal/entities"
	"Accounting/internal/utils"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

var statusTransitions = map[string][]string{
	"NEW":    {"AUTH"},
	"AUTH":   {"CHARGE", "CANCEL"},
	"CHARGE": {"REFUND"},
}

func CreateTransaction(terminal entities.Terminal, orderID string, amount float64, code string, message *string) (*entities.Transaction, error) {
	tx := &entities.Transaction{
		ID:            utils.GenerateID(),
		TerminalID:    terminal.ID,
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
	log.Println("Changing transaction status:", tx.ID, "to", newStatus)
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

func CheckClientSecret(db *gorm.DB, clientID string, clientSecret string) (bool, error) {
	var terminal entities.Terminal
	err := db.Where("client_id = ?", clientID).First(&terminal).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if terminal.ClientSecret != clientSecret {
		return false, nil
	}
	return true, nil
}
func GetTerminalByClientID(db *gorm.DB, clientID string) (*entities.Terminal, error) {
	var terminal entities.Terminal
	err := db.Where("client_id = ?", clientID).First(&terminal).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &terminal, nil
}
