package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateQuarantine{}

type MsgCreateQuarantine struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Id string `json:"id" yaml:"id"`
  User_id string `json:"user_id" yaml:"user_id"`
  Start_at string `json:"start_at" yaml:"start_at"`
  End_at string `json:"end_at" yaml:"end_at"`
}

func NewMsgCreateQuarantine(creator sdk.AccAddress, id string, user_id string, start_at string, end_at string) MsgCreateQuarantine {
  return MsgCreateQuarantine{
    ID: uuid.New().String(),
		Creator: creator,
    Id: id,
    User_id: user_id,
    Start_at: start_at,
    End_at: end_at,
	}
}

func (msg MsgCreateQuarantine) Route() string {
  return RouterKey
}

func (msg MsgCreateQuarantine) Type() string {
  return "CreateQuarantine"
}

func (msg MsgCreateQuarantine) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateQuarantine) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateQuarantine) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}