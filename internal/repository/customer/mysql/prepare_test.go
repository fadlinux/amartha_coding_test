package mysql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_mySqlCustomerRepo_prepareAddStmt(t *testing.T) {
	type fields struct {
		conn *sql.DB
	}

	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "normal case",
			fields: fields{
				conn: conn,
			},
		},
	}
	for _, tt := range tests {
		mock.ExpectPrepare("INSERT")

		t.Run(tt.name, func(t *testing.T) {
			a := &mySqlCustomerRepo{
				conn: tt.fields.conn,
			}
			a.prepareAddStmt()
		})
	}
}
