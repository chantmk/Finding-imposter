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
  Id string `json:"id" yaml:"id"`
  Place_id string `json:"place_id" yaml:"place_id"`
  User_id string `json:"user_id" yaml:"user_id"`
  Checkin_at string `json:"checkin_at" yaml:"checkin_at"`
  Checkout_at string `json:"checkout_at" yaml:"checkout_at"`
}

func NewMsgCreateLog(creator sdk.AccAddress, id string, place_id string, user_id string, checkin_at string, checkout_at string) MsgCreateLog {
  return MsgCreateLog{
    ID: uuid.New().String(),
		Creator: creator,
    Id: id,
    Place_id: place_id,
    User_id: user_id,
    Checkin_at: checkin_at,
    Checkout_at: checkout_at,
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