package loan

import (
	"context"
	"log"
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
