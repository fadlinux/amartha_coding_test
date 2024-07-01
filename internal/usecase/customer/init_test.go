package customer

import (
	"testing"

	rCustomer "github.com/fadlinux/amartha_coding_test/internal/repository/customer"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomerUsecase(t *testing.T) {
	var (
		mysqlUserRepo rCustomer.Repository
	)
	usecaseMock := NewCustomerUsecase(mysqlUserRepo)
	assert.NotNil(t, usecaseMock)
}
