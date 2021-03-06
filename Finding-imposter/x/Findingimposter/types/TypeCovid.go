package types

import (
	"time"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Covid struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  CovidID string `json:"covidID" yaml:"covidID"`
  CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
  Status string `json:"status" yaml:"status"`
  PubKey []string `json:"pubKey" yaml:"pubKey"`
}