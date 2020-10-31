package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateLog(ctx sdk.Context, Log types.Log) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.LogPrefix + Log.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(Log)
	store.Set(key, value)
}

func listLog(ctx sdk.Context, k Keeper) ([]byte, error) {
  var LogList []types.Log
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.LogPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var Log types.Log
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &Log)
    LogList = append(LogList, Log)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, LogList)
  return res, nil
}