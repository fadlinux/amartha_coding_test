package loan

import (
	"context"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

// Repository : function contract for loan
type Repository interface {
	AddLoan(ctx context.Context, req mLoan.AddLoanRequest) (lastID int64, err error)
	UpdateLoan(ctx context.Context, loan_id int64, delinquent int) (err error)
	FetchByCifId(ctx context.Context, cifId int64) (count int64, err error)
	FetchLoanId(ctx context.Context, loanId int64) (data mLoan.AddLoanRequest, err error)
}
