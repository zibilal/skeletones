package apitest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
	"net/http/httptest"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestRouter(t *testing.T) {

	router := httprouter.New()

	apiv1 := NewApiV1("v1.1", "apitest")
	apiv1.AddEndpoint("/order", "GET", ListOrders)
	apiv1.AddEndpoint("/order", "POST", CreateOrder)
	apiv1.AddEndpoint("/order", "PUT", UpdateOrder)
	apiv1.InitApi(router)

	t.Log("Testing create order")
	{
		req, err := http.NewRequest("GET", "/apitest/v1.1/order", nil)

		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil", success)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("%s expected status code %d, got %d", failed, http.StatusOK, status)
		} else {
			t.Logf("%s expected status code %d", success, status)
		}
	}
}
