package loan

import (
	"context"
	"testing"

	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
	rLoan "github.com/fadlinux/amartha_coding_test/internal/repository/loan"
)

func Test_loanUC_AddLoan(t *testing.T) {
	rloanMock := &rLoan.RepositoryMock{
		AddLoanFunc: func(ctx context.Context, req mLoan.AddLoanRequest) (int64, error) {
			return 1, nil
		},
	}
	type fields struct {
		mysqlUserRepo rLoan.Repository
	}
	type args struct {
		ctx    context.Context
		data   mLoan.AddLoanResponse
		params mLoan.AddLoanRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case 1: normal case",
			args: args{
				ctx: context.Background(),
				params: mLoan.AddLoanRequest{
					CifID:        "1",
					TotalAmount:  5000000,
					InterestRate: 10,
					TotalWeeks:   50,
				},
				data: mLoan.AddLoanResponse{
					Message: "success",
				},
			},
			fields: fields{
				mysqlUserRepo: rloanMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &loanUsecase{
				mysqlLoanRepo: tt.fields.mysqlUserRepo,
			}
			if _, err := ac.AddLoan(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("loanUC.AddLoan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
