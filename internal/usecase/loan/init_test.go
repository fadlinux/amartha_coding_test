package loan

import (
	"testing"

	rLoan "github.com/fadlinux/amartha_coding_test/internal/repository/loan"
	"github.com/stretchr/testify/assert"
)

func TestNewLoanUsecase(t *testing.T) {
	var (
		mysqlLoanRepo rLoan.Repository
	)
	usecaseMock := NewLoanUsecase(mysqlLoanRepo)
	assert.NotNil(t, usecaseMock)
}
