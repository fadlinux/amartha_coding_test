package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	cHttp "github.com/fadlinux/amartha_coding_test/common/http"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
	"github.com/julienschmidt/httprouter"
)

func (d Delivery) HandleAddPayment(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var result mPayment.AddPaymentResponse
	statusCode := http.StatusCreated
	//var err error

	loanId := req.FormValue("loan_id")
	amount, _ := strconv.ParseInt(req.FormValue("amount"), 10, 64)

	reqData := mPayment.AddPaymentRequest{
		LoanID: loanId,
		Amount: amount,
	}

	if loanId == "" {
		statusCode = http.StatusBadRequest
		result.Message = "Bad Request"
	} else {
		loanIdInt, _ := strconv.ParseInt(loanId, 10, 64)
		exist, errExist := d.loanUC.GetExistCifId(context.Background(), loanIdInt)

		if errExist != nil {
			log.Println("HandleAddPayment GetExistCifId error :", errExist)
			statusCode = http.StatusBadGateway
			result.Message = "error on HandleAddPayment "
		}

		if exist {
			var paymentId int64

			//get data loan
			dataLoan, errLoan := d.loanUC.GetLoanByLoanId(context.Background(), loanIdInt)
			if errLoan != nil {
				log.Println("HandleAddPayment GetExistCifId error :", errExist)
				statusCode = http.StatusBadGateway
				result.Message = "error on HandleAddPayment "
			}

			//check amount payment equal Weekly payment ? yes add payment, not show error
			if amount == int64(dataLoan.WeeklyPayment) {
				//Add payment
				paymentId, _ = d.paymentUC.AddPayment(context.Background(), reqData)
				dataPayment, _ := d.paymentUC.GetId(context.Background(), loanIdInt)

				weeksPayment := make([]float64, dataLoan.TotalWeeks)
				for x := range weeksPayment {
					if x > len(dataPayment) {
						weeksPayment[x] = float64(dataLoan.WeeklyPayment)
					}
				}
				var totalPayment int64
				for _, y := range dataPayment {
					totalPayment = totalPayment + y.Amount
				}

				total := dataLoan.TotalAmount + ((dataLoan.TotalAmount * dataLoan.InterestRate) / 100)
				outStanding := total - totalPayment

				status := "Open"
				if outStanding <= 0 {
					status = "Close"
				}

				result.PaymentID = strconv.Itoa(int(paymentId))
				result.Payment = mPayment.PaymentResponse{
					TotalAmount:   dataLoan.TotalAmount,
					TotalWeeks:    dataLoan.TotalWeeks,
					InterestRate:  dataLoan.InterestRate,
					WeeklyPayment: dataLoan.WeeklyPayment,
					WeeksPayment:  weeksPayment,
					OutStanding:   outStanding,
					Status:        status,
				}
				statusCode = http.StatusAccepted
				result.Message = "Success add payment"

			} else {
				statusCode = http.StatusConflict
				result.Message = "Error, payment amount not match weekly payment"
				result.Payment = mPayment.PaymentResponse{}
			}
		} else {
			statusCode = http.StatusNotFound
			result.Message = "Error, loan id not found in system"
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
