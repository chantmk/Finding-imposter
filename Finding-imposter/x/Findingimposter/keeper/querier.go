package keeper

import (
  // this line is used by starport scaffolding
	"github.com/Finding-imposter/x/Findingimposter/types"
		
	
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for Findingimposter clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListQuarantine:
			return listQuarantine(ctx, k)
		case types.QueryListCovid:
			return listCovid(ctx, k)
		case types.QueryListLog:
			return listLog(ctx, k)
		case types.QueryListDoctor:
			return listDoctor(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown Findingimposter query endpoint")
		}
	}
}