package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Doctor struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Address string `json:"address" yaml:"address"`
  IsDoctor string `json:"isDoctor" yaml:"isDoctor"`
}