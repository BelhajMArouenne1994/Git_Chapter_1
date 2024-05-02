package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
    "github.com/google/uuid"
)

func TestCreateRecipient(t *testing.T) {
    argRecipient := CreateRecipientParams{
        OcdMasterId: uuid.New().String(),
        Username: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
        Role: sql.NullString{String: "TEST_RECIPIENT", Valid: true},
        CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
        
    }
    
    createdRecipient, error := testQueries.CreateRecipient(context.Background(), argRecipient);
    require.NoError(t, error)
    require.NotEmpty(t, createdRecipient)
}