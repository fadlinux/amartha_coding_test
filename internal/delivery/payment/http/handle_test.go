package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mPayment "github.com/fadlinux/amartha_coding_test/internal/model/payment"
	uPayment "github.com/fadlinux/amartha_coding_test/internal/usecase/payment"
	"github.com/julienschmidt/httprouter"
)

func TestDelivery_HandleAddPayment(t *testing.T) {
	w := httptest.NewRecorder()
	reqAddPayment, _ := http.NewRequest("POST", "/payment", nil)
	paymentUCMock := &uPayment.UsecaseMock{
		AddPaymentFunc: func(ctx context.Context, request mPayment.AddPaymentRequest) (int64, error) {
			return 1, nil
		},
	}
	type fields struct {
		paymentUC uPayment.Usecase
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		expectEligible      bool
		expectErrIsEligible error
		expectErrDoActivity error
	}{
		{
			name: "case 1: valid test",
			args: args{
				w:   w,
				req: reqAddPayment,
			},
			fields: fields{
				paymentUC: paymentUCMock,
			},
			expectEligible:      true,
			expectErrIsEligible: nil,
			expectErrDoActivity: errors.New("error here"),
		},
		{
			name: "case 2: valid test 2",
			args: args{
				w:   w,
				req: reqAddPayment,
			},
			fields: fields{
				paymentUC: paymentUCMock,
			},
			expectErrIsEligible: errors.New("error here"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delivery{
				paymentUC: tt.fields.paymentUC,
			}
			d.HandleAddPayment(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}
