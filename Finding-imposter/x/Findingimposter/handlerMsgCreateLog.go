package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateLog(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLog) (*sdk.Result, error) {
	var log = types.Log{
		Creator: msg.Creator,
		ID:      msg.ID,
    Place_id: msg.Place_id,
    Check_in_at: msg.Check_in_at,
    Check_out_at: msg.Check_out_at,
	}
	k.CreateLog(ctx, log)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
