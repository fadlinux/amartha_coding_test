package mysql

import (
	"database/sql"

	rPayment "github.com/fadlinux/amartha_coding_test/internal/repository/payment"
)

type (
	mySqlPaymentRepo struct {
		conn *sql.DB
	}
)

// NewMySQLPaymentRepo create new object for Payment Mysql Repository
func NewMySQLPaymentRepo(Conn *sql.DB) rPayment.Repository {
	repo := &mySqlPaymentRepo{
		conn: Conn,
	}
	repo.prepareAddPaymentsStmt()
	repo.prepareUpdatePaymentsStmt()
	repo.prepareGetExistCustomerByCifID()
	repo.prepareGetId()
	repo.prepareGetCountPaymentByCifId()

	return repo
}
