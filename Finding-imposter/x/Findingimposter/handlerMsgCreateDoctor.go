package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
	"github.com/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreateDoctor(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateDoctor) (*sdk.Result, error) {
	var doctor = types.Doctor{
		Creator: msg.Creator,
		ID:      msg.ID,
    Address: msg.Address,
    IsDoctor: msg.IsDoctor,
	}
	k.CreateDoctor(ctx, doctor)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
