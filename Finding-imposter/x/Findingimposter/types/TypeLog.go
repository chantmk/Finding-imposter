package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Log struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Id string `json:"id" yaml:"id"`
  Place_id string `json:"place_id" yaml:"place_id"`
  User_id string `json:"user_id" yaml:"user_id"`
  Checkin_at string `json:"checkin_at" yaml:"checkin_at"`
  Checkout_at string `json:"checkout_at" yaml:"checkout_at"`
}