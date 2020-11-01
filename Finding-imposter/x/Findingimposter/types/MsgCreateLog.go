package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateLog{}

type MsgCreateLog struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Place_id string `json:"place_id" yaml:"place_id"`
  Check_in_at string `json:"check_in_at" yaml:"check_in_at"`
  Check_out_at string `json:"check_out_at" yaml:"check_out_at"`
}

func NewMsgCreateLog(creator sdk.AccAddress, place_id string, check_in_at string, check_out_at string) MsgCreateLog {
  return MsgCreateLog{
    ID: uuid.New().String(),
		Creator: creator,
    Place_id: place_id,
    Check_in_at: check_in_at,
    Check_out_at: check_out_at,
	}
}

func (msg MsgCreateLog) Route() string {
  return RouterKey
}

func (msg MsgCreateLog) Type() string {
  return "CreateLog"
}

func (msg MsgCreateLog) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateLog) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateLog) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}