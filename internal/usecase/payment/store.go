package payment

import (
	"context"
	"log"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

func (uc *paymentUsecase) AddPayment(ctx context.Context, request mPayment.AddPaymentRequest) (lastID int64, err error) {
	lastID, err = uc.mysqlPaymentRepo.AddPayment(ctx, request)
	if err != nil {
		log.Println("Usecase AddPayment error :", err)
	}

	return
}
