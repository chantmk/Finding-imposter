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
  Id string `json:"id" yaml:"id"`
  User_id string `json:"user_id" yaml:"user_id"`
  Status string `json:"status" yaml:"status"`
}

func NewMsgCreatePatient(creator sdk.AccAddress, id string, user_id string, status string) MsgCreatePatient {
  return MsgCreatePatient{
    ID: uuid.New().String(),
		Creator: creator,
    Id: id,
    User_id: user_id,
    Status: status,
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