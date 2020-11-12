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
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
)

func GetCmdCreateCovid(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-covid [status] [pub_key]",
		Short: "Creates a new covid",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsStatus := string(args[0])
	  argsCreated_at := time.Now().Format("02/01/2006 15:04")
	  argsPub_key := []string(args[1:])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateCovid(cliCtx.GetFromAddress(), argsStatus, argsCreated_at, argsPub_key)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
