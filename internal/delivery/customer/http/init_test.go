package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerHTTP(t *testing.T) {
	delivery := NewCustomerHTTP(nil)
	assert.NotNil(t, delivery)
}
