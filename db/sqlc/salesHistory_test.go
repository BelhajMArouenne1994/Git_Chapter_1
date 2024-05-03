package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateSalesHistory(t *testing.T) {
	salesHistoryParams := CreateSalesHistoryParams{
		Brand:                   "Acme Corp",
		Country:                 "US",
		OcdTicketId:             uuid.New().String(),
		TechnicalCreationDate:   time.Now(),
		TechnicalLastUpdateDate: time.Now(),
		Source:                  "Online",
		SourceName:              "Main Website",
		SourceChannel:           "Web",
		SourcePersonId:          "12345",
		SourceTicketNumber:      "10001",
		SourceStoreType:         "Retail",
		SourceStatusOrder:       "Completed",
		OcdContactMasterId:      "",
		OcdContactVersionId:     "VID123456",
		OcdStoreId:              "SID123456",
	}

	createdSalesHistory, error := testQueries.CreateSalesHistory(context.Background(), salesHistoryParams)
	require.NoError(t, error)
	require.NotEmpty(t, createdSalesHistory)
}
