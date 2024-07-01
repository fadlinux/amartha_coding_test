package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
	uCustomer "github.com/fadlinux/amartha_coding_test/internal/usecase/customer"
	"github.com/julienschmidt/httprouter"
)

func TestDelivery_HandleAddCustomer(t *testing.T) {
	w := httptest.NewRecorder()
	reqAddCustomer, _ := http.NewRequest("POST", "/customer", nil)
	customerUCMock := &uCustomer.UsecaseMock{
		AddCustomerFunc: func(ctx context.Context, request mCustomer.AddCustomerRequest) (int64, error) {
			return 1, nil
		},
	}
	type fields struct {
		customerUC uCustomer.Usecase
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
				req: reqAddCustomer,
			},
			fields: fields{
				customerUC: customerUCMock,
			},
			expectEligible:      true,
			expectErrIsEligible: nil,
			expectErrDoActivity: errors.New("error here"),
		},
		{
			name: "case 2: valid test 2",
			args: args{
				w:   w,
				req: reqAddCustomer,
			},
			fields: fields{
				customerUC: customerUCMock,
			},
			expectErrIsEligible: errors.New("error here"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delivery{
				customerUC: tt.fields.customerUC,
			}
			d.HandleAddCustomer(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}

func TestDelivery_HandleUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	reqAddCustomer, _ := http.NewRequest("PUT", "/customer", nil)
	customerUCMock := &uCustomer.UsecaseMock{
		AddCustomerFunc: func(ctx context.Context, request mCustomer.AddCustomerRequest) (int64, error) {
			return 1, nil
		},
		UpdateCustomerFunc: func(ctx context.Context, cifId int64, request mCustomer.AddCustomerRequest) error {
			return nil
		},
	}
	type fields struct {
		customerUC uCustomer.Usecase
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
				req: reqAddCustomer,
			},
			fields: fields{
				customerUC: customerUCMock,
			},
			expectEligible:      true,
			expectErrIsEligible: nil,
			expectErrDoActivity: errors.New("error here"),
		},
		{
			name: "case 2: valid test 2",
			args: args{
				w:   w,
				req: reqAddCustomer,
			},
			fields: fields{
				customerUC: customerUCMock,
			},
			expectErrIsEligible: errors.New("error here"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delivery{
				customerUC: tt.fields.customerUC,
			}
			d.HandleUpdateCustomer(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}
func TestDelivery_OptionHandler(t *testing.T) {

	w := httptest.NewRecorder()
	reqCustomer, _ := http.NewRequest("GET,DELETE,OPTIONS,POST,PUT", "/dynamic-placeholder", nil)
	optionUCMock := &uCustomer.UsecaseMock{}

	type fields struct {
		customerUC uCustomer.Usecase
	}
	type args struct {
		w      http.ResponseWriter
		req    *http.Request
		params httprouter.Params
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case 1",
			fields: fields{
				customerUC: optionUCMock,
			},
			args: args{
				w:   w,
				req: reqCustomer,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delivery{
				customerUC: tt.fields.customerUC,
			}
			d.OptionHandler(tt.args.w, tt.args.req, tt.args.params)
		})
	}
}
