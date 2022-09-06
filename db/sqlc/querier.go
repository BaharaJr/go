// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	GenericSearch(ctx context.Context, arg GenericSearchParams) ([]Account, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (Account, error)
	GetAccountByOwner(ctx context.Context, owner string) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
