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
	addStmt, err = ar.conn.Prepare("INSERT INTO customer(fullname,ktp_no,address,create_date) VALUES($1,$2,$3,now()) RETURNING cif_id")
	if err != nil {
		log.Fatal("[Postgre: prepareAddStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlCustomerRepo) prepareUpdateStmt() {
	var err error
	updateStmt, err = ar.conn.Prepare("UPDATE customer SET fullname = $1, ktp_no = $2, address = $3, update_date = now() WHERE cif_id = $4 ")
	if err != nil {
		log.Fatal("[Postgre: prepareUpdateStmt] Prepare statement fail :", err)
	}
}

func (ar *mySqlCustomerRepo) prepareGetExistCustomerByCifID() {
	var err error
	getExistCustomerByCifID, err = ar.conn.Prepare("SELECT COUNT(cif_id) AS total FROM customer WHERE cif_id = $1 ")
	if err != nil {
		log.Fatal("[Postgre: prepareGetExistCustomerByCifID] Prepare statement fail :", err)
	}
}

func (ar *mySqlCustomerRepo) prepareGetCustomerByCifID() {
	var err error
	getCustomerByCifID, err = ar.conn.Prepare("SELECT fullname,ktp_no,address AS TOTAL FROM customer WHERE cif_id =  $1")
	if err != nil {
		log.Fatal("[Postgre: prepareGetCustomerByCifID] Prepare statement fail :", err)
	}
}
