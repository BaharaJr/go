package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	sender := CreateRandomAccount(t)
	receiver := CreateRandomAccount(t)

	//run n concurrent transactions
	n := 5
	amount := int64(10)

	errors := make(chan error)
	results := make(chan TransferResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), CreateTransferParams{
				Sender:   sender.ID,
				Receiver: receiver.ID,
				Amount:   amount,
			})
			errors <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		error := <-errors
		require.NoError(t, error)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.Sender, sender.ID)
		require.Equal(t, transfer.Receiver, receiver.ID)
		require.Equal(t, transfer.Amount, amount)
		require.NotZero(t, transfer.Created)

		_, err := store.GetTransfer(context.Background(), transfer.ID)

		require.NoError(t, err)

		//check entries
		senderEntry := result.SenderEntry

		require.NotEmpty(t, senderEntry)
		require.Equal(t, senderEntry.Amount, -amount)
		require.Equal(t, senderEntry.Account, sender.ID)
		require.NotZero(t, senderEntry.Created)

		_, err = store.GetEntryId(context.Background(), senderEntry.ID)

		require.NoError(t, err)

		receiverEntry := result.ReceiverEntry

		require.NotEmpty(t, receiverEntry)
		require.Equal(t, receiverEntry.Amount, amount)
		require.Equal(t, receiverEntry.Account, receiver.ID)
		require.NotZero(t, receiverEntry.Created)

		_, err = store.GetEntryId(context.Background(), receiverEntry.ID)

		require.NoError(t, err)

		//! TODO: check account balances
	}

}
