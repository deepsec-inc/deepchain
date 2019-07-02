// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import discovery "deepchain/discovery/client"
import mock "github.com/stretchr/testify/mock"
import protosdiscovery "deepchain/protos/discovery"

// ServiceResponse is an autogenerated mock type for the ServiceResponse type
type ServiceResponse struct {
	mock.Mock
}

// ForChannel provides a mock function with given fields: _a0
func (_m *ServiceResponse) ForChannel(_a0 string) discovery.ChannelResponse {
	ret := _m.Called(_a0)

	var r0 discovery.ChannelResponse
	if rf, ok := ret.Get(0).(func(string) discovery.ChannelResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.ChannelResponse)
		}
	}

	return r0
}

// ForLocal provides a mock function with given fields:
func (_m *ServiceResponse) ForLocal() discovery.LocalResponse {
	ret := _m.Called()

	var r0 discovery.LocalResponse
	if rf, ok := ret.Get(0).(func() discovery.LocalResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.LocalResponse)
		}
	}

	return r0
}

// Raw provides a mock function with given fields:
func (_m *ServiceResponse) Raw() *protosdiscovery.Response {
	ret := _m.Called()

	var r0 *protosdiscovery.Response
	if rf, ok := ret.Get(0).(func() *protosdiscovery.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protosdiscovery.Response)
		}
	}

	return r0
}
