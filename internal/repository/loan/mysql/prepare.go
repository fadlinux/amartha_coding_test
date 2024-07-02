package mysql

import (
	"database/sql"
	"log"
)

var (
	addStmt                 *sql.Stmt
	updateStmt              *sql.Stmt
	getExistCustomerByCifID *sql.Stmt
	getLoanByLoanId         *sql.Stmt
)

func (ar *mySqlLoanRepo) prepareAddLoansStmt() {
	var err error
	addStmt, err = ar.conn.Prepare("INSERT INTO loan (cif_id,total_amount,total_weeks,interest_rate, status, weekly_payment, create_time) VALUES ($1, $2, $3, $4, $5, $6,now()) RETURNING loan_id")
	if err != nil {
		log.Fatal("[Postgre: prepareAddLoansStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlLoanRepo) prepareUpdateLoansStmt() {
	var err error
	updateStmt, err = ar.conn.Prepare("UPDATE loan SET delinquent = $1 WHERE loan_id =  $2 ")
	if err != nil {
		log.Fatal("[Postgre: prepareUpdateLoansStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlLoanRepo) prepareGetExistCustomerByCifID() {
	var err error
	getExistCustomerByCifID, err = ar.conn.Prepare("SELECT COUNT(loan_id) AS total FROM loan WHERE loan_id =  $1 ")
	if err != nil {
		log.Fatal("[Postgre:prepareGetExistCustomerByCifIDprepareGetByCifID] Prepare statement fail :", err)
	}
}

func (ar *mySqlLoanRepo) prepareGetLoanByLoanId() {
	var err error
	getLoanByLoanId, err = ar.conn.Prepare("SELECT cif_id,total_amount,total_weeks,interest_rate, status, weekly_payment FROM loan WHERE loan_id =  $1")
	if err != nil {
		log.Fatal("[Postgre: prepareGetLoanByLoanId] Prepare statement fail :", err)
	}
}
