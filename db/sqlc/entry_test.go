package db

import (
	"context"
	"testing"

	"github.com/BaharaJr/go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	arg := CreateEntryParams{
		Amount:  util.RandomMoney(),
		Account: account.ID,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, account.ID, entry.Account)
	return entry
}
func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}
func TestGetEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	entry1, error := testQueries.GetEntryId(context.Background(), entry.ID)
	require.NoError(t, error)
	require.NotEmpty(t, entry1)
	require.Equal(t, entry1.ID, entry.ID)
}
func TestUpdateEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	arg := UpdateEntryParams{
		Amount: util.RandomMoney(),
		ID:     entry.ID,
	}
	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)
	require.Equal(t, updatedEntry.ID, entry.ID)
	require.NotEqual(t, updatedEntry.Amount, entry.Amount)
}
func TestGetEntries(t *testing.T) {
	arg := GetEntriesParams{
		Limit:  3,
		Offset: 0,
	}
	entries, error := testQueries.GetEntries(context.Background(), arg)
	require.NoError(t, error)
	require.NotEmpty(t, entries)
	require.Equal(t, len(entries), 3)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestDeleteEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	error := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, error)

	entry1, error := testQueries.GetEntryId(context.Background(), entry.ID)
	require.Error(t, error)
	require.Empty(t, entry1)
}
