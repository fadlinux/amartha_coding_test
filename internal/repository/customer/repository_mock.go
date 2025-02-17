// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package customer

import (
	"context"
	mCustomer "github.com/fadlinux/amartha_coding_test/internal/model/customer"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			AddCustomerFunc: func(ctx context.Context, req mCustomer.AddCustomerRequest) (int64, error) {
//				panic("mock out the AddCustomer method")
//			},
//			FetchByCifIdFunc: func(ctx context.Context, cifId int64) (int64, error) {
//				panic("mock out the FetchByCifId method")
//			},
//			FetchCustomerByCifIdFunc: func(ctx context.Context, cifId int64) (mCustomer.AddCustomerRequest, error) {
//				panic("mock out the FetchCustomerByCifId method")
//			},
//			UpdateCustomerFunc: func(ctx context.Context, cifId int64, req mCustomer.AddCustomerRequest) error {
//				panic("mock out the UpdateCustomer method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// AddCustomerFunc mocks the AddCustomer method.
	AddCustomerFunc func(ctx context.Context, req mCustomer.AddCustomerRequest) (int64, error)

	// FetchByCifIdFunc mocks the FetchByCifId method.
	FetchByCifIdFunc func(ctx context.Context, cifId int64) (int64, error)

	// FetchCustomerByCifIdFunc mocks the FetchCustomerByCifId method.
	FetchCustomerByCifIdFunc func(ctx context.Context, cifId int64) (mCustomer.AddCustomerRequest, error)

	// UpdateCustomerFunc mocks the UpdateCustomer method.
	UpdateCustomerFunc func(ctx context.Context, cifId int64, req mCustomer.AddCustomerRequest) error

	// calls tracks calls to the methods.
	calls struct {
		// AddCustomer holds details about calls to the AddCustomer method.
		AddCustomer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req mCustomer.AddCustomerRequest
		}
		// FetchByCifId holds details about calls to the FetchByCifId method.
		FetchByCifId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CifId is the cifId argument value.
			CifId int64
		}
		// FetchCustomerByCifId holds details about calls to the FetchCustomerByCifId method.
		FetchCustomerByCifId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CifId is the cifId argument value.
			CifId int64
		}
		// UpdateCustomer holds details about calls to the UpdateCustomer method.
		UpdateCustomer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CifId is the cifId argument value.
			CifId int64
			// Req is the req argument value.
			Req mCustomer.AddCustomerRequest
		}
	}
	lockAddCustomer          sync.RWMutex
	lockFetchByCifId         sync.RWMutex
	lockFetchCustomerByCifId sync.RWMutex
	lockUpdateCustomer       sync.RWMutex
}

