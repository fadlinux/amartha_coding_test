// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package loan

import (
	"context"
	mLoan "github.com/fadlinux/amartha_coding_test/internal/model/loan"
	"sync"
)

// Ensure, that UsecaseMock does implement Usecase.
// If this is not the case, regenerate this file with moq.
var _ Usecase = &UsecaseMock{}

// UsecaseMock is a mock implementation of Usecase.
//
//	func TestSomethingThatUsesUsecase(t *testing.T) {
//
//		// make and configure a mocked Usecase
//		mockedUsecase := &UsecaseMock{
//			AddLoanFunc: func(ctx context.Context, request mLoan.AddLoanRequest) (int64, error) {
//				panic("mock out the AddLoan method")
//			},
//			GetExistCifIdFunc: func(ctx context.Context, cifId int64) (bool, error) {
//				panic("mock out the GetExistCifId method")
//			},
//			GetLoanByLoanIdFunc: func(ctx context.Context, loanId int64) (mLoan.AddLoanRequest, error) {
//				panic("mock out the GetLoanByLoanId method")
//			},
//			UpdateLoanFunc: func(ctx context.Context, loanId int64, delinquent int) error {
//				panic("mock out the UpdateLoan method")
//			},
//		}
//
//		// use mockedUsecase in code that requires Usecase
//		// and then make assertions.
//
//	}
type UsecaseMock struct {
	// AddLoanFunc mocks the AddLoan method.
	AddLoanFunc func(ctx context.Context, request mLoan.AddLoanRequest) (int64, error)

	// GetExistCifIdFunc mocks the GetExistCifId method.
	GetExistCifIdFunc func(ctx context.Context, cifId int64) (bool, error)

	// GetLoanByLoanIdFunc mocks the GetLoanByLoanId method.
	GetLoanByLoanIdFunc func(ctx context.Context, loanId int64) (mLoan.AddLoanRequest, error)

	// UpdateLoanFunc mocks the UpdateLoan method.
	UpdateLoanFunc func(ctx context.Context, loanId int64, delinquent int) error

	// calls tracks calls to the methods.
	calls struct {
		// AddLoan holds details about calls to the AddLoan method.
		AddLoan []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Request is the request argument value.
			Request mLoan.AddLoanRequest
		}
		// GetExistCifId holds details about calls to the GetExistCifId method.
		GetExistCifId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CifId is the cifId argument value.
			CifId int64
		}
		// GetLoanByLoanId holds details about calls to the GetLoanByLoanId method.
		GetLoanByLoanId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// LoanId is the loanId argument value.
			LoanId int64
		}
		// UpdateLoan holds details about calls to the UpdateLoan method.
		UpdateLoan []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// LoanId is the loanId argument value.
			LoanId int64
			// Delinquent is the delinquent argument value.
			Delinquent int
		}
	}
	lockAddLoan         sync.RWMutex
	lockGetExistCifId   sync.RWMutex
	lockGetLoanByLoanId sync.RWMutex
	lockUpdateLoan      sync.RWMutex
}

