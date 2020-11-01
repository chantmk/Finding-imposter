package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Patient struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Status string `json:"status" yaml:"status"`
  User_id string `json:"user_id" yaml:"user_id"`
}