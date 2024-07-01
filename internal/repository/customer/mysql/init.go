package mysql

import (
	"database/sql"

	rCustomer "github.com/fadlinux/amartha_coding_test/internal/repository/customer"
)

type (
	mySqlCustomerRepo struct {
		conn *sql.DB
	}
)

// NewMySQLCustomerRepo create new object for Customer Mysql Repository
func NewMySQLCustomerRepo(Conn *sql.DB) rCustomer.Repository {
	repo := &mySqlCustomerRepo{
		conn: Conn,
	}

	repo.prepareAddStmt()
	repo.prepareUpdateStmt()
	repo.prepareGetExistCustomerByCifID()
	repo.prepareGetCustomerByCifID()

	return repo
}
