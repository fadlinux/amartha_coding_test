package main

import (
	"net/http"

	configCmd "github.com/fadlinux/amartha_coding_test/cmd/config"

	customerCmd "github.com/fadlinux/amartha_coding_test/cmd/customer"
	loanCMD "github.com/fadlinux/amartha_coding_test/cmd/loan"
	paymentCMD "github.com/fadlinux/amartha_coding_test/cmd/payment"

	cHttp "github.com/fadlinux/amartha_coding_test/common/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"
)

func initModule() {
	configCmd.Initialize()
	customerCmd.Initialize()
	loanCMD.Initialize()
	paymentCMD.Initialize()
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

	router.POST("/customer", customerCmd.HTTPDelivery.HandleAddCustomer)
	router.PUT("/customer", customerCmd.HTTPDelivery.HandleUpdateCustomer)

	router.POST("/loans", loanCMD.HTTPDelivery.HandleAddLoan)

	//update loans for delinquent if any
	router.PUT("/loans", loanCMD.HTTPDelivery.HandleAddLoan)

	//router.GET("/loans", loanCMD.HTTPDelivery.HandleAddLoan)

	//update deliciant IsDelinquent
	//make payment
	router.POST("/payment", paymentCMD.HTTPDelivery.HandleAddPayment)

}

func healthcheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cHttp.Render(w, "success", 0, req.FormValue("callback"))
}
