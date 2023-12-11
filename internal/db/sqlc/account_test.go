package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/zomgra/simplybank/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomMoney(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
func getAccount(ID int64) (Account, error) {
	account, err := testQueries.GetAccount(context.Background(), ID)

	return account, err
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
func TestGetAccount(t *testing.T) {
	// create account1
	account1 := createRandomAccount(t)
	account2, err := getAccount(account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.ID, account2.ID)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	newBalance := util.RandomMoney()
	arg := UpdateAccountParams{account1.ID, newBalance}
	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	account2, err := getAccount(account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	require.Equal(t, account2.Balance, newBalance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
}
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := getAccount(account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := GetAccountsParams{Limit: 5, Offset: 5}
	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)

	}
}
