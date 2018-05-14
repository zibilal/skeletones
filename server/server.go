package server

type Server interface {
	Serve(name string) error
}
