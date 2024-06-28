package main

import (
	"net/http"

	configCmd "github.com/fadlinux/amartha_coding_test/cmd/config"
	customerCmd "github.com/fadlinux/amartha_coding_test/customer/config"

	cHttp "github.com/fadlinux/amartha_coding_test/common/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"
)

func initModule() {
	configCmd.Initialize()
	customerCmd.Initialize()
}

func main() {
	initModule()
	router := httprouter.New()
	initRoute(router)
	grace.Serve(":8080", router)
}

func initRoute(router *httprouter.Router) {
	// Healthcheck
	router.HEAD("/healthcheck", healthcheck)
	router.GET("/healthcheck", healthcheck)

	router.POST("/customer/create", customerCmd.HTTPDelivery.HandleAddUser)

	//allow origin
	router.OPTIONS("/customer/create", customerCmd.HTTPDelivery.OptionHandler)

}

func healthcheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cHttp.Render(w, "success", 0, req.FormValue("callback"))
}
