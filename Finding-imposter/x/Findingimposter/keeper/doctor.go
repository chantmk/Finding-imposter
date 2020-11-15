package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateDoctor(ctx sdk.Context, doctor types.Doctor) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.DoctorPrefix + doctor.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(doctor)
	store.Set(key, value)
}

func listDoctor(ctx sdk.Context, k Keeper) ([]byte, error) {
  var doctorList []types.Doctor
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.DoctorPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var doctor types.Doctor
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &doctor)
    doctorList = append(doctorList, doctor)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, doctorList)
  return res, nil
}