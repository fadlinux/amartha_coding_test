package http

import (
	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
)

type Delivery struct {
	customerUC uCustomer.Usecase
}

// NewCustomerHTTP : create new instance for user http delivery
func NewCustomerHTTP(customerUC uCustomer.Usecase) Delivery {
	return Delivery{customerUC}
}
