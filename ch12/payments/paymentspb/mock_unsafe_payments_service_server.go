// Code generated by mockery v2.14.0. DO NOT EDIT.

package paymentspb

import mock "github.com/stretchr/testify/mock"

// MockUnsafePaymentsServiceServer is an autogenerated mock type for the UnsafePaymentsServiceServer type
type MockUnsafePaymentsServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedPaymentsServiceServer provides a mock function with given fields:
func (_m *MockUnsafePaymentsServiceServer) mustEmbedUnimplementedPaymentsServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewMockUnsafePaymentsServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUnsafePaymentsServiceServer creates a new instance of MockUnsafePaymentsServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafePaymentsServiceServer(t mockConstructorTestingTNewMockUnsafePaymentsServiceServer) *MockUnsafePaymentsServiceServer {
	mock := &MockUnsafePaymentsServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
