package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
	"github.com/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateCovid(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCovid) (*sdk.Result, error) {
	var covid = types.Covid{
		Creator: msg.Creator,
		ID:      msg.ID,
    CovidID: msg.CovidID,
    CreatedAt: msg.CreatedAt,
    Status: msg.Status,
	}
	k.CreateCovid(ctx, covid)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
