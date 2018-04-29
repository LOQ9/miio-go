// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import net "net"

// OutboundConn is an autogenerated mock type for the OutboundConn type
type OutboundConn struct {
	mock.Mock
}

// WriteTo provides a mock function with given fields: _a0, _a1
func (_m *OutboundConn) WriteTo(_a0 []byte, _a1 net.Addr) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte, net.Addr) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, net.Addr) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
