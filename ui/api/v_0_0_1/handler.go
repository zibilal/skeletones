package v_0_0_1

import "github.com/zibilal/skeletones/service"

type TerracottaHandler struct {
	services []service.Service
}

func NewTerracottaHandler(services ...service.Service) *TerracottaHandler {
	t := new(TerracottaHandler)
	t.services = make([]service.Service, 0)
	t.services = append(t.services, services...)

	return t
}

func (t *TerracottaHandler) Get(ctx interface{}) error {
	return nil
}

func (t *TerracottaHandler) Create(ctx interface{}) error {
	return nil
}

func (t *TerracottaHandler) Update(ctx interface{}) error {
	return nil
}

func (t *TerracottaHandler) Delete(ctx interface{}) error {
	return nil
}
