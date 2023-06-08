package keeper_test

import (
	"context"
	"testing"

	keepertest "aggevm/testutil/keeper"
	"aggevm/x/aggevm/keeper"
	"aggevm/x/aggevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AggevmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
