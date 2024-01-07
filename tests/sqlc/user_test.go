package tests

import (
	"backend-master-class/api/request_params"
	"backend-master-class/util"
	"context"
	"reflect"
	"testing"
	"time"

	db "backend-master-class/db/sqlc"

	"github.com/stretchr/testify/require"
)

func TestCreateUp(t *testing.T) {
	createRandomUser(t)
}

func createRandomUser(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username:       util.RandomOwner(),
		FullName:       util.RandomOwner(),
		HashedPassword: "secret",
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	fetchedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)
	// require.Equal(t, createdAccount.ID, fetchedAccount.ID)
	// require.Equal(t, createdAccount.Owner, fetchedAccount.Owner)
	// require.Equal(t, createdAccount.Balance, fetchedAccount.Balance)
	// require.Equal(t, createdAccount.Currency, fetchedAccount.Currency)
	require.WithinDuration(t, createdUser.CreatedAt, fetchedUser.CreatedAt, time.Second)
	require.True(t, reflect.DeepEqual(createdUser, fetchedUser))
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	limit := 5
	req := request_params.ListUsersRequest{
		PageSize: 5,
		PageID:   1,
	}
	fetchedUsers, err := testQueries.ListUser(context.Background(), db.ListUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	require.NoError(t, err)
	require.Len(t, fetchedUsers, limit)

	for _, value := range fetchedUsers {
		require.NotEmpty(t, value)
	}
}
