package customer

import (
	"context"
	"log"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
)

func (uc *customerUsecase) AddCustomer(ctx context.Context, request mCustomer.AddCustomerRequest) (lastID int64, err error) {
	lastID, err = uc.mysqlCustomerRepo.AddCustomer(ctx, request)
	if err != nil {
		log.Println("Usecase AddCustomer error :", err)
	}

	return
}

func (uc *customerUsecase) UpdateCustomer(ctx context.Context, cifId int64, request mCustomer.AddCustomerRequest) (err error) {
	err = uc.mysqlCustomerRepo.UpdateCustomer(ctx, cifId, request)
	if err != nil {
		log.Println("Usecase UpdateCustomer error :", err)
	}

	return
}
