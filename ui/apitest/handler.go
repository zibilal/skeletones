package apitest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func ListOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("List of orders")
}

func CreateOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func UpdateOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
