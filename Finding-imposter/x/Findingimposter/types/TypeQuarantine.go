package types

import (
	"time"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Quarantine struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  UserAddress sdk.AccAddress `json: "userAddress" yaml:"userAddress"`
  StartAt time.Time `json:"startAt" yaml:"startAt"`
  EndAt time.Time `json:"endAt" yaml:"endAt"`
}