package db

import (
	"context"
	"testing"

	"github.com/go-api-payline/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomStore(t *testing.T) Stores {

	arg := CreateStoreParams{
		StoreAccessID: 1,
		Name: util.CreateRandomStore(),
	}

	store, err := testQueries.CreateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store)


	return store
}

func TestCreateRandomStore(t *testing.T) {
	CreateRandomStore(t)
}
