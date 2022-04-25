package db

import (
	"context"
	"testing"

	"github.com/Emmrys-Jay/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomBalance(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.Amount)
	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.CreatedAt, transfer2.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	account3 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		arg := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account3.ID,
			Amount:        util.RandomBalance(),
		}
		testQueries.CreateTransfer(context.Background(), arg)
	}

	for i := 0; i < 5; i++ {
		arg := CreateTransferParams{
			FromAccountID: account3.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomBalance(),
		}
		testQueries.CreateTransfer(context.Background(), arg)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        2,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for i := 0; i < 5; i++ {
		require.NotEmpty(t, transfers[i].ID)
		require.NotEmpty(t, transfers[i].FromAccountID)
		require.NotEmpty(t, transfers[i].ToAccountID)
		require.NotEmpty(t, transfers[i].Amount)
		require.NotEmpty(t, transfers[i].CreatedAt)
	}
}
