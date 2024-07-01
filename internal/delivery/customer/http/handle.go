package http

import (
	"context"
	"net/http"
	"strconv"

	cHttp "github.com/fadlinux/amartha_coding_test/common/http"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
	"github.com/julienschmidt/httprouter"
)

func (d Delivery) HandleAddCustomer(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var result mCustomer.AddCustomerResponse
	statusCode := http.StatusCreated
	fullname := req.FormValue("fullname")
	ktpNo := req.FormValue("ktp_no")
	address := req.FormValue("address")

	reqData := mCustomer.AddCustomerRequest{
		Fullname: fullname,
		KTP:      ktpNo,
		Address:  address,
	}

	if fullname == "" {
		statusCode = http.StatusBadRequest
		result.Message = "Bad Request"
	} else {
		lastID, _ := d.customerUC.AddCustomer(context.Background(), reqData)
		result.Message = "Add Customer successfully"
		result.CifID = lastID
	}

	cHttp.RenderHTTPJSON(w, result, statusCode, req.FormValue("callback"))
}

func (d Delivery) HandleUpdateCustomer(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var result mCustomer.AddCustomerResponse
	statusCode := http.StatusCreated

	cidId, _ := strconv.ParseInt(req.FormValue("cif_id"), 10, 64)
	fullname := req.FormValue("fullname")
	ktpNo := req.FormValue("ktp_no")
	address := req.FormValue("address")

	reqData := mCustomer.AddCustomerRequest{
		Fullname: fullname,
		KTP:      ktpNo,
		Address:  address,
	}

	if cidId == 0 || fullname == "" {
		statusCode = http.StatusBadRequest
		result.Message = "Bad Request"
	} else {
		err := d.customerUC.UpdateCustomer(context.Background(), cidId, reqData)
		if err != nil {
			result.Message = "Update customer error, please check your data"
		} else {
			result.Message = "Update customer successfully"
		}
	}

	cHttp.RenderHTTPJSON(w, result, statusCode, req.FormValue("callback"))
}

// OptionHandler set http option
func (d Delivery) OptionHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "GET,DELETE,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-MD5")
	w.WriteHeader(http.StatusOK)
}
