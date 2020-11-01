package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateQuarantine(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateQuarantine) (*sdk.Result, error) {
	var quarantine = types.Quarantine{
		Creator: msg.Creator,
		ID:      msg.ID,
    User_id: msg.User_id,
    Start_at: msg.Start_at,
    End_at: msg.End_at,
	}
	k.CreateQuarantine(ctx, quarantine)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
