package network

import (
	"net/http"

	"github.com/hyperhq/hypercli/api/server/httputils"
	"github.com/hyperhq/hypercli/api/server/router"
	"github.com/hyperhq/hypercli/api/server/router/local"
	"github.com/hyperhq/hypercli/errors"
	"golang.org/x/net/context"
)

// networkRouter is a router to talk with the network controller
type networkRouter struct {
	backend Backend
	routes  []router.Route
}

// NewRouter initializes a new network router
func NewRouter(b Backend) router.Router {
	r := &networkRouter{
		backend: b,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the network controller
func (r *networkRouter) Routes() []router.Route {
	return r.routes
}

func (r *networkRouter) initRoutes() {
	r.routes = []router.Route{
		// GET
		local.NewGetRoute("/networks", r.controllerEnabledMiddleware(r.getNetworksList)),
		local.NewGetRoute("/networks/{id:.*}", r.controllerEnabledMiddleware(r.getNetwork)),
		// POST
		local.NewPostRoute("/networks/create", r.controllerEnabledMiddleware(r.postNetworkCreate)),
		local.NewPostRoute("/networks/{id:.*}/connect", r.controllerEnabledMiddleware(r.postNetworkConnect)),
		local.NewPostRoute("/networks/{id:.*}/disconnect", r.controllerEnabledMiddleware(r.postNetworkDisconnect)),
		// DELETE
		local.NewDeleteRoute("/networks/{id:.*}", r.controllerEnabledMiddleware(r.deleteNetwork)),
	}
}

func (r *networkRouter) controllerEnabledMiddleware(handler httputils.APIFunc) httputils.APIFunc {
	if r.backend.NetworkControllerEnabled() {
		return handler
	}
	return networkControllerDisabled
}

func networkControllerDisabled(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	return errors.ErrorNetworkControllerNotEnabled.WithArgs()
}
