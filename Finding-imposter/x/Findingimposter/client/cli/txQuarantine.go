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

func GetCmdCreateQuarantine(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-quarantine [startAt] [endAt]",
		Short: "Creates a new quarantine",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsUserAddress, err_addr := sdk.AccAddressFromBech32(string(args[0]))
			if err_addr != nil {
				return err_addr
			}
			argsStartAt := time.Now()
			argsEndAt := time.Now().AddDate(0, 0, 14)

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateQuarantine(cliCtx.GetFromAddress(), argsUserAddress, argsStartAt, argsEndAt)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
