package keeper

import (
	"context"

	"aggevm/x/aggevm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EvmAll(goCtx context.Context, req *types.QueryAllEvmRequest) (*types.QueryAllEvmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var evms []types.Evm
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	evmStore := prefix.NewStore(store, types.KeyPrefix(types.EvmKeyPrefix))

	pageRes, err := query.Paginate(evmStore, req.Pagination, func(key []byte, value []byte) error {
		var evm types.Evm
		if err := k.cdc.Unmarshal(value, &evm); err != nil {
			return err
		}

		evms = append(evms, evm)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEvmResponse{Evm: evms, Pagination: pageRes}, nil
}

func (k Keeper) Evm(goCtx context.Context, req *types.QueryGetEvmRequest) (*types.QueryGetEvmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEvm(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEvmResponse{Evm: val}, nil
}
