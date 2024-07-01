package customer

import (
	"context"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
)

// Usecase : contract of functions in Usecase layer
type Usecase interface {
	AddCustomer(ctx context.Context, request mCustomer.AddCustomerRequest) (lastID int64, err error)
	UpdateCustomer(ctx context.Context, cifId int64, request mCustomer.AddCustomerRequest) (err error)
	GetExistCifId(ctx context.Context, cifId int64) (exist bool, err error)
	GetCustomer(ctx context.Context, cifId int64) (data mCustomer.AddCustomerRequest, err error)
}
