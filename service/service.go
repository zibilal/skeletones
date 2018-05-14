package service

type Service interface {
	Serve(data ...interface{})
}
