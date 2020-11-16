package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
  "github.com/Finding-imposter/x/Findingimposter/types"
)

func GetCmdListCovid(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-covid",
		Short: "list all covid",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListCovid, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Covid\n%s\n", err.Error())
				return nil
			}
			var out []types.Covid
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdListPendingCovid(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-pending-covid",
		Short: "list all pending covid",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListPendingCovid, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list pending Covid\n%s\n", err.Error())
				return nil
			}
			var out []types.Covid
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}