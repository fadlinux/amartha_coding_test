package http

import (
	uPayment "github.com/fadlinux/amartha_coding_test/internal/usecase/payment"

	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
)

type Delivery struct {
	paymentUC uPayment.Usecase
	loanUC    uLoan.Usecase
}

// NewPaymentHTTP : create new instance for payment http delivery
func NewPaymentHTTP(paymentUC uPayment.Usecase, loanUC uLoan.Usecase) Delivery {
	return Delivery{paymentUC, loanUC}
}
