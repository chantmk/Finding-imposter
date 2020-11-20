package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/Finding-imposter/x/Findingimposter/types"
)

func listCovidHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-covid", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func listPendingCovidHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-pending-covid", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type listSpecCovid struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Address []string `json:"address"`
}
func listSpecCovidHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-covid", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		var req listSpecCovid
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		address := req.Address

		var out []types.Covid
		cliCtx.Codec.MustUnmarshalJSON(res, &out)

		var filteredOut []types.Covid
			for _, covid := range out {
				var pubkey []string
				pubkey = covid.PubKey
				if covid.Status == "APPROVED" {
					for _, key := range pubkey {
						if isin(key, address){
							filteredOut = append(filteredOut, covid)
						}
					}
				}
			}
		// response = cliCtx.Codec.MusMarshalJSON(filteredOut)
		rest.PostProcessResponse(w, cliCtx, filteredOut)
	}
}
