package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateQuarantine(ctx sdk.Context, quarantine types.Quarantine) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.QuarantinePrefix + quarantine.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(quarantine)
	store.Set(key, value)
}

func listQuarantine(ctx sdk.Context, k Keeper) ([]byte, error) {
  var quarantineList []types.Quarantine
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.QuarantinePrefix))
  for ; iterator.Valid(); iterator.Next() {
    var quarantine types.Quarantine
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &quarantine)
    quarantineList = append(quarantineList, quarantine)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, quarantineList)
  return res, nil
}