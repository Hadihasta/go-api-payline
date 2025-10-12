package db

// jika dijalankan run pacage test tidak akan 100% karna di role.sql.go masih ada yang belum di test
import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-api-payline/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRole(t *testing.T) Roles {

	role, err := testQueries.CreateRoles(context.Background(), util.RandomRole())
	require.NoError(t, err)
	require.NotEmpty(t, role)

	// kolom hasil generate pake PascalCase
	require.Contains(t, []string{"customer", "super_admin", "owner"}, role.RoleName)
	require.NotZero(t, role.ID)

	return role
}

func TestCreateRole(t *testing.T) {
	CreateRandomRole(t)
}

func TestGetRoles(t *testing.T) {
	role1 := CreateRandomRole(t)
	role2, err := testQueries.GetRoles(context.Background(), role1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, role2)

	require.Equal(t, role1.ID, role2.ID)
}

func TestUpdateRoles(t *testing.T) {
	role1 := CreateRandomRole(t)

	arg := UpdateRolesParams{
		ID:       role1.ID,
		RoleName: util.RandomRole(),
	}

	role2, err := testQueries.UpdateRoles(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, role2)

	require.Equal(t, role1.ID, role2.ID) // ID should stay the same
	require.Equal(t, arg.RoleName, role2.RoleName)
	// require.WithinDuration(t,role1.Createdat, role2.CreatedAt, time.Second) // kalau ada field created at
}

func TestDeleteRoles(t *testing.T) {
	role1 := CreateRandomRole(t)
	err := testQueries.DeleteRoles(context.Background(), role1.ID)
	require.NoError(t, err)

	role2, err := testQueries.GetRoles(context.Background(), role1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, role2)
}

func TestListRoles(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRole(t)
	}

	arg := ListRolesParams{
		Limit:  5,
		Offset: 5,
	}

	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roles, 5)

	for _, roles := range roles {
		require.NotEmpty(t, roles)
	}
}
