package Findingimposter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/keeper"
)

func handleMsgCreatePatient(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePatient) (*sdk.Result, error) {
	var patient = types.Patient{
		Creator: msg.Creator,
		ID:      msg.ID,
    Status: msg.Status,
    User_id: msg.User_id,
	}
	k.CreatePatient(ctx, patient)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
