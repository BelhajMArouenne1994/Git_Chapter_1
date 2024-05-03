// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createRecipientStmt, err = db.PrepareContext(ctx, createRecipient); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRecipient: %w", err)
	}
	if q.createSalesHistoryStmt, err = db.PrepareContext(ctx, createSalesHistory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSalesHistory: %w", err)
	}
	if q.createSurveyUrlStmt, err = db.PrepareContext(ctx, createSurveyUrl); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSurveyUrl: %w", err)
	}
	if q.deleteRecipientStmt, err = db.PrepareContext(ctx, deleteRecipient); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRecipient: %w", err)
	}
	if q.deleteSalesHistoryStmt, err = db.PrepareContext(ctx, deleteSalesHistory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSalesHistory: %w", err)
	}
	if q.deleteSurveyUrlStmt, err = db.PrepareContext(ctx, deleteSurveyUrl); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSurveyUrl: %w", err)
	}
	if q.findSalesHistoryByBrandStmt, err = db.PrepareContext(ctx, findSalesHistoryByBrand); err != nil {
		return nil, fmt.Errorf("error preparing query FindSalesHistoryByBrand: %w", err)
	}
	if q.findSalesHistoryByCountryStmt, err = db.PrepareContext(ctx, findSalesHistoryByCountry); err != nil {
		return nil, fmt.Errorf("error preparing query FindSalesHistoryByCountry: %w", err)
	}
	if q.findSurveyUrlsByBrandAndCountryStmt, err = db.PrepareContext(ctx, findSurveyUrlsByBrandAndCountry); err != nil {
		return nil, fmt.Errorf("error preparing query FindSurveyUrlsByBrandAndCountry: %w", err)
	}
	if q.getRecipientStmt, err = db.PrepareContext(ctx, getRecipient); err != nil {
		return nil, fmt.Errorf("error preparing query GetRecipient: %w", err)
	}
	if q.getSalesHistoryByIdStmt, err = db.PrepareContext(ctx, getSalesHistoryById); err != nil {
		return nil, fmt.Errorf("error preparing query GetSalesHistoryById: %w", err)
	}
	if q.getSurveyUrlByIdStmt, err = db.PrepareContext(ctx, getSurveyUrlById); err != nil {
		return nil, fmt.Errorf("error preparing query GetSurveyUrlById: %w", err)
	}
	if q.linkSurveyUrlToRecipientStmt, err = db.PrepareContext(ctx, linkSurveyUrlToRecipient); err != nil {
		return nil, fmt.Errorf("error preparing query LinkSurveyUrlToRecipient: %w", err)
	}
	if q.listNewestSalesHistoryStmt, err = db.PrepareContext(ctx, listNewestSalesHistory); err != nil {
		return nil, fmt.Errorf("error preparing query ListNewestSalesHistory: %w", err)
	}
	if q.listNewestSurveyUrlsStmt, err = db.PrepareContext(ctx, listNewestSurveyUrls); err != nil {
		return nil, fmt.Errorf("error preparing query ListNewestSurveyUrls: %w", err)
	}
	if q.listNewestSurveyUrlsByBrandAndCountryStmt, err = db.PrepareContext(ctx, listNewestSurveyUrlsByBrandAndCountry); err != nil {
		return nil, fmt.Errorf("error preparing query ListNewestSurveyUrlsByBrandAndCountry: %w", err)
	}
	if q.listRecipientsStmt, err = db.PrepareContext(ctx, listRecipients); err != nil {
		return nil, fmt.Errorf("error preparing query ListRecipients: %w", err)
	}
	if q.updateRecipientStmt, err = db.PrepareContext(ctx, updateRecipient); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateRecipient: %w", err)
	}
	if q.updateSalesHistoryStmt, err = db.PrepareContext(ctx, updateSalesHistory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSalesHistory: %w", err)
	}
	if q.updateSurveyUrlStmt, err = db.PrepareContext(ctx, updateSurveyUrl); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSurveyUrl: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createRecipientStmt != nil {
		if cerr := q.createRecipientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRecipientStmt: %w", cerr)
		}
	}
	if q.createSalesHistoryStmt != nil {
		if cerr := q.createSalesHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSalesHistoryStmt: %w", cerr)
		}
	}
	if q.createSurveyUrlStmt != nil {
		if cerr := q.createSurveyUrlStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSurveyUrlStmt: %w", cerr)
		}
	}
	if q.deleteRecipientStmt != nil {
		if cerr := q.deleteRecipientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRecipientStmt: %w", cerr)
		}
	}
	if q.deleteSalesHistoryStmt != nil {
		if cerr := q.deleteSalesHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSalesHistoryStmt: %w", cerr)
		}
	}
	if q.deleteSurveyUrlStmt != nil {
		if cerr := q.deleteSurveyUrlStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSurveyUrlStmt: %w", cerr)
		}
	}
	if q.findSalesHistoryByBrandStmt != nil {
		if cerr := q.findSalesHistoryByBrandStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findSalesHistoryByBrandStmt: %w", cerr)
		}
	}
	if q.findSalesHistoryByCountryStmt != nil {
		if cerr := q.findSalesHistoryByCountryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findSalesHistoryByCountryStmt: %w", cerr)
		}
	}
	if q.findSurveyUrlsByBrandAndCountryStmt != nil {
		if cerr := q.findSurveyUrlsByBrandAndCountryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findSurveyUrlsByBrandAndCountryStmt: %w", cerr)
		}
	}
	if q.getRecipientStmt != nil {
		if cerr := q.getRecipientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRecipientStmt: %w", cerr)
		}
	}
	if q.getSalesHistoryByIdStmt != nil {
		if cerr := q.getSalesHistoryByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSalesHistoryByIdStmt: %w", cerr)
		}
	}
	if q.getSurveyUrlByIdStmt != nil {
		if cerr := q.getSurveyUrlByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSurveyUrlByIdStmt: %w", cerr)
		}
	}
	if q.linkSurveyUrlToRecipientStmt != nil {
		if cerr := q.linkSurveyUrlToRecipientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing linkSurveyUrlToRecipientStmt: %w", cerr)
		}
	}
	if q.listNewestSalesHistoryStmt != nil {
		if cerr := q.listNewestSalesHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNewestSalesHistoryStmt: %w", cerr)
		}
	}
	if q.listNewestSurveyUrlsStmt != nil {
		if cerr := q.listNewestSurveyUrlsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNewestSurveyUrlsStmt: %w", cerr)
		}
	}
	if q.listNewestSurveyUrlsByBrandAndCountryStmt != nil {
		if cerr := q.listNewestSurveyUrlsByBrandAndCountryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNewestSurveyUrlsByBrandAndCountryStmt: %w", cerr)
		}
	}
	if q.listRecipientsStmt != nil {
		if cerr := q.listRecipientsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRecipientsStmt: %w", cerr)
		}
	}
	if q.updateRecipientStmt != nil {
		if cerr := q.updateRecipientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateRecipientStmt: %w", cerr)
		}
	}
	if q.updateSalesHistoryStmt != nil {
		if cerr := q.updateSalesHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSalesHistoryStmt: %w", cerr)
		}
	}
	if q.updateSurveyUrlStmt != nil {
		if cerr := q.updateSurveyUrlStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSurveyUrlStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                        DBTX
	tx                                        *sql.Tx
	createRecipientStmt                       *sql.Stmt
	createSalesHistoryStmt                    *sql.Stmt
	createSurveyUrlStmt                       *sql.Stmt
	deleteRecipientStmt                       *sql.Stmt
	deleteSalesHistoryStmt                    *sql.Stmt
	deleteSurveyUrlStmt                       *sql.Stmt
	findSalesHistoryByBrandStmt               *sql.Stmt
	findSalesHistoryByCountryStmt             *sql.Stmt
	findSurveyUrlsByBrandAndCountryStmt       *sql.Stmt
	getRecipientStmt                          *sql.Stmt
	getSalesHistoryByIdStmt                   *sql.Stmt
	getSurveyUrlByIdStmt                      *sql.Stmt
	linkSurveyUrlToRecipientStmt              *sql.Stmt
	listNewestSalesHistoryStmt                *sql.Stmt
	listNewestSurveyUrlsStmt                  *sql.Stmt
	listNewestSurveyUrlsByBrandAndCountryStmt *sql.Stmt
	listRecipientsStmt                        *sql.Stmt
	updateRecipientStmt                       *sql.Stmt
	updateSalesHistoryStmt                    *sql.Stmt
	updateSurveyUrlStmt                       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                        tx,
		tx:                                        tx,
		createRecipientStmt:                       q.createRecipientStmt,
		createSalesHistoryStmt:                    q.createSalesHistoryStmt,
		createSurveyUrlStmt:                       q.createSurveyUrlStmt,
		deleteRecipientStmt:                       q.deleteRecipientStmt,
		deleteSalesHistoryStmt:                    q.deleteSalesHistoryStmt,
		deleteSurveyUrlStmt:                       q.deleteSurveyUrlStmt,
		findSalesHistoryByBrandStmt:               q.findSalesHistoryByBrandStmt,
		findSalesHistoryByCountryStmt:             q.findSalesHistoryByCountryStmt,
		findSurveyUrlsByBrandAndCountryStmt:       q.findSurveyUrlsByBrandAndCountryStmt,
		getRecipientStmt:                          q.getRecipientStmt,
		getSalesHistoryByIdStmt:                   q.getSalesHistoryByIdStmt,
		getSurveyUrlByIdStmt:                      q.getSurveyUrlByIdStmt,
		linkSurveyUrlToRecipientStmt:              q.linkSurveyUrlToRecipientStmt,
		listNewestSalesHistoryStmt:                q.listNewestSalesHistoryStmt,
		listNewestSurveyUrlsStmt:                  q.listNewestSurveyUrlsStmt,
		listNewestSurveyUrlsByBrandAndCountryStmt: q.listNewestSurveyUrlsByBrandAndCountryStmt,
		listRecipientsStmt:                        q.listRecipientsStmt,
		updateRecipientStmt:                       q.updateRecipientStmt,
		updateSalesHistoryStmt:                    q.updateSalesHistoryStmt,
		updateSurveyUrlStmt:                       q.updateSurveyUrlStmt,
	}
}
