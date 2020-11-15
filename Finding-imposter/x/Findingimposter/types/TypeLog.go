package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Log struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  LogID string `json:"logID" yaml:"logID"`
  PlaceID string `json:"placeID" yaml:"placeID"`
  CreatedAt string `json:"createdAt" yaml:"createdAt"`
  Action string `json:"action" yaml:"action"`
}