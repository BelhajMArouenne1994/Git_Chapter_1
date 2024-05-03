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
			return fmt.Errorf(`error %v failed to rollback: %v;`, err, rbError)
		}
		return err
	}

	return tx.Commit()
}

type NpsRecipientCreationTxParams struct {
	CreateRecipientParams CreateRecipientParams `json:"createRecipient"`
}

type NpsRecipientCreationTxResult struct {
	Recipient Recipient `json:"recipient"`
}

func (nps *Nps) NpsRecipientCreationTx(ctx context.Context, arg NpsRecipientCreationTxParams) (NpsRecipientCreationTxResult, error) {
	var result NpsRecipientCreationTxResult

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

type NpsCreationSalesHistoryTxParams struct {
	CreateSalesHistoryParams CreateSalesHistoryParams `json:"createSalesHistory"`
}

type NpsCreationSalesHistoryTxResult struct {
	SalesHistory SalesHistory `json:"salesHistory"`
}

func (nps *Nps) NpsCreationSalesHistoryTx(ctx context.Context, arg NpsCreationSalesHistoryTxParams) (NpsCreationSalesHistoryTxResult, error) {
	var result NpsCreationSalesHistoryTxResult

	err := nps.execTx(ctx, func(q *Queries) error {
		var err error // Declare err variable here
		result.SalesHistory, err = q.CreateSalesHistory(ctx, CreateSalesHistoryParams(arg.CreateSalesHistoryParams))
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
