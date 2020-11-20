package cli

import (
	"bufio"
	"time"
	"errors"
	"fmt"
  
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
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCovidID := string(args[0])
			argsCreatedAt := time.Now()
			argsStatus := string(args[1])
			argsPubKey := []string(args[2:])		
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			creator := cliCtx.GetFromAddress()
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateCovid(creator, argsCovidID, argsCreatedAt, argsStatus, argsPubKey)
			err := msg.ValidateBasic()
			check := isDoctor(cliCtx, cdc, creator.String())
			if err != nil {
				return err
			}
			
			if argsStatus == "PENDING" {
				return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
			} else if argsStatus == "REJECTED" {
				if check {
					return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
				} else {
					return errors.New("You are not doctor")
				}
			} else if argsStatus == "APPROVED" {
				if check {
					return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
				} else {
					return errors.New("You are not doctor")
				}
			} else {
				return errors.New("invalid status")
			}
		},
	}
}

func isDoctor(cliCtx context.CLIContext,cdc *codec.Codec, creator string) (bool) {
	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListDoctor, "Findingimposter"), nil)
	if err != nil {
		fmt.Printf("could not list Doctor\n%s\n", err.Error())
		return false
	}
	var out []types.Doctor
	cdc.MustUnmarshalJSON(res, &out)
	for _, doctor := range out {
		if doctor.Address == creator {
			return true
		}
	}
	return false
}
