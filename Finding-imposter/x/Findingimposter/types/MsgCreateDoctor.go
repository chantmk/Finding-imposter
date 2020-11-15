package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateDoctor{}

type MsgCreateDoctor struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Address string `json:"address" yaml:"address"`
  IsDoctor string `json:"isDoctor" yaml:"isDoctor"`
}

func NewMsgCreateDoctor(creator sdk.AccAddress, address string, isDoctor string) MsgCreateDoctor {
  return MsgCreateDoctor{
    ID: uuid.New().String(),
		Creator: creator,
    Address: address,
    IsDoctor: isDoctor,
	}
}

func (msg MsgCreateDoctor) Route() string {
  return RouterKey
}

func (msg MsgCreateDoctor) Type() string {
  return "CreateDoctor"
}

func (msg MsgCreateDoctor) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateDoctor) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateDoctor) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}