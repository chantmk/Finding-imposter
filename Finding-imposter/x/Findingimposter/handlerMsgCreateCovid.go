package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateCovid(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCovid) (*sdk.Result, error) {
	var covid = types.Covid{
		Creator: msg.Creator,
		ID:      msg.ID,
		Status: msg.Status,
		Created_at: msg.Created_at,
		Pub_key: msg.Pub_key,
	}
	k.CreateCovid(ctx, covid)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
