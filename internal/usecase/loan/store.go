package loan

import (
	"context"
	"log"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

func (uc *loanUsecase) AddLoan(ctx context.Context, request mLoan.AddLoanRequest) (lastID int64, err error) {
	lastID, err = uc.mysqlLoanRepo.AddLoan(ctx, request)
	if err != nil {
		log.Println("Usecase AddLoan error :", err)
	}

	return
}

func (uc *loanUsecase) UpdateLoan(ctx context.Context, cifId int64, request mLoan.AddLoanRequest) (err error) {
	err = uc.mysqlLoanRepo.UpdateLoan(ctx, cifId, request)
	if err != nil {
		log.Println("Usecase UpdateLoan error :", err)
	}

	return
}
