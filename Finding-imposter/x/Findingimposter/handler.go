package Findingimposter

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/keeper"
	"github.com/Finding-imposter/x/Findingimposter/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding
		case types.MsgCreateCovid:
			return handleMsgCreateCovid(ctx, k, msg)
		case types.MsgCreateLog:
			return handleMsgCreateLog(ctx, k, msg)
		case types.MsgCreateDoctor:
			return handleMsgCreateDoctor(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
