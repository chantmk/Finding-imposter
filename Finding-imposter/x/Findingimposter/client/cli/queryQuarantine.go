package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
  "github.com/Finding-imposter/x/Findingimposter/types"
)

func GetCmdListQuarantine(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-quarantine",
		Short: "list all quarantine",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListQuarantine, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Quarantine\n%s\n", err.Error())
				return nil
			}
			var out []types.Quarantine
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
