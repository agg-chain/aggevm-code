package keeper

import (
	"aggevm/x/aggevm/types"
)

var _ types.QueryServer = Keeper{}
