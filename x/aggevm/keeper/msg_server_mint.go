package keeper

import (
	"context"
	"errors"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"aggevm/x/aggevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errors.New("invalid amount")
	}
	err := k.bankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin("uagg", amount)))
	if err != nil {
		return nil, err
	}
	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, errors.New("to address error")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, to, sdk.NewCoins(sdk.NewCoin("uagg", amount)))
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
