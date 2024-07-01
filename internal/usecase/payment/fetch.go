package payment

import (
	"context"
	"log"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

func (uc *paymentUsecase) GetId(ctx context.Context, loanId int64) (data []mPayment.AddPaymentRequest, err error) {
	data, err = uc.mysqlPaymentRepo.GetId(ctx, loanId)
	if err != nil {
		log.Println("Usecase GetId error :", err)
	}

	return

}

func (uc *paymentUsecase) GetSumTotal(ctx context.Context, loanId int64) (total int64, err error) {
	total, err = uc.mysqlPaymentRepo.GetSumTotal(ctx, loanId)
	if err != nil {
		log.Println("Usecase GetSumTotal error :", err)
	}

	return

}
