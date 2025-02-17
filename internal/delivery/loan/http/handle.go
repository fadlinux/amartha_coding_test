package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	cHttp "github.com/fadlinux/amartha_coding_test/common/http"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
	"github.com/julienschmidt/httprouter"
)

func (d Delivery) HandleAddLoan(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var result mLoan.AddLoanResponse
	statusCode := http.StatusCreated

	cifId, _ := strconv.ParseInt(req.FormValue("cif_id"), 10, 64)
	totalAmount, _ := strconv.ParseInt(req.FormValue("total_amount"), 10, 64)
	interestRate, _ := strconv.ParseInt(req.FormValue("interest_rate"), 10, 64)
	totalWeeks, _ := strconv.ParseInt(req.FormValue("total_weeks"), 10, 64)

	cifIdStr := strconv.Itoa(int(cifId))

	if cifId == 0 {
		statusCode = http.StatusBadRequest
		result.Message = "Bad Request"
	} else {
		exist, errExist := d.customerUC.GetExistCifId(context.Background(), cifId)

		if errExist != nil {
			log.Println("HandleAddLoan GetExistCifId error :", errExist)
			statusCode = http.StatusBadGateway
			result.Message = "error on HandleAddLoan "
		}
		if exist {

			//Assumption : check duplicate CIF ID as customer id only 1 submit loan
			checkDuplicate, _ := d.loanUC.GetExistCifId(context.Background(), cifId)
			if !checkDuplicate {
				outStanding := totalAmount + ((totalAmount * interestRate) / 100)
				weeksPayment := make([]float64, totalWeeks) //set how many week for installment term
				weeklyPayment := float32(outStanding / totalWeeks)

				//Assumption : create new one loans status is open, close if paid off all payment
				status := "Open"
				result.Loan = mLoan.LoanResponse{
					TotalAmount:   totalAmount,
					InterestRate:  interestRate,
					OutStanding:   outStanding,
					TotalWeeks:    totalWeeks,
					WeeklyPayment: weeklyPayment,
					WeeksPayment:  weeksPayment,
					Status:        status,
				}

				reqData := mLoan.AddLoanRequest{
					CifID:         cifIdStr,
					TotalAmount:   totalAmount,
					InterestRate:  interestRate,
					TotalWeeks:    totalWeeks,
					WeeklyPayment: weeklyPayment,
					Status:        status,
				}
				LoanID, _ := d.loanUC.AddLoan(context.Background(), reqData)

				dataCustomer, _ := d.customerUC.GetCustomer(context.Background(), cifId)
				result.LoanID = strconv.FormatInt(LoanID, 10)
				result.CustomerData = mLoan.CustomerData{
					FullName: dataCustomer.Fullname,
					KTP:      dataCustomer.KTP,
					Address:  dataCustomer.Address,
				}

				statusCode = http.StatusForbidden
				result.Message = "Customer loans has been successfully created"
			} else {
				statusCode = http.StatusConflict
				result.Message = "Error: CIF Id Customer already exist "
			}
		} else {
			statusCode = http.StatusConflict
			result.Message = "Customer not exist in system"
		}

	}

	cHttp.RenderHTTPJSON(w, result, statusCode, req.FormValue("callback"))
}

func (d Delivery) HandleUpdateLoan(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var resp struct {
		Message string `json:"message"`
	}
	statusCode := http.StatusCreated
	delinquent, _ := strconv.Atoi(req.FormValue("delinquent"))
	loan_id, _ := strconv.ParseInt(req.FormValue("loan_id"), 0, 64)
	d.loanUC.UpdateLoan(context.Background(), loan_id, delinquent)
	resp.Message = "Success update loans"

	cHttp.RenderHTTPJSON(w, resp, statusCode, req.FormValue("callback"))
}

// OptionHandler set http option
func (d Delivery) OptionHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "GET,DELETE,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-MD5")
	w.WriteHeader(http.StatusOK)
}
