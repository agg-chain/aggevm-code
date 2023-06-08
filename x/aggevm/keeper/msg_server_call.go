package keeper

import (
	"context"

	"aggevm/x/aggevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Call(goCtx context.Context, msg *types.MsgCall) (*types.MsgCallResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCallResponse{}, nil
}
