package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/Finding-imposter/x/Findingimposter/types"
)

func listLogHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-log", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type listSpecLog struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Address []string `json:"address"`
}
func listSpecLogHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-log", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		var req listSpecLog
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		address := req.Address

		var out []types.Log
		cliCtx.Codec.MustUnmarshalJSON(res, &out)

		var filteredOut []types.Log
			for _, log := range out {
				if isin(log.Creator.String(),address) {
					filteredOut = append(filteredOut, log)
				}
			}
		// response = cliCtx.Codec.MusMarshalJSON(filteredOut)
		rest.PostProcessResponse(w, cliCtx, filteredOut)
	}
}
