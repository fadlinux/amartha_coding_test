package mysql

import (
	"database/sql"
	"log"
)

var (
	addStmt                 *sql.Stmt
	updateStmt              *sql.Stmt
	getExistCustomerByCifID *sql.Stmt
	getCustomerByCifID      *sql.Stmt
)

func (ar *mySqlCustomerRepo) prepareAddStmt() {
	var err error
	addStmt, err = ar.conn.Prepare("INSERT INTO customer(fullname,ktp_no,address,create_date) VALUES(?,?,?,now())")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareAddStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlCustomerRepo) prepareUpdateStmt() {
	var err error
	updateStmt, err = ar.conn.Prepare("UPDATE customer SET fullname = ?, ktp_no = ?, address = ?, update_date = now() WHERE cif_id = ? ")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareUpdateStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlCustomerRepo) prepareGetExistCustomerByCifID() {
	var err error
	getExistCustomerByCifID, err = ar.conn.Prepare("SELECT COUNT(cif_id) AS total FROM customer WHERE cif_id = ? ")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareGetByCifID] Prepare statement fail :", err)
	}
}

func (ar *mySqlCustomerRepo) prepareGetCustomerByCifID() {
	var err error
	getCustomerByCifID, err = ar.conn.Prepare("SELECT fullname,ktp_no,address AS TOTAL FROM customer WHERE cif_id = ? ")
	if err != nil {
		log.Fatal("[User MySQL Repo: prepareGetCustomerByCifID] Prepare statement fail :", err)
	}
}
