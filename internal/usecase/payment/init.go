package payment

import (
	rPayment "github.com/fadlinux/amartha_coding_test/internal/repository/payment"
)

type (
	paymentUsecase struct {
		mysqlPaymentRepo rPayment.Repository
	}
)

// NewPaymentUsecase : create new instance for paymentUsecase
func NewPaymentUsecase(mysqlPayment rPayment.Repository) Usecase {
	return &paymentUsecase{mysqlPayment}
}
