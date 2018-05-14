package apiwrapper

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/zibilal/skeletones/ui/api"
)

type EchoRouter struct {
	echoRouter *echo.Echo
	endpoints  []api.Endpoint
}

func NewEchoRouter(e *echo.Echo) *EchoRouter {
	router := new(EchoRouter)
	router.echoRouter = e
	router.endpoints = make([]api.Endpoint, 0)
	return router
}

func (r *EchoRouter) Append(endpoint api.Endpoint) {
	r.endpoints = append(r.endpoints, endpoint)
}

func (r *EchoRouter) Endpoints() []api.Endpoint {
	return r.endpoints
}

func (r *EchoRouter) Route() error {
	for _, endpoint := range r.endpoints {
		switch endpoint.Method {
		case echo.GET:
			r.echoRouter.GET(endpoint.Path, func(ctx echo.Context) error {
				return endpoint.Handler(ctx)
			})
		case echo.DELETE:
			r.echoRouter.DELETE(endpoint.Path, func(ctx echo.Context) error {
				return endpoint.Handler(ctx)
			})
		case echo.POST:
			r.echoRouter.POST(endpoint.Path, func(ctx echo.Context) error {
				return endpoint.Handler(ctx)
			})
		case echo.PUT:
			r.echoRouter.PUT(endpoint.Path, func(ctx echo.Context) error {
				return endpoint.Handler(ctx)
			})
		default:
			return fmt.Errorf("unknown method %s for endpoint %s", endpoint.Method, endpoint.Path)
		}
	}

	return nil
}
