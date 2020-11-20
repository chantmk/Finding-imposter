package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
  "github.com/Finding-imposter/x/Findingimposter/types"
)

func GetCmdListLog(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-log",
		Short: "list all log",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListLog, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Log\n%s\n", err.Error())
				return nil
			}
			var out []types.Log
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdListSpecLog(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-spec-log [address] ...",
		Short: "list specific address log",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListLog, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list specific Log\n%s\n", err.Error())
				return nil
			}
			var out []types.Log
			cdc.MustUnmarshalJSON(res, &out)
			
			var filteredOut []types.Log
			for _, log := range out {
				if isin(log.Creator.String(),args) {
					filteredOut = append(filteredOut, log)
				}
			}
			return cliCtx.PrintOutput(filteredOut)
		},
	}
}