package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg CreateTransferParams) (TransferResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

//Result of th transfer transaction

type TransferResult struct {
	Transfer      Transfer `json:"transfer"`
	Sender        Account  `json:"sender"`
	Receiver      Account  `json:"receiver"`
	SenderEntry   Entry    `json:"sender_entry"`
	ReceiverEntry Entry    `json:"receiver_entry"`
}

// create new store
// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// create store execution context
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	var err error
	var tx *sql.Tx
	tx, err = store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if roleBackError := tx.Rollback(); roleBackError != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, roleBackError)
		}
		return err
	}
	return tx.Commit()
}

// TransferTx performs a money transfer transaction from one account to another
func (store *SQLStore) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferResult, error) {
	var result TransferResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		createTransferArg := CreateTransferParams{
			Sender:   arg.Sender,
			Receiver: arg.Receiver,
			Amount:   arg.Amount,
		}
		result.Transfer, err = q.CreateTransfer(ctx, createTransferArg)
		if err != nil {
			return err
		}
		result.ReceiverEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			Account: result.Transfer.Receiver,
			Amount:  arg.Amount,
		})
		if err != nil {
			return err
		}
		result.SenderEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			Account: result.Transfer.Sender,
			Amount:  -arg.Amount,
		})
		if err != nil {
			return err
		}

		//TODO UPDATE ACCOUNT BALANCE
		return nil

	})
	return result, err
}
