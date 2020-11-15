package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
	"github.com/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateLog(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLog) (*sdk.Result, error) {
	var log = types.Log{
		Creator: msg.Creator,
		ID:      msg.ID,
    LogID: msg.LogID,
    PlaceID: msg.PlaceID,
    CreatedAt: msg.CreatedAt,
    Action: msg.Action,
	}
	k.CreateLog(ctx, log)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
