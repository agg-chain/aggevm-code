package keeper

import (
	"aggevm/x/aggevm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEvm set a specific evm in the store from its index
func (k Keeper) SetEvm(ctx sdk.Context, evm types.Evm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EvmKeyPrefix))
	b := k.cdc.MustMarshal(&evm)
	store.Set(types.EvmKey(
		evm.Index,
	), b)
}

// GetEvm returns a evm from its index
func (k Keeper) GetEvm(
	ctx sdk.Context,
	index string,

) (val types.Evm, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EvmKeyPrefix))

	b := store.Get(types.EvmKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEvm removes a evm from the store
func (k Keeper) RemoveEvm(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EvmKeyPrefix))
	store.Delete(types.EvmKey(
		index,
	))
}

// GetAllEvm returns all evm
func (k Keeper) GetAllEvm(ctx sdk.Context) (list []types.Evm) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EvmKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Evm
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
