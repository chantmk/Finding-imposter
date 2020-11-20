package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Finding-imposter/x/Findingimposter/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group Findingimposter queries under a subcommand
	FindingimposterQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	FindingimposterQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding
			GetCmdListQuarantine(queryRoute, cdc),
			GetCmdListSpecQuarantine(queryRoute, cdc),
			GetCmdListCovid(queryRoute, cdc),
			GetCmdListSpecCovid(queryRoute, cdc),
			GetCmdListLog(queryRoute, cdc),
			GetCmdListDoctor(queryRoute, cdc),
			GetCmdListPendingCovid(queryRoute, cdc),
			GetCmdListSpecLog(queryRoute, cdc),
		)...,
	)

	return FindingimposterQueryCmd
}

func isin(address string, list []string) bool {
    for _, a := range list {
        if a == address {
            return true
        }
    }
    return false
}