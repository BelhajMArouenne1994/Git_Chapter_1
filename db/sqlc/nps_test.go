package db

import (
	"context"
	"database/sql"
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
			Username: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
			Role: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
			CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
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
				Brand:                   sql.NullString{String: "Acme Corp", Valid: true},
				Country:                 sql.NullString{String: "US", Valid: true},
				OcdTicketId:             uuid.New().String(),
				TechnicalCreationDate:   sql.NullTime{Time: time.Now(), Valid: true},
				TechnicalLastUpdateDate: sql.NullTime{Time: time.Now(), Valid: true},
				Source:                  sql.NullString{String: "Online", Valid: true},
				SourceName:              sql.NullString{String: "Main Website", Valid: true},
				SourceChannel:           sql.NullString{String: "Web", Valid: true},
				SourcePersonId:          sql.NullString{String: "12345", Valid: true},
				SourceTicketNumber:      sql.NullString{String: "10001", Valid: true},
				SourceStoreType:         sql.NullString{String: "Retail", Valid: true},
				SourceStatusOrder:       sql.NullString{String: "Completed", Valid: true},
					OcdContactMasterId:      sql.NullString{String: results.Recipient.OcdMasterId, Valid: true},
				OcdContactVersionId:     sql.NullString{String: "VID123456", Valid: true},
				OcdStoreId:              sql.NullString{String: "SID123456", Valid: true},
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