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
	addStmt, err = ar.conn.Prepare("INSERT INTO loans(cif_id,total_amount,total_weeks,interest_rate, status, weekly_payment, create_time) VALUES(?,?,?,?,?,?,now())")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareAddLoansStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlLoanRepo) prepareUpdateLoansStmt() {
	var err error
	updateStmt, err = ar.conn.Prepare("UPDATE loans SET delinquent = ? WHERE loan_id = ? ")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareUpdateLoansStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlLoanRepo) prepareGetExistCustomerByCifID() {
	var err error
	getExistCustomerByCifID, err = ar.conn.Prepare("SELECT COUNT(loan_id) AS total FROM loans WHERE loan_id = ? ")
	if err != nil {
		log.Fatal("[User MySQL Repo:prepareGetExistCustomerByCifIDprepareGetByCifID] Prepare statement fail :", err)
	}
}

func (ar *mySqlLoanRepo) prepareGetLoanByLoanId() {
	var err error
	getLoanByLoanId, err = ar.conn.Prepare("SELECT cif_id,total_amount,total_weeks,interest_rate, status, weekly_payment FROM loans WHERE loan_id = ?")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareGetLoanByLoanId] Prepare statement fail :", err)
	}
}
