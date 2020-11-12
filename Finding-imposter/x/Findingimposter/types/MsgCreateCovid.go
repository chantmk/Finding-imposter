package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateCovid{}

type MsgCreateCovid struct {
  ID      string
  Creator sdk.AccAddress 	`json:"creator" yaml:"creator"`
  Status 		string		`json:"status" yaml:"status"`
  Created_at	string   	`json:"created_at" yaml:"create_at"`
  Pub_key []string 			`json:"pub_key" yaml:"pub_key"`
}

func NewMsgCreateCovid(creator sdk.AccAddress, status string, created_at string, pub_key []string)  MsgCreateCovid{
  return MsgCreateCovid{
    ID: uuid.New().String(),
	Creator: creator,
	Status: status,
	Created_at: created_at,
    Pub_key: pub_key,
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