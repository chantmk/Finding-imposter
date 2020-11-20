package types

import (
	"time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateQuarantine{}

type MsgCreateQuarantine struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  UserAddress sdk.AccAddress `json:"userAddress" yaml:"userAddress"`
  StartAt time.Time `json:"startAt" yaml:"startAt"`
  EndAt time.Time `json:"endAt" yaml:"endAt"`
}

func NewMsgCreateQuarantine(creator sdk.AccAddress,userAddress sdk.AccAddress, startAt time.Time, endAt time.Time) MsgCreateQuarantine {
  return MsgCreateQuarantine{
    ID: uuid.New().String(),
		Creator: creator,
	UserAddress: userAddress,
    StartAt: startAt,
    EndAt: endAt,
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