package db

import (
	"context"
	"testing"

	"github.com/go-api-payline/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomStoreAccess(t *testing.T) StoresAccess {

	access, err := testQueries.CreateStoreAccess(context.Background(), util.RandomStoreAccess())
	require.NoError(t, err)
	require.NotEmpty(t, access)

	// kolom hasil generate pake PascalCase
	require.Contains(t, []string{"menu_qr ", "no_acces", "pos"}, access.Name)
	require.NotZero(t, access.ID)

	return access
}

func TestCreateStoreAccess(t *testing.T) {
	CreateRandomStoreAccess(t)
}
