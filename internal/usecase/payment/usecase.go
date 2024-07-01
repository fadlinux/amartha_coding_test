package payment

import (
	"context"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

// Usecase : contract of functions in Usecase layer
type Usecase interface {
	AddPayment(ctx context.Context, request mPayment.AddPaymentRequest) (lastID int64, err error)
	GetId(ctx context.Context, loanId int64) (data []mPayment.AddPaymentRequest, err error)
	GetSumTotal(ctx context.Context, loanId int64) (total int64, err error)
}
