package http

import (
	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"

	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
)

type Delivery struct {
	loanUC     uLoan.Usecase
	customerUC uCustomer.Usecase
}

// NewLoanHTTP : create new instance for loan http delivery
func NewLoanHTTP(loanUC uLoan.Usecase, customerUC uCustomer.Usecase) Delivery {
	return Delivery{loanUC, customerUC}
}
