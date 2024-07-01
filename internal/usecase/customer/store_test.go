package customer

import (
	"context"
	"testing"

	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
	rCustomer "github.com/fadlinux/amartha_coding_test/internal/repository/customer"
)

func Test_customerUC_AddCustomer(t *testing.T) {
	rcustomerMock := &rCustomer.RepositoryMock{
		AddCustomerFunc: func(ctx context.Context, req mCustomer.AddCustomerRequest) (int64, error) {
			return 1, nil
		},
	}
	type fields struct {
		mysqlCustomerRepo rCustomer.Repository
	}
	type args struct {
		ctx    context.Context
		data   mCustomer.AddCustomerResponse
		params mCustomer.AddCustomerRequest
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
				params: mCustomer.AddCustomerRequest{
					Fullname: "amartha",
					KTP:      "xxx",
					Address:  "jakarta",
				},
				data: mCustomer.AddCustomerResponse{
					Message: "success",
					CifID:   1,
				},
			},
			fields: fields{
				mysqlCustomerRepo: rcustomerMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &customerUsecase{
				mysqlCustomerRepo: tt.fields.mysqlCustomerRepo,
			}
			if _, err := ac.AddCustomer(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("customerUC.AddCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
