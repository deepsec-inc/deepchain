// Code generated by mockery v1.0.0
package mocks

import endorsement "deepchain/core/handlers/endorsement/api"
import endorser "deepchain/core/endorser"
import mock "github.com/stretchr/testify/mock"

// PluginMapper is an autogenerated mock type for the PluginMapper type
type PluginMapper struct {
	mock.Mock
}

// PluginFactoryByName provides a mock function with given fields: name
func (_m *PluginMapper) PluginFactoryByName(name endorser.PluginName) endorsement.PluginFactory {
	ret := _m.Called(name)

	var r0 endorsement.PluginFactory
	if rf, ok := ret.Get(0).(func(endorser.PluginName) endorsement.PluginFactory); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(endorsement.PluginFactory)
		}
	}

	return r0
}
