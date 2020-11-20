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
	checkDoctor := isDoctor(ctx, k, covid.Creator.String())
	if covid.Status == "APPROVE" {
		if checkDoctor {
			createQC(ctx, k, covid)
			store.Set(key, value)
		}
	} else if covid.Status == "REJECTED" {
		if checkDoctor {
			store.Set(key,value)
		}
	}
}

func isDoctor(ctx sdk.Context, k Keeper, user string) (bool){
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.DoctorPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var doctor types.Doctor
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &doctor)
		if doctor.IsDoctor == "True" && doctor.Address == user {
			return true
		} 
	}
	return false
}

func createQC(ctx sdk.Context, k Keeper, covid types.Covid) {
	pub := covid.PubKey
	store := ctx.KVStore(k.storeKey)

	//list log
	var logList []types.Log
	for _, pubkey := range pub {
		iterator := sdk.KVStorePrefixIterator(store, []byte(types.LogPrefix))
		for ; iterator.Valid(); iterator.Next() {
			var log types.Log
			k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &log)
			if pubkey == log.Creator.String() {
				logList = append(logList, log)
			}
		}
	}

	//create quarantine for who that need
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LogPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var log types.Log
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &log)
		// for _, covidLog := range logList {
		// 	if 
		// } 
	}
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
	visitedCovidID := make(map[string]bool)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.CovidPrefix))
	var covid types.Covid
	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &covid)
		if covid.Status != PENDING {
			visitedCovidID[covid.CovidID] = true
		}
	}
	iterator2 := sdk.KVStorePrefixIterator(store, []byte(types.CovidPrefix))
	for ; iterator2.Valid(); iterator2.Next() {
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator2.Key()), &covid)
		if (covid.Status == PENDING && !visitedCovidID[covid.CovidID]){
			covidList = append(covidList, covid)
		}
	}

		res := codec.MustMarshalJSONIndent(k.cdc, covidList)
  		return res, nil
}
