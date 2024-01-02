package db

import (
	"context"
	"testing"
	"time"

	"github.com/mauricio-mds/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	createRandomTransfer(t, a1, a2)
}

func TestGetTransfer(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, a1, a2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, a1, a2)
	}

	arg := ListTransfersParams{
		FromAccountID: a1.ID,
		ToAccountID: a2.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	}
}

func createRandomTransfer(t *testing.T, a1 Account, a2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: a1.ID,
		ToAccountID: a2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err :=  testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	return transfer
}