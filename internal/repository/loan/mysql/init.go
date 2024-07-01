package mysql

import (
	"database/sql"

	rLoan "github.com/fadlinux/amartha_coding_test/internal/repository/loan"
)

type (
	mySqlLoanRepo struct {
		conn *sql.DB
	}
)

// NewMySQLLoanRepo create new object for Loan Mysql Repository
func NewMySQLLoanRepo(Conn *sql.DB) rLoan.Repository {
	repo := &mySqlLoanRepo{
		conn: Conn,
	}
	repo.prepareAddLoansStmt()
	repo.prepareUpdateLoansStmt()
	repo.prepareGetExistCustomerByCifID()

	return repo
}
