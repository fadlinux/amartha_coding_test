package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"
)

func initModule() {

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

}

func healthcheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cHttp.Render(w, "success", 0, req.FormValue("callback"))
}
