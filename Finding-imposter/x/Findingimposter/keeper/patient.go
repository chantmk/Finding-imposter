package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreatePatient(ctx sdk.Context, patient types.Patient) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PatientPrefix + patient.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(patient)
	store.Set(key, value)
}

func listPatient(ctx sdk.Context, k Keeper) ([]byte, error) {
  var patientList []types.Patient
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.PatientPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var patient types.Patient
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &patient)
    patientList = append(patientList, patient)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, patientList)
  return res, nil
}