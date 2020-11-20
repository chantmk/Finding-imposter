package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers Findingimposter-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/Findingimposter/quarantine", listQuarantineHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/quarantine", createQuarantineHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Findingimposter/quarantine/list", listSpecQuarantineHandler(cliCtx, "Findingimposter")).Methods("POST")
	r.HandleFunc("/Findingimposter/covid", listCovidHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/covid", createCovidHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Findingimposter/covid/pending", listPendingCovidHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/covid/list", listSpecCovidHandler(cliCtx, "Findingimposter")).Methods("POST")
	r.HandleFunc("/Findingimposter/log", listLogHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/log", createLogHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Findingimposter/log/list", listSpecLogHandler(cliCtx, "Findingimposter")).Methods("POST")
	r.HandleFunc("/Findingimposter/doctor", listDoctorHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/doctor", createDoctorHandler(cliCtx)).Methods("POST")
}

func isin(address string, list []string) bool {
    for _, a := range list {
        if a == address {
            return true
        }
    }
    return false
}