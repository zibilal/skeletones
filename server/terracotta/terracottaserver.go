package terracotta

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/zibilal/skeletones/service"
	"github.com/zibilal/skeletones/ui/api"
	"github.com/zibilal/skeletones/ui/api/apiwrapper"
	"github.com/zibilal/skeletones/ui/api/v_0_0_1"
)

type TerracottaServer struct {
	name     string
	port     int
	maps     map[string]api.VersionApi
	echoType *echo.Echo
}

func NewTerracottaServer(name string, port int, services ...service.Service) *TerracottaServer {
	s := new(TerracottaServer)
	s.name = name

	s.maps = make(map[string]api.VersionApi)

	s.echoType = echo.New()
	echoRouter := apiwrapper.NewEchoRouter(s.echoType)
	terracottaHandler := v_0_0_1.NewTerracottaHandler(services...)
	s.maps["v.0.0.1"] = v_0_0_1.NewV_0_0_1(echoRouter, terracottaHandler)

	return s
}

func (s *TerracottaServer) Serve(name string) error {
	switch name {
	case "v.0.0.1":
		s.maps[name].DefineRoute()

		return s.echoType.Start(fmt.Sprintf(":%d", s.port))
	default:
		return fmt.Errorf("unrecognized server %s", name)
	}
}
