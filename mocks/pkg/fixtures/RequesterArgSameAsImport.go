// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	json "encoding/json"

	mock "github.com/stretchr/testify/mock"
)

// RequesterArgSameAsImport is an autogenerated mock type for the RequesterArgSameAsImport type
type RequesterArgSameAsImport struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RequesterArgSameAsImport) Get(_a0 string) *json.RawMessage {
	ret := _m.Called(_a0)

	var r0 *json.RawMessage
	if rf, ok := ret.Get(0).(func(string) *json.RawMessage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*json.RawMessage)
		}
	}

	return r0
}
