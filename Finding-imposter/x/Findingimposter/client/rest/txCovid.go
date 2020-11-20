package rest

import (
	"net/http"
	"time"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Finding-imposter/x/Findingimposter/types"
)

type createCovidRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	CovidID string `json:"covidID"`
	Status string `json:"status"`
	PubKey []string `json:"pubKey"`
	
}

func createCovidHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createCovidRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		createdAt := time.Now()
		msg := types.NewMsgCreateCovid(creator,  req.CovidID,  createdAt,  req.Status, req.PubKey)
		check := isDoctor(cliCtx, w, creator.String())
		if req.Status == "PENDING" {
			utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
		} else if req.Status == "REJECTED" {
			if check {
				utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
			} else {
				rest.WriteErrorResponse(w, http.StatusBadRequest, "Wrong user, you must be a doctor")
			}
		} else if req.Status == "APPROVED" {
			if check {
				utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
			} else {
				rest.WriteErrorResponse(w, http.StatusBadRequest, "Wrong user, you must be a doctor")
			}
		} else {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Invalid status")
		}
	}
}

func isDoctor(cliCtx context.CLIContext, w http.ResponseWriter, creator string) (bool) {
	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list-doctor", "Findingimposter"), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return false
		}
	var out []types.Doctor
	cliCtx.Codec.MustUnmarshalJSON(res, &out)
	for _, doctor := range out {
		if doctor.Address == creator {
			return true
		}
	}
	return false
}