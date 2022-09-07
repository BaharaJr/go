// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Currency string

const (
	CurrencyUSD  Currency = "USD"
	CurrencyTshs Currency = "Tshs"
	CurrencyEUR  Currency = "EUR"
)

func (e *Currency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currency(s)
	case string:
		*e = Currency(s)
	default:
		return fmt.Errorf("unsupported scan type for Currency: %T", src)
	}
	return nil
}

type NullCurrency struct {
	Currency Currency
	Valid    bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrency) Scan(value interface{}) error {
	if value == nil {
		ns.Currency, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currency.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrency) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Currency, nil
}

type Account struct {
	ID       uuid.UUID      `json:"id"`
	Created  time.Time      `json:"created"`
	Code     sql.NullString `json:"code"`
	Owner    string         `json:"owner"`
	Balance  int64          `json:"balance"`
	Currency string         `json:"currency"`
}

type Entry struct {
	ID      uuid.UUID `json:"id"`
	Amount  int64     `json:"amount"`
	Account uuid.UUID `json:"account"`
	Created time.Time `json:"created"`
}

type Transfer struct {
	ID       uuid.UUID `json:"id"`
	Sender   uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
	Amount   int64     `json:"amount"`
	Created  time.Time `json:"created"`
}
