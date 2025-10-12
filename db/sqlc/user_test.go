package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-api-payline/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) Users {

	arg := CreateUserParams{
		RoleID:      util.RandomInt(1, 9),
		Email:       sql.NullString{String: util.RandomEmail(), Valid: true},
		PhoneNumber: util.RandomPhoneNumber(),
		Name:        util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)


	// kolom hasil generate pake PascalCase
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	require.Equal(t, arg.Name, user.Name)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	// ambil data dari database dan simpan ke user2 berdasarkan user1.name
	user2, err := testQueries.GetUser(context.Background(), user1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
}
