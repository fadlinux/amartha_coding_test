package loan

import (
	"context"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

// Usecase : contract of functions in Usecase layer
type Usecase interface {
	AddLoan(ctx context.Context, request mLoan.AddLoanRequest) (lastID int64, err error)
	GetExistCifId(ctx context.Context, cifId int64) (exit bool, err error)
}
