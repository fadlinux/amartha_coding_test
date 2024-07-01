package mysql

import (
	"context"
	"log"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

func (ar *mySqlPaymentRepo) AddPayment(ctx context.Context, req mPayment.AddPaymentRequest) (lastID int64, err error) {
	res, err := addStmt.Exec(req.LoanID, req.Amount)

	if err != nil {
		log.Println("Repository AddPayment error", err)
		return
	}

	lastID, err = res.LastInsertId()
	if err != nil {
		log.Println("response AddPayment LastInsertId error : ", err)
		return
	}

	return
}
