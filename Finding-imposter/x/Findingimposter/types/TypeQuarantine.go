package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Quarantine struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  UserAddress sdk.AccAddress `json: "userAddress" yaml:"userAddress"`
  StartAt string `json:"startAt" yaml:"startAt"`
  EndAt string `json:"endAt" yaml:"endAt"`
}