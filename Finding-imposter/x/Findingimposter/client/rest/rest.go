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
	r.HandleFunc("/Findingimposter/patient", listPatientHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/patient", createPatientHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Findingimposter/Log", listLogHandler(cliCtx, "Findingimposter")).Methods("GET")
	r.HandleFunc("/Findingimposter/Log", createLogHandler(cliCtx)).Methods("POST")
}
