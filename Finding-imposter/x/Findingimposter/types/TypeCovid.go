package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// enum for status
// type Status_type string
// const(
// 	Pending 	Status_type = "pending"
// 	Verified				= "verified"
// 	Rejected				= "rejected"
// )

type Covid struct {
	Creator sdk.AccAddress 	`json:"creator" yaml:"creator"`
	ID      	string      `json:"id" yaml:"id"`
	Status 		string		`json:"status" yaml:"status"`
	Created_at	string   	`json:"created_at" yaml:"created_at"`
  	Pub_key 	[]string 	`json:"pub_key" yaml:"pub_key"`
}