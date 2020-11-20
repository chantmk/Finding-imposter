package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/Finding-imposter/x/Findingimposter/types")

func listQuarantineHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-quarantine", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type listSpecQuarantine struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Address []string `json:address`
}
func listSpecQuarantineHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-quarantine", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		var req listSpecQuarantine
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		address := req.Address

		var out []types.Quarantine
		cliCtx.Codec.MustUnmarshalJSON(res, &out)

		var filteredOut []types.Quarantine
			for _, quarantine := range out {
				if isin(quarantine.UserAddress.String(),address) {
					filteredOut = append(filteredOut, quarantine)
				}
			}
		// response = cliCtx.Codec.MusMarshalJSON(filteredOut)
		rest.PostProcessResponse(w, cliCtx, filteredOut)
	}
}
