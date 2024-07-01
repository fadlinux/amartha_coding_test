package mysql

import (
	"testing"
)

func Test_NewDBConnection(t *testing.T) {
	NewDBConnection("postgres", "host")
}
