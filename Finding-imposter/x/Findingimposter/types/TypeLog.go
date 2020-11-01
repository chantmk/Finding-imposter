package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Log struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Place_id string `json:"place_id" yaml:"place_id"`
  Check_in_at string `json:"check_in_at" yaml:"check_in_at"`
  Check_out_at string `json:"check_out_at" yaml:"check_out_at"`
}