package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Nps struct {
	*Queries
	db *sql.DB
}

func NewNps(db *sql.DB) *Nps {
	return &Nps{
		db:      db,
		Queries: New(db),
	}
}

func (nps *Nps) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := nps.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("Error %v failed to rollback: %v", err, rbError)
		}
		return err
	}

	return tx.Commit()
}

type NpsCreationTxParams struct {
	CreateRecipientParams    CreateRecipientParams    `json:"CreateRecipient"`
}

type NpsCreationTxResult struct {
	Recipient    Recipient    `json:"recipient"`
}

func (nps *Nps) NpsCreationTx(ctx context.Context, arg NpsCreationTxParams) (NpsCreationTxResult, error) {
	var result NpsCreationTxResult

	err := nps.execTx(ctx, func(q *Queries) error {
		var err error // Declare err variable here
		result.Recipient, err = q.CreateRecipient(ctx, CreateRecipientParams(arg.CreateRecipientParams))
		if err != nil {
            return err
        }
		return nil
	})

	if err != nil {
        return result, err
    }
	return result, nil
}
