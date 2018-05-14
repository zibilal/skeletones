package v_0_0_1

import "github.com/zibilal/skeletones/ui/api"

type V_0_0_1 struct {
	Endpoints []api.Endpoint
	Router    api.Router
	Handler   api.ApiHandler
}

func NewV_0_0_1(router api.Router, handler api.ApiHandler) *V_0_0_1 {
	v := new(V_0_0_1)
	v.Router = router
	v.Endpoints = make([]api.Endpoint, 0)
	v.Handler = handler

	return v
}

func (v *V_0_0_1) DefineRoute() error {

	endpoints := []api.Endpoint{

		{
			"GET",
			"/trct/v.0.0.1/order",
			v.Handler.Get,
		},
		{
			"POST",
			"trct/v.0.0.1/order",
			v.Handler.Create,
		},
		{
			"PUT",
			"trct/v.0.0.1/order",
			v.Handler.Update,
		},
		{
			"DELETE",
			"trct/v.0.0.1/order",
			v.Handler.Delete,
		},
	}

	for _, endpoint := range endpoints {
		v.Router.Append(endpoint)
	}

	return v.Router.Route()
}
