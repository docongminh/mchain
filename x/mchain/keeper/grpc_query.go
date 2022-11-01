package keeper

import (
	"mchain/x/mchain/types"
)

var _ types.QueryServer = Keeper{}
