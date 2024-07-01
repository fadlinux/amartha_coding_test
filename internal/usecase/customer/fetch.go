package customer

import (
	"context"
	"log"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
)

func (uc *customerUsecase) GetExistCifId(ctx context.Context, cifId int64) (exist bool, err error) {
	var count int64
	count, err = uc.mysqlCustomerRepo.FetchByCifId(ctx, cifId)
	if err != nil {
		exist = false
		log.Println("Usecase GetExistCifId error :", err)
	}

	if count > 0 {
		exist = true
	}

	return
}

func (uc *customerUsecase) GetCustomer(ctx context.Context, cifId int64) (data mCustomer.AddCustomerRequest, err error) {
	data, err = uc.mysqlCustomerRepo.FetchCustomerByCifId(ctx, cifId)
	if err != nil {
		log.Println("Usecase GetCustomer error :", err)
	}

	return
}
