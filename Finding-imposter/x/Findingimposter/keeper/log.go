package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateLog(ctx sdk.Context, log types.Log) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.LogPrefix + log.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(log)
	store.Set(key, value)
}

func listLog(ctx sdk.Context, k Keeper) ([]byte, error) {
  var logList []types.Log
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.LogPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var log types.Log
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &log)
    logList = append(logList, log)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, logList)
  return res, nil
}