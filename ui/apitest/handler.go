package apitest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ListOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("List of orders")
}

func CreateOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func UpdateOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
