// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	packet "miio-go/protocol/packet"
	mock "github.com/stretchr/testify/mock"
)

// Outbound is an autogenerated mock type for the Outbound type
type Outbound struct {
	mock.Mock
}

// Call provides a mock function with given fields: method, params
func (_m *Outbound) Call(method string, params interface{}) ([]byte, error) {
	ret := _m.Called(method, params)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, interface{}) []byte); ok {
		r0 = rf(method, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(method, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallAndDeserialize provides a mock function with given fields: method, params, resp
func (_m *Outbound) CallAndDeserialize(method string, params interface{}, resp interface{}) error {
	ret := _m.Called(method, params, resp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, interface{}) error); ok {
		r0 = rf(method, params, resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handle provides a mock function with given fields: pkt
func (_m *Outbound) Handle(pkt *packet.Packet) error {
	ret := _m.Called(pkt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*packet.Packet) error); ok {
		r0 = rf(pkt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *Outbound) Send(_a0 *packet.Packet) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*packet.Packet) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
