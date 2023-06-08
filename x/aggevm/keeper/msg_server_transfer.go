package keeper

import (
	"context"
	"errors"

	"aggevm/x/aggevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.New("from address error")
	}
	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, errors.New("to address error")
	}
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errors.New("invalid amount")
	}
	err = k.bankKeeper.SendCoins(ctx, from, to, sdk.NewCoins(sdk.NewCoin("uagg", amount)))
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferResponse{}, nil
}
