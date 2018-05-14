package apiwrapper

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/zibilal/skeletones/ui/api"
	"net/http/httptest"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestNewEchoRouter(t *testing.T) {
	echolib := echo.New()
	eroute := NewEchoRouter(echolib)
	eroute.Append(api.Endpoint{
		Path:   "/test/echo",
		Method: echo.POST,
		Handler: func(ctx interface{}) error {
			t.Log("[POST]echo executed")
			switch ctx.(type) {
			case echo.Context:
				t.Log("expected type context")
				return nil
			default:
				t.Logf("unexpected type context %t", ctx)
				return errors.New("unexpected type context")
			}
		},
	})
	eroute.Append(api.Endpoint{
		Path:   "/test/echo",
		Method: echo.GET,
		Handler: func(ctx interface{}) error {
			t.Log("[GET]echo executed")
			switch ctx.(type) {
			case echo.Context:
				t.Log("expected type context")
				return nil
			default:
				t.Logf("unexpected type context %T", ctx)
				return errors.New("unexpected type context")
			}
		},
	})

	err := eroute.Route()

	t.Log("Testing api wrapper, Route()")
	{
		if err != nil {
			t.Errorf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}
	}

	t.Log("Testing api wrapper, endpoints")
	{
		endpoints := eroute.Endpoints()
		if len(endpoints) == 2 {
			t.Logf("%s expected route have two endpoints", success)
		} else {
			t.Errorf("%s expected route have two enpoints, got %d", failed, len(endpoints))
		}

		req := httptest.NewRequest(echo.GET, "/test/echo", nil)
		resp := httptest.NewRecorder()
		echoCtx := echolib.NewContext(req, resp)
		err := endpoints[1].Handler(echoCtx)

		t.Log("ZE Error", err)
	}
}
