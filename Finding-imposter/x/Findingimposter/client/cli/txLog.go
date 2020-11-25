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

func GetCmdCreateLog(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-log [logID] [placeID] [action]",
		Short: "Creates a new log",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsLogID := string(args[0])
      argsPlaceID := string(args[1])
      argsCreatedAt := time.Now()
      argsAction := string(args[2])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateLog(cliCtx.GetFromAddress(), argsLogID, argsPlaceID, argsCreatedAt, argsAction)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
