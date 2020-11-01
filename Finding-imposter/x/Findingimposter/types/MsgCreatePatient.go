package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreatePatient{}

type MsgCreatePatient struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Status string `json:"status" yaml:"status"`
  User_id string `json:"user_id" yaml:"user_id"`
}

func NewMsgCreatePatient(creator sdk.AccAddress, status string, user_id string) MsgCreatePatient {
  return MsgCreatePatient{
    ID: uuid.New().String(),
		Creator: creator,
    Status: status,
    User_id: user_id,
	}
}

func (msg MsgCreatePatient) Route() string {
  return RouterKey
}

func (msg MsgCreatePatient) Type() string {
  return "CreatePatient"
}

func (msg MsgCreatePatient) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePatient) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePatient) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}