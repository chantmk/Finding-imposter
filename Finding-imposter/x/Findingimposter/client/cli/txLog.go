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
		Use:   "create-Log [id] [place_id] [user_id] [checkin_at] [checkout_at]",
		Short: "Creates a new Log",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsId := string(args[0])
      argsPlace_id := string(args[1])
      argsUser_id := string(args[2])
      argsCheckin_at := string(args[3])
      argsCheckout_at := string(args[4])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateLog(cliCtx.GetFromAddress(), argsId, argsPlace_id, argsUser_id, argsCheckin_at, argsCheckout_at)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
