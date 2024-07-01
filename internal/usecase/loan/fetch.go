package loan

import (
	"context"
	"log"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

func (uc *loanUsecase) GetExistCifId(ctx context.Context, cifId int64) (exist bool, err error) {
	var count int64
	count, err = uc.mysqlLoanRepo.FetchByCifId(ctx, cifId)
	if err != nil {
		exist = false
		log.Println("Usecase GetExistCifId error :", err)
	}

	if count > 0 {
		exist = true
	}

	return
}

func (uc *loanUsecase) GetLoanByLoanId(ctx context.Context, loanId int64) (data mLoan.AddLoanRequest, err error) {
	data, err = uc.mysqlLoanRepo.FetchLoanId(ctx, loanId)
	if err != nil {
		log.Println("Usecase GetLoanByLoanId error :", err)
	}

	return

}
