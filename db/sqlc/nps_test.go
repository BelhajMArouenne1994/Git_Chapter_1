package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)



func TestNpsCreationTx(t *testing.T) {

	argRecipient := CreateRecipientParams{
		OcdMasterId: uuid.New().String(),
		Username: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
		Role: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	nps := NewNps(testDB)

	result, err := nps.NpsCreationTx(context.Background(), NpsCreationTxParams {
		CreateRecipientParams: argRecipient,
	}) 

	if err!= nil {
        t.Error(err)
    }
	fmt.Println(result)
}


