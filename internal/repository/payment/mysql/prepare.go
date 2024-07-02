package mysql

import (
	"database/sql"
	"log"
)

var (
	addStmt                 *sql.Stmt
	updateStmt              *sql.Stmt
	getExistCustomerByCifID *sql.Stmt
	getId                   *sql.Stmt
	getSumTotal             *sql.Stmt
)

func (ar *mySqlPaymentRepo) prepareAddPaymentsStmt() {
	var err error
	addStmt, err = ar.conn.Prepare("INSERT INTO payment(loan_id,amount,payment_date) VALUES( $1, $2,now()) RETURNING payment_id")
	if err != nil {
		log.Fatal("[Postgre: prepareAddPaymentsStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlPaymentRepo) prepareUpdatePaymentsStmt() {
	var err error
	updateStmt, err = ar.conn.Prepare("UPDATE payment SET amount =  $1 WHERE payment_id =  $2 ")
	if err != nil {
		log.Fatal("[Postgre: prepareUpdatePaymentsStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlPaymentRepo) prepareGetExistCustomerByCifID() {
	var err error
	getExistCustomerByCifID, err = ar.conn.Prepare("SELECT COUNT(payment_id) AS total FROM payment WHERE payment_id = $1 ")
	if err != nil {
		log.Fatal("[Postgre: prepareGetByCifID] Prepare statement fail :", err)
	}
}

func (ar *mySqlPaymentRepo) prepareGetId() {
	var err error
	getId, err = ar.conn.Prepare("SELECT amount FROM payment WHERE loan_id =  $1 ")
	if err != nil {
		log.Fatal("[Postgre: prepareGetId] Prepare statement fail :", err)
	}
}

func (ar *mySqlPaymentRepo) prepareGetCountPaymentByCifId() {
	var err error
	getSumTotal, err = ar.conn.Prepare("SELECT SUM(amount) as total_payment FROM payment WHERE loan_id =  $1")
	if err != nil {
		log.Fatal("[Postgre: prepareGetId] Prepare statement fail :", err)
	}
}
