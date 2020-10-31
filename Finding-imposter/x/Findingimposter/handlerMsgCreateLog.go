package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateLog(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLog) (*sdk.Result, error) {
	var Log = types.Log{
		Creator: msg.Creator,
		ID:      msg.ID,
    Id: msg.Id,
    Place_id: msg.Place_id,
    User_id: msg.User_id,
    Checkin_at: msg.Checkin_at,
    Checkout_at: msg.Checkout_at,
	}
	k.CreateLog(ctx, Log)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
