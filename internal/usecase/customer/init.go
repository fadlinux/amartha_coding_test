package customer

import (
	rCustomer "github.com/fadlinux/amartha_coding_test/internal/repository/customer"
)

type (
	customerUsecase struct {
		mysqlCustomerRepo rCustomer.Repository
	}
)

// NewCustomerUsecase : create new instance for customerUsecase
func NewCustomerUsecase(mysqlCustomer rCustomer.Repository) Usecase {
	return &customerUsecase{mysqlCustomer}
}
