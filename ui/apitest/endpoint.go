package apitest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Endpoint struct {
	Url     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type EndPointHandlerFunc func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
