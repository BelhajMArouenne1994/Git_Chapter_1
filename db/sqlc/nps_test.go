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

	errs := make(chan error)
	results := make(chan NpsCreationTxResult)

	// Create recipients
	for i:=0 ; i < 10; i++ {
		argRecipient := CreateRecipientParams{
			OcdMasterId: uuid.New().String(),
			Username: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
			Role: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
			CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		}
		
		go func ()  {
			result, err := nps.NpsCreationTx(context.Background(), NpsCreationTxParams {
				CreateRecipientParams: argRecipient,
			}) 

			errs <- err
			results <- result
		} ()
	}

	for i:=0 ; i < 10; i++ {
		err := <- errs
		require.NoError(t, err)

		results := <- results
		require.NotEmpty(t, results.Recipient)
	}

	//Create saleshistory

}


