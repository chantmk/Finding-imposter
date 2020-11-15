package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding
		cdc.RegisterConcrete(MsgCreateQuarantine{}, "Findingimposter/CreateQuarantine", nil)
		cdc.RegisterConcrete(MsgCreateCovid{}, "Findingimposter/CreateCovid", nil)
		cdc.RegisterConcrete(MsgCreateLog{}, "Findingimposter/CreateLog", nil)
		cdc.RegisterConcrete(MsgCreateDoctor{}, "Findingimposter/CreateDoctor", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
