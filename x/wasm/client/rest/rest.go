// Deprecated: the rest package will be removed. You can use the GRPC gateway instead
package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers staking-related REST handlers to a router
// Deprecated: the rest package will be removed. You can use the GRPC gateway instead
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	// registerQueryRoutes(cliCtx, r)
	// registerTxRoutes(cliCtx, r)
	// registerNewTxRoutes(cliCtx, r)
}
