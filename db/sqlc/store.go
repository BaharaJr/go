package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

// store provides all queries and functions on transactions
type Store struct {
	*Queries
	db *sql.DB
}

// TransferTxParams contains all necessary input parameters to transfer money from one account to another
type TransferTxParams struct {
	Sender   uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
	Amount   int64     `json:"amount"`
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
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// create store execution context
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

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
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferResult, error) {
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
