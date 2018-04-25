package apitest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ApiV1 struct {
	versionName string
	baseName    string
	endpoints   []Endpoint
}

func NewApiV1(version, base string) *ApiV1 {
	apiv1 := new(ApiV1)
	apiv1.versionName = version
	apiv1.baseName = base
	return apiv1
}

func (a *ApiV1) AddEndpoint(url, method string, handler EndPointHandlerFunc) {
	a.endpoints = append(a.endpoints, Endpoint{
		Url:     url,
		Method:  method,
		Handler: handler,
	})
}

func (a *ApiV1) InitApi(route *httprouter.Router) error {

	for _, e := range a.endpoints {
		switch e.Method {
		case http.MethodGet:
			route.GET(fmt.Sprintf("/%s/%s%s", a.baseName, a.versionName, e.Url), e.Handler)
		case http.MethodPost:
			route.POST(fmt.Sprintf("/%s/%s%s", a.baseName, a.versionName, e.Url), e.Handler)
		case http.MethodPut:
			route.PUT(fmt.Sprintf("/%s/%s%s", a.baseName, a.versionName, e.Url), e.Handler)
		case http.MethodDelete:
			route.DELETE(fmt.Sprintf("/%s/%s%s", a.baseName, a.versionName, e.Url), e.Handler)
		default:
			return fmt.Errorf("unknown method %s", e.Method)
		}
	}

	return nil

}
