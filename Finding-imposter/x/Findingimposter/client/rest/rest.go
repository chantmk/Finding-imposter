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
	r.HandleFunc("/Findingimposter/covid", listCovidHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/pendingCovid", listPendingCovidHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/covid", createCovidHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Findingimposter/log", listLogHandler(cliCtx, "Findingimposter")).Methods("GET") // all
	r.HandleFunc("/Findingimposter/log/list", listLogByAddressHandler(cliCtx, "Findingimposter")).Methods("POST") // by address
	r.HandleFunc("/Findingimposter/log", createLogHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Findingimposter/doctor", listDoctorHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/doctor", createDoctorHandler(cliCtx)).Methods("POST")
}