// AddCustomer calls AddCustomerFunc.
func (mock *RepositoryMock) AddCustomer(ctx context.Context, req mCustomer.AddCustomerRequest) (int64, error) {
	if mock.AddCustomerFunc == nil {
		panic("RepositoryMock.AddCustomerFunc: method is nil but Repository.AddCustomer was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req mCustomer.AddCustomerRequest
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockAddCustomer.Lock()
	mock.calls.AddCustomer = append(mock.calls.AddCustomer, callInfo)
	mock.lockAddCustomer.Unlock()
	return mock.AddCustomerFunc(ctx, req)
}

// AddCustomerCalls gets all the calls that were made to AddCustomer.
// Check the length with:
//
//	len(mockedRepository.AddCustomerCalls())
func (mock *RepositoryMock) AddCustomerCalls() []struct {
	Ctx context.Context
	Req mCustomer.AddCustomerRequest
} {
	var calls []struct {
		Ctx context.Context
		Req mCustomer.AddCustomerRequest
	}
	mock.lockAddCustomer.RLock()
	calls = mock.calls.AddCustomer
	mock.lockAddCustomer.RUnlock()
	return calls
}

// FetchByCifId calls FetchByCifIdFunc.
func (mock *RepositoryMock) FetchByCifId(ctx context.Context, cifId int64) (int64, error) {
	if mock.FetchByCifIdFunc == nil {
		panic("RepositoryMock.FetchByCifIdFunc: method is nil but Repository.FetchByCifId was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		CifId int64
	}{
		Ctx:   ctx,
		CifId: cifId,
	}
	mock.lockFetchByCifId.Lock()
	mock.calls.FetchByCifId = append(mock.calls.FetchByCifId, callInfo)
	mock.lockFetchByCifId.Unlock()
	return mock.FetchByCifIdFunc(ctx, cifId)
}

// FetchByCifIdCalls gets all the calls that were made to FetchByCifId.
// Check the length with:
//
//	len(mockedRepository.FetchByCifIdCalls())
func (mock *RepositoryMock) FetchByCifIdCalls() []struct {
	Ctx   context.Context
	CifId int64
} {
	var calls []struct {
		Ctx   context.Context
		CifId int64
	}
	mock.lockFetchByCifId.RLock()
	calls = mock.calls.FetchByCifId
	mock.lockFetchByCifId.RUnlock()
	return calls
}

// FetchCustomerByCifId calls FetchCustomerByCifIdFunc.
func (mock *RepositoryMock) FetchCustomerByCifId(ctx context.Context, cifId int64) (mCustomer.AddCustomerRequest, error) {
	if mock.FetchCustomerByCifIdFunc == nil {
		panic("RepositoryMock.FetchCustomerByCifIdFunc: method is nil but Repository.FetchCustomerByCifId was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		CifId int64
	}{
		Ctx:   ctx,
		CifId: cifId,
	}
	mock.lockFetchCustomerByCifId.Lock()
	mock.calls.FetchCustomerByCifId = append(mock.calls.FetchCustomerByCifId, callInfo)
	mock.lockFetchCustomerByCifId.Unlock()
	return mock.FetchCustomerByCifIdFunc(ctx, cifId)
}

// FetchCustomerByCifIdCalls gets all the calls that were made to FetchCustomerByCifId.
// Check the length with:
//
//	len(mockedRepository.FetchCustomerByCifIdCalls())
func (mock *RepositoryMock) FetchCustomerByCifIdCalls() []struct {
	Ctx   context.Context
	CifId int64
} {
	var calls []struct {
		Ctx   context.Context
		CifId int64
	}
	mock.lockFetchCustomerByCifId.RLock()
	calls = mock.calls.FetchCustomerByCifId
	mock.lockFetchCustomerByCifId.RUnlock()
	return calls
}

// UpdateCustomer calls UpdateCustomerFunc.
func (mock *RepositoryMock) UpdateCustomer(ctx context.Context, cifId int64, req mCustomer.AddCustomerRequest) error {
	if mock.UpdateCustomerFunc == nil {
		panic("RepositoryMock.UpdateCustomerFunc: method is nil but Repository.UpdateCustomer was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		CifId int64
		Req   mCustomer.AddCustomerRequest
	}{
		Ctx:   ctx,
		CifId: cifId,
		Req:   req,
	}
	mock.lockUpdateCustomer.Lock()
	mock.calls.UpdateCustomer = append(mock.calls.UpdateCustomer, callInfo)
	mock.lockUpdateCustomer.Unlock()
	return mock.UpdateCustomerFunc(ctx, cifId, req)
}

// UpdateCustomerCalls gets all the calls that were made to UpdateCustomer.
// Check the length with:
//
//	len(mockedRepository.UpdateCustomerCalls())
func (mock *RepositoryMock) UpdateCustomerCalls() []struct {
	Ctx   context.Context
	CifId int64
	Req   mCustomer.AddCustomerRequest
} {
	var calls []struct {
		Ctx   context.Context
		CifId int64
		Req   mCustomer.AddCustomerRequest
	}
	mock.lockUpdateCustomer.RLock()
	calls = mock.calls.UpdateCustomer
	mock.lockUpdateCustomer.RUnlock()
	return calls
}
