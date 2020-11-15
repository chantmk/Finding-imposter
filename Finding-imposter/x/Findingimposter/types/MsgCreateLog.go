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
  LogID string `json:"logID" yaml:"logID"`
  PlaceID string `json:"placeID" yaml:"placeID"`
  CreatedAt string `json:"createdAt" yaml:"createdAt"`
  Action string `json:"action" yaml:"action"`
}

func NewMsgCreateLog(creator sdk.AccAddress, logID string, placeID string, createdAt string, action string) MsgCreateLog {
  return MsgCreateLog{
    ID: uuid.New().String(),
		Creator: creator,
    LogID: logID,
    PlaceID: placeID,
    CreatedAt: createdAt,
    Action: action,
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