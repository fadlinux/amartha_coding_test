package mysql

import (
	"context"
	"database/sql"
	"testing"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_mySqlLoanRepo_AddLoan(t *testing.T) {
	conn, sqlm, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()
	type fields struct {
		conn *sql.DB
	}
	type args struct {
		ctx    context.Context
		params mLoan.AddLoanRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case 1: normal case",
			fields: fields{
				conn: conn,
			},
			args: args{
				ctx: context.Background(),
				params: mLoan.AddLoanRequest{
					CifID:        "1",
					TotalAmount:  5000000,
					InterestRate: 10,
					TotalWeeks:   50,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sqlm.ExpectPrepare("INSERT")
		sqlm.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(0, 0))
		t.Run(tt.name, func(t *testing.T) {
			ac := &mySqlLoanRepo{
				conn: tt.fields.conn,
			}
			ac.prepareAddLoansStmt()
			if _, err := ac.AddLoan(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("mySqlLoanRepo.AddLoan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
