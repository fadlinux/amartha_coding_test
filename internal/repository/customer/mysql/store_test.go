package mysql

import (
	"context"
	"database/sql"
	"testing"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_mySqlCustomerRepo_AddCustomer(t *testing.T) {
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
		params mCustomer.AddCustomerRequest
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
				params: mCustomer.AddCustomerRequest{
					Fullname: "amartha",
					KTP:      "xxx",
					Address:  "xxx",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sqlm.ExpectPrepare("INSERT")
		sqlm.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(0, 0))
		t.Run(tt.name, func(t *testing.T) {
			ac := &mySqlCustomerRepo{
				conn: tt.fields.conn,
			}
			ac.prepareAddStmt()
			if _, err := ac.AddCustomer(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("mySqlCustomerRepo.AddCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
