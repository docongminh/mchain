package mchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mchain/testutil/keeper"
	"mchain/testutil/nullify"
	"mchain/x/mchain"
	"mchain/x/mchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MchainKeeper(t)
	mchain.InitGenesis(ctx, *k, genesisState)
	got := mchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
