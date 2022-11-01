package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mchain/testutil/keeper"
	"mchain/x/mchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
