package mysql

import (
	"context"
	"log"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

func (ar *mySqlPaymentRepo) AddPayment(ctx context.Context, req mPayment.AddPaymentRequest) (lastID int64, err error) {
	err = addStmt.QueryRow(req.LoanID, req.Amount).Scan(&lastID)

	if err != nil {
		log.Println("Repository AddPayment error", err)
		return
	}

	return
}
