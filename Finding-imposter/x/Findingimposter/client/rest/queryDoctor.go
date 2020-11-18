package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/Finding-imposter/x/Findingimposter/types"
)

func listDoctorHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-doctor", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type listSpecDoctor struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Address []string `json:address`
}

func listSpecDoctorHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-doctor", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		var req listSpecDoctor
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		address := req.Address

		var out []types.Doctor
		cliCtx.Codec.MustUnmarshalJSON(res, &out)

		var filteredOut []types.Doctor
			for _, doctor := range out {
				if isin(doctor.Creator.String(),address) {
					filteredOut = append(filteredOut, doctor)
				}
			}
		// response = cliCtx.Codec.MusMarshalJSON(filteredOut)
		rest.PostProcessResponse(w, cliCtx, filteredOut)
	}
}
