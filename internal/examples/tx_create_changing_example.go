package examples

import (
	"Accounting/internal/database"
	"Accounting/internal/entities"
	"Accounting/internal/service"
	"Accounting/internal/utils"
	"log"
)

/*
These methods demonstrate the principle of creating and changing the status of transactions
in the database.
These examples are only an example of the use of functionality and do not claim the title of standard.
*/

func CreateTxExample() *entities.Transaction {
	var generatedTerminal = entities.Terminal{
		ID:           utils.GenerateID(),
		ClientID:     utils.GenerateString(5),
		ClientSecret: utils.GenerateString(10),
		UUID:         utils.GenerateUUID(),
	}
	database.DB.Create(&generatedTerminal)
	log.Println("Terminal generate successfully:", generatedTerminal)
	var generatedTransaction, err = service.CreateTransaction(generatedTerminal, utils.GenerateString(5),
		utils.GenerateRandomFloat(10, 50000), "NEW", nil)
	if err != nil {
		log.Fatalf("Transaction error: %v", err)
	} else {
		log.Println(generatedTransaction)
	}
	return generatedTransaction
}
