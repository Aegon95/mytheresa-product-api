// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	sqlx "github.com/jmoiron/sqlx"
	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Setup provides a mock function with given fields:
func (_m *Database) Setup() *sqlx.DB {
	ret := _m.Called()

	var r0 *sqlx.DB
	if rf, ok := ret.Get(0).(func() *sqlx.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.DB)
		}
	}

	return r0
}
