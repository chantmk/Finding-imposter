package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
	"github.com/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateQuarantine(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateQuarantine) (*sdk.Result, error) {
	var quarantine = types.Quarantine{
		Creator: msg.Creator,
		ID:      msg.ID,
    StartAt: msg.StartAt,
    EndAt: msg.EndAt,
	}
	k.CreateQuarantine(ctx, quarantine)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
