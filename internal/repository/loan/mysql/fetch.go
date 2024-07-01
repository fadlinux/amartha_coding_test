package mysql

import (
	"context"
	"log"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
)

func (ar *mySqlLoanRepo) FetchByCifId(ctx context.Context, cifId int64) (count int64, err error) {
	err = getExistCustomerByCifID.QueryRow(cifId).Scan(&count)
	if err != nil {
		log.Println("Repository FetchByCifId error", err)
		return
	}

	return
}

func (ar *mySqlLoanRepo) FetchLoanId(ctx context.Context, loanId int64) (data mLoan.AddLoanRequest, err error) {
	rows, err := getLoanByLoanId.Query(loanId)
	if err != nil {
		log.Println("Repository FetchLoanId error", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		item := mLoan.AddLoanRequest{}
		rows.Scan(
			&item.CifID,
			&item.TotalAmount,
			&item.TotalWeeks,
			&item.InterestRate,
			&item.Status,
			&item.WeeklyPayment,
		)
		data.CifID = item.CifID
		data.TotalAmount = item.TotalAmount
		data.TotalWeeks = item.TotalWeeks
		data.InterestRate = item.InterestRate
		data.WeeklyPayment = item.WeeklyPayment
		data.Status = item.Status
	}

	return
}
