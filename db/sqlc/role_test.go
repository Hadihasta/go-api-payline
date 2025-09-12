package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRole(t *testing.T) {

	role, err := testQueries.CreateRoles(context.Background(), "customer")
	require.NoError(t, err)
	require.NotEmpty(t, role)

	// kolom hasil generate pake PascalCase
	require.Equal(t, "customer", role.RoleName)
	require.NotZero(t, role.ID)
}
