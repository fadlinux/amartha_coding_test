package mysql

import (
	"context"
	"log"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
)

func (ar *mySqlCustomerRepo) AddCustomer(ctx context.Context, req mCustomer.AddCustomerRequest) (lastID int64, err error) {
	err = addStmt.QueryRow(req.Fullname, req.KTP, req.Address).Scan(&lastID)

	if err != nil {
		log.Println("Repository AddCustomer error", err)
		return
	}

	return
}

func (ar *mySqlCustomerRepo) UpdateCustomer(ctx context.Context, cifId int64, req mCustomer.AddCustomerRequest) (err error) {
	rows, err := updateStmt.Query(req.Fullname, req.KTP, req.Address, cifId)

	if err != nil {
		log.Println("Repository UpdateCustomer error", err)
		return
	}

	defer rows.Close()

	return
}
