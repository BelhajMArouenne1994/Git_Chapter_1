package db

import (
    "context"
    "database/sql"
    "time"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/google/uuid"
)

func TestCreateSalesHistory(t *testing.T) {
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
		OcdContactMasterId:      sql.NullString{String: "", Valid: false},
		OcdContactVersionId:     sql.NullString{String: "VID123456", Valid: true},
		OcdStoreId:              sql.NullString{String: "SID123456", Valid: true},
	}

	createdSalesHistory, error := testQueries.CreateSalesHistory(context.Background(), salesHistoryParams);
	require.NoError(t, error)
	require.NotEmpty(t, createdSalesHistory)
}