// AddLoan calls AddLoanFunc.
func (mock *UsecaseMock) AddLoan(ctx context.Context, request mLoan.AddLoanRequest) (int64, error) {
	if mock.AddLoanFunc == nil {
		panic("UsecaseMock.AddLoanFunc: method is nil but Usecase.AddLoan was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Request mLoan.AddLoanRequest
	}{
		Ctx:     ctx,
		Request: request,
	}
	mock.lockAddLoan.Lock()
	mock.calls.AddLoan = append(mock.calls.AddLoan, callInfo)
	mock.lockAddLoan.Unlock()
	return mock.AddLoanFunc(ctx, request)
}

// AddLoanCalls gets all the calls that were made to AddLoan.
// Check the length with:
//
//	len(mockedUsecase.AddLoanCalls())
func (mock *UsecaseMock) AddLoanCalls() []struct {
	Ctx     context.Context
	Request mLoan.AddLoanRequest
} {
	var calls []struct {
		Ctx     context.Context
		Request mLoan.AddLoanRequest
	}
	mock.lockAddLoan.RLock()
	calls = mock.calls.AddLoan
	mock.lockAddLoan.RUnlock()
	return calls
}

// GetExistCifId calls GetExistCifIdFunc.
func (mock *UsecaseMock) GetExistCifId(ctx context.Context, cifId int64) (bool, error) {
	if mock.GetExistCifIdFunc == nil {
		panic("UsecaseMock.GetExistCifIdFunc: method is nil but Usecase.GetExistCifId was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		CifId int64
	}{
		Ctx:   ctx,
		CifId: cifId,
	}
	mock.lockGetExistCifId.Lock()
	mock.calls.GetExistCifId = append(mock.calls.GetExistCifId, callInfo)
	mock.lockGetExistCifId.Unlock()
	return mock.GetExistCifIdFunc(ctx, cifId)
}

// GetExistCifIdCalls gets all the calls that were made to GetExistCifId.
// Check the length with:
//
//	len(mockedUsecase.GetExistCifIdCalls())
func (mock *UsecaseMock) GetExistCifIdCalls() []struct {
	Ctx   context.Context
	CifId int64
} {
	var calls []struct {
		Ctx   context.Context
		CifId int64
	}
	mock.lockGetExistCifId.RLock()
	calls = mock.calls.GetExistCifId
	mock.lockGetExistCifId.RUnlock()
	return calls
}

// GetLoanByLoanId calls GetLoanByLoanIdFunc.
func (mock *UsecaseMock) GetLoanByLoanId(ctx context.Context, loanId int64) (mLoan.AddLoanRequest, error) {
	if mock.GetLoanByLoanIdFunc == nil {
		panic("UsecaseMock.GetLoanByLoanIdFunc: method is nil but Usecase.GetLoanByLoanId was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		LoanId int64
	}{
		Ctx:    ctx,
		LoanId: loanId,
	}
	mock.lockGetLoanByLoanId.Lock()
	mock.calls.GetLoanByLoanId = append(mock.calls.GetLoanByLoanId, callInfo)
	mock.lockGetLoanByLoanId.Unlock()
	return mock.GetLoanByLoanIdFunc(ctx, loanId)
}

// GetLoanByLoanIdCalls gets all the calls that were made to GetLoanByLoanId.
// Check the length with:
//
//	len(mockedUsecase.GetLoanByLoanIdCalls())
func (mock *UsecaseMock) GetLoanByLoanIdCalls() []struct {
	Ctx    context.Context
	LoanId int64
} {
	var calls []struct {
		Ctx    context.Context
		LoanId int64
	}
	mock.lockGetLoanByLoanId.RLock()
	calls = mock.calls.GetLoanByLoanId
	mock.lockGetLoanByLoanId.RUnlock()
	return calls
}

// UpdateLoan calls UpdateLoanFunc.
func (mock *UsecaseMock) UpdateLoan(ctx context.Context, loanId int64, delinquent int) error {
	if mock.UpdateLoanFunc == nil {
		panic("UsecaseMock.UpdateLoanFunc: method is nil but Usecase.UpdateLoan was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		LoanId     int64
		Delinquent int
	}{
		Ctx:        ctx,
		LoanId:     loanId,
		Delinquent: delinquent,
	}
	mock.lockUpdateLoan.Lock()
	mock.calls.UpdateLoan = append(mock.calls.UpdateLoan, callInfo)
	mock.lockUpdateLoan.Unlock()
	return mock.UpdateLoanFunc(ctx, loanId, delinquent)
}

// UpdateLoanCalls gets all the calls that were made to UpdateLoan.
// Check the length with:
//
//	len(mockedUsecase.UpdateLoanCalls())
func (mock *UsecaseMock) UpdateLoanCalls() []struct {
	Ctx        context.Context
	LoanId     int64
	Delinquent int
} {
	var calls []struct {
		Ctx        context.Context
		LoanId     int64
		Delinquent int
	}
	mock.lockUpdateLoan.RLock()
	calls = mock.calls.UpdateLoan
	mock.lockUpdateLoan.RUnlock()
	return calls
}
