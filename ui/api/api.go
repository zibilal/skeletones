package api

type Router interface {
	Route() error
	Append(endpoint Endpoint)
	Endpoints() []Endpoint
}

type Endpoint struct {
	Method  string
	Path    string
	Handler RouteHandler
}

type RouteHandler func(interface{}) error

type ApiHandler interface {
	Get(ctx interface{}) error
	Create(ctx interface{}) error
	Update(ctx interface{}) error
	Delete(ctx interface{}) error
}

type VersionApi interface {
	DefineRoute() error
}
