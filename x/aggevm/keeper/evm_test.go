package keeper_test

import (
	"strconv"
	"testing"

	keepertest "aggevm/testutil/keeper"
	"aggevm/testutil/nullify"
	"aggevm/x/aggevm/keeper"
	"aggevm/x/aggevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEvm(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Evm {
	items := make([]types.Evm, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEvm(ctx, items[i])
	}
	return items
}

func TestEvmGet(t *testing.T) {
	keeper, ctx := keepertest.AggevmKeeper(t)
	items := createNEvm(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEvm(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEvmRemove(t *testing.T) {
	keeper, ctx := keepertest.AggevmKeeper(t)
	items := createNEvm(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEvm(ctx,
			item.Index,
		)
		_, found := keeper.GetEvm(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEvmGetAll(t *testing.T) {
	keeper, ctx := keepertest.AggevmKeeper(t)
	items := createNEvm(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEvm(ctx)),
	)
}
