package payment

import (
	"context"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

// Repository : function contract for payment
type Repository interface {
	AddPayment(ctx context.Context, req mPayment.AddPaymentRequest) (lastID int64, err error)
	GetId(ctx context.Context, loanId int64) (data []mPayment.AddPaymentRequest, err error)
	GetSumTotal(ctx context.Context, loanId int64) (totalPayment int64, err error)
}
