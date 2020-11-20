package cli

import (
	"bufio"
	"time"
  
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/Finding-imposter/x/Findingimposter/types"
)

func GetCmdCreateCovid(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-covid [covidID] [status] [pubKey]",
		Short: "Creates a new covid",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsCovidID := string(args[0])
      argsCreatedAt := time.Now()
	  argsStatus := string(args[1])
	  argsPubKey := []string(args[2:])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateCovid(cliCtx.GetFromAddress(), argsCovidID, argsCreatedAt, argsStatus, argsPubKey)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
