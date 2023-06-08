package aggevm_test

import (
	"testing"

	keepertest "aggevm/testutil/keeper"
	"aggevm/testutil/nullify"
	"aggevm/x/aggevm"
	"aggevm/x/aggevm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EvmList: []types.Evm{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AggevmKeeper(t)
	aggevm.InitGenesis(ctx, *k, genesisState)
	got := aggevm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EvmList, got.EvmList)
	// this line is used by starport scaffolding # genesis/test/assert
}
