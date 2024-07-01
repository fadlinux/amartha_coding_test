package loan

import (
	rLoan "github.com/fadlinux/amartha_coding_test/internal/repository/loan"
)

type (
	loanUsecase struct {
		mysqlLoanRepo rLoan.Repository
	}
)

// NewLoanUsecase : create new instance for loanUsecase
func NewLoanUsecase(mysqlLoan rLoan.Repository) Usecase {
	return &loanUsecase{mysqlLoan}
}
