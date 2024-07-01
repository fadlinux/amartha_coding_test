package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLoanHTTP(t *testing.T) {
	delivery := NewLoanHTTP(nil, nil)
	assert.NotNil(t, delivery)
}
