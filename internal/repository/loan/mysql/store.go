package mysql

import (
	"context"
	"log"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

func (ar *mySqlLoanRepo) AddLoan(ctx context.Context, req mLoan.AddLoanRequest) (lastID int64, err error) {
	res, err := addStmt.Exec(req.CifID, req.TotalAmount, req.TotalWeeks, req.InterestRate, req.Status, req.WeeklyPayment)

	if err != nil {
		log.Println("Repository AddLoan error", err)
		return
	}

	lastID, err = res.LastInsertId()
	if err != nil {
		log.Println("response AddLoan LastInsertId error : ", err)
		return
	}

	return
}

func (ar *mySqlLoanRepo) UpdateLoan(ctx context.Context, cifId int64, req mLoan.AddLoanRequest) (err error) {
	rows, err := updateStmt.Query(req.CifID, req.InterestRate, req.TotalAmount, req.TotalWeeks)

	if err != nil {
		log.Println("Repository UpdateLoan error", err)
		return
	}

	defer rows.Close()

	return
}
