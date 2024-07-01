package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
	uLoan "github.com/fadlinux/amartha_coding_test/internal/usecase/loan"
	"github.com/julienschmidt/httprouter"
)

func TestDelivery_HandleAddLoan(t *testing.T) {
	w := httptest.NewRecorder()
	reqAddLoan, _ := http.NewRequest("POST", "/loan", nil)
	loanUCMock := &uLoan.UsecaseMock{
		AddLoanFunc: func(ctx context.Context, request mLoan.AddLoanRequest) (int64, error) {
			return 1, nil
		},
	}
	type fields struct {
		loanUC uLoan.Usecase
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
				req: reqAddLoan,
			},
			fields: fields{
				loanUC: loanUCMock,
			},
			expectEligible:      true,
			expectErrIsEligible: nil,
			expectErrDoActivity: errors.New("error here"),
		},
		{
			name: "case 2: valid test 2",
			args: args{
				w:   w,
				req: reqAddLoan,
			},
			fields: fields{
				loanUC: loanUCMock,
			},
			expectErrIsEligible: errors.New("error here"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Delivery{
				loanUC: tt.fields.loanUC,
			}
			d.HandleAddLoan(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}
