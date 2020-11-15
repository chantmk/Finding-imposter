package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateCovid{}

type MsgCreateCovid struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  CovidID string `json:"covidID" yaml:"covidID"`
  CreatedAt string `json:"createdAt" yaml:"createdAt"`
  Status string `json:"status" yaml:"status"`
}

func NewMsgCreateCovid(creator sdk.AccAddress, covidID string, createdAt string, status string) MsgCreateCovid {
  return MsgCreateCovid{
    ID: uuid.New().String(),
		Creator: creator,
    CovidID: covidID,
    CreatedAt: createdAt,
    Status: status,
	}
}

func (msg MsgCreateCovid) Route() string {
  return RouterKey
}

func (msg MsgCreateCovid) Type() string {
  return "CreateCovid"
}

func (msg MsgCreateCovid) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateCovid) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateCovid) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}