package mysql

import (
	"context"
	"log"
)

func (ar *mySqlLoanRepo) FetchByCifId(ctx context.Context, cifId int64) (count int64, err error) {
	err = getExistCustomerByCifID.QueryRow(cifId).Scan(&count)

	if err != nil {
		log.Println("Repository FetchByCifId error", err)
		return
	}

	return
}
