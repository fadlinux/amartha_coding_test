package mysql

import (
	"context"
	"log"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

func (ar *mySqlLoanRepo) AddLoan(ctx context.Context, req mLoan.AddLoanRequest) (lastID int64, err error) {
	err = addStmt.QueryRow(req.CifID, req.TotalAmount, req.TotalWeeks, req.InterestRate, req.Status, req.WeeklyPayment).Scan(&lastID)

	if err != nil {
		log.Println("Repository AddLoan error", err)
		return
	}

	return
}

func (ar *mySqlLoanRepo) UpdateLoan(ctx context.Context, loan_id int64, delinquent int) (err error) {
	rows, err := updateStmt.Query(delinquent, loan_id)

	if err != nil {
		log.Println("Repository UpdateLoan error", err)
		return
	}

	defer rows.Close()

	return
}
