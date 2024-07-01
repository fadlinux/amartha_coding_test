package mysql

import (
	"context"
	"fmt"
	"log"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
)

func (ar *mySqlPaymentRepo) GetId(ctx context.Context, loanId int64) (data []mPayment.AddPaymentRequest, err error) {
	rows, err := getId.Query(loanId)
	if err != nil {
		log.Println("Repository GetId error", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		item := mPayment.AddPaymentRequest{}
		rows.Scan(
			&item.Amount,
		)

		fmt.Println("item Amount :", item.Amount)

		data = append(data, item)
	}
	return
}

func (ar *mySqlPaymentRepo) GetSumTotal(ctx context.Context, loanId int64) (totalPayment int64, err error) {
	rows, err := getSumTotal.Query(loanId)
	if err != nil {
		log.Println("Repository GetSumTotal error", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		item := mPayment.TotalPaymentResponse{}
		rows.Scan(
			&item.TotalPayment,
		)
		totalPayment = item.TotalPayment
	}

	return
}
