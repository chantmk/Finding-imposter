package cli

import (
	"bufio"
  
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/chantmk/Finding-imposter/x/Findingimposter/types"
)

func GetCmdCreateLog(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-log [place_id] [check_in_at] [check_out_at]",
		Short: "Creates a new log",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsPlace_id := string(args[0])
      argsCheck_in_at := string(args[1])
      argsCheck_out_at := string(args[2])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateLog(cliCtx.GetFromAddress(), argsPlace_id, argsCheck_in_at, argsCheck_out_at)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
