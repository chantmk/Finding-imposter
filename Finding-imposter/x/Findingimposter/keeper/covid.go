package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateCovid(ctx sdk.Context, covid types.Covid) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.CovidPrefix + covid.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(covid)
	store.Set(key, value)
}

func listCovid(ctx sdk.Context, k Keeper) ([]byte, error) {
  var covidList []types.Covid
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.CovidPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var covid types.Covid
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &covid)
    covidList = append(covidList, covid)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, covidList)
  return res, nil
}

func listPendingCovid(ctx sdk.Context, k Keeper)([]byte, error) {
	var covidList []types.Covid
	PENDING := "PENDING"
	// var visitedCovidID map[string]bool
	visitedCovidID := make(map[string]bool)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.CovidPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var covid types.Covid
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &covid)
		if covid.Status != PENDING {
			visitedCovidID[covid.CovidID] = true
		} else if (covid.Status == PENDING && !visitedCovidID[covid.CovidID]){
			covidList = append(covidList, covid)
		}
	}	
		res := codec.MustMarshalJSONIndent(k.cdc, covidList)
  		return res, nil
}
