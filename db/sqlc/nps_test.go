package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)



func TestNpsCreationTx(t *testing.T) {

	nps := NewNps(testDB)

	errsRecipients := make(chan error)
	resultsRecipients := make(chan NpsRecipientCreationTxResult)

	// Create recipients
	for i:=0 ; i < 10; i++ {
		argRecipient := CreateRecipientParams{
			OcdMasterId: uuid.New().String(),
			Username:"TEST_RECIPIENT",
			Role:"TEST_RECIPIENT",
			CreatedAt: time.Now().UTC(),
		}
		go func ()  {
			result, err := nps.NpsRecipientCreationTx(context.Background(), NpsRecipientCreationTxParams {
				CreateRecipientParams: argRecipient,
			}) 

			errsRecipients <- err
			resultsRecipients <- result
		} ()
	}



	errSalesHistory := make(chan error)
	resultsSalesHistory := make(chan NpsCreationSalesHistoryTxResult)
	for i:=0 ; i < 10; i++ {
		err := <- errsRecipients
		require.NoError(t, err)

		results := <- resultsRecipients
		require.NotEmpty(t, results.Recipient)

		//Create saleshistory

		//ocdTicketIds := make([]string, 10) // Slice to store ocdTicketId
		for i:=0 ; i < 10; i++ {
			salesHistoryParams := CreateSalesHistoryParams{
				Brand:                  "Acme Corp",
				Country:                "US",
				OcdTicketId:             uuid.New().String(),
				TechnicalCreationDate:   time.Now(),
				TechnicalLastUpdateDate: time.Now(),
				Source:                 "Online",
				SourceName:             "Main Website",
				SourceChannel:          "Web",
				SourcePersonId:         "12345",
				SourceTicketNumber:     "10001",
				SourceStoreType:        "Retail",
				SourceStatusOrder:      "Completed",
				OcdContactMasterId:     results.Recipient.OcdMasterId,
				OcdContactVersionId:    "VID123456",
				OcdStoreId:             "SID123456",
			}

			go func () {
				result, err := nps.NpsCreationSalesHistoryTx(context.Background(), NpsCreationSalesHistoryTxParams {
					CreateSalesHistoryParams : salesHistoryParams,
				}) 
				errSalesHistory <- err
				resultsSalesHistory <- result
			} ()
		}
	}




	for i:=0 ; i < 10; i++ {
		err := <- errSalesHistory
        require.NoError(t, err)

		result := <- resultsSalesHistory
		require.NotEmpty(t, result.SalesHistory)
	}
}