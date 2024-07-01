package customer

import (
	"context"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
)

// Repository : function contract for customer
type Repository interface {
	AddCustomer(ctx context.Context, req mCustomer.AddCustomerRequest) (lastID int64, err error)
	UpdateCustomer(ctx context.Context, cifId int64, req mCustomer.AddCustomerRequest) (err error)
	FetchByCifId(ctx context.Context, cifId int64) (count int64, err error)
	FetchCustomerByCifId(ctx context.Context, cifId int64) (data mCustomer.AddCustomerRequest, err error)
}
