package mysql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewMySQLCustomerRepo(t *testing.T) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()

	mock.ExpectPrepare("INSERT")

	repoMock := NewMySQLCustomerRepo(conn)
	assert.NotNil(t, repoMock)
}
