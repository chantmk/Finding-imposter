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

func GetCmdCreateQuarantine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-quarantine [user_id] [start_at] [end_at]",
		Short: "Creates a new quarantine",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsUser_id := string(args[0])
      argsStart_at := string(args[1])
      argsEnd_at := string(args[2])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateQuarantine(cliCtx.GetFromAddress(), argsUser_id, argsStart_at, argsEnd_at)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
