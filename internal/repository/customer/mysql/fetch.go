package mysql

import (
	"context"
	"log"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
)

func (ar *mySqlCustomerRepo) FetchByCifId(ctx context.Context, cifId int64) (count int64, err error) {
	err = getExistCustomerByCifID.QueryRow(cifId).Scan(&count)

	if err != nil {
		log.Println("Repository FetchByCifId error", err)
		return
	}

	return
}

func (ar *mySqlCustomerRepo) FetchCustomerByCifId(ctx context.Context, cifId int64) (data mCustomer.AddCustomerRequest, err error) {
	rows, err := getCustomerByCifID.Query(cifId)
	if err != nil {
		log.Println("Repository FetchCustomerByCifId error", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		item := mCustomer.AddCustomerRequest{}
		rows.Scan(
			&item.Fullname,
			&item.KTP,
			&item.Address,
		)

		data.Fullname = item.Fullname
		data.KTP = item.KTP
		data.Address = item.Address
	}

	return
}
