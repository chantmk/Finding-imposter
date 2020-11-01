package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Quarantine struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  User_id string `json:"user_id" yaml:"user_id"`
  Start_at string `json:"start_at" yaml:"start_at"`
  End_at string `json:"end_at" yaml:"end_at"`
}