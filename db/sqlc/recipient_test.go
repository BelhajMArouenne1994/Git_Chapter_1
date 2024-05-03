package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateRecipient(t *testing.T) {
	argRecipient := CreateRecipientParams{
		OcdMasterId: uuid.New().String(),
		Username:    uuid.New().String(),
		Role:        uuid.New().String(),
		CreatedAt:   time.Now().UTC(),
	}

	createdRecipient, error := testQueries.CreateRecipient(context.Background(), argRecipient)
	require.NoError(t, error)
	require.NotEmpty(t, createdRecipient)
}
