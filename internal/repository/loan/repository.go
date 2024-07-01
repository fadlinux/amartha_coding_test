package loan

import (
	"context"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

// Repository : function contract for loan
type Repository interface {
	AddLoan(ctx context.Context, req mLoan.AddLoanRequest) (lastID int64, err error)
	UpdateLoan(ctx context.Context, cifId int64, req mLoan.AddLoanRequest) (err error)
	FetchByCifId(ctx context.Context, cifId int64) (count int64, err error)
}
