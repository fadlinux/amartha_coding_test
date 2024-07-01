package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaymentHTTP(t *testing.T) {
	delivery := NewPaymentHTTP(nil, nil)
	assert.NotNil(t, delivery)
}
