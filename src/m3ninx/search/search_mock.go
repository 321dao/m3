// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3ninx/search/types.go

package search

import (
	"reflect"

	"github.com/m3db/m3ninx/doc"
	"github.com/m3db/m3ninx/index"
	"github.com/m3db/m3ninx/postings"

	"github.com/golang/mock/gomock"
)

// MockExecutor is a mock of Executor interface
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return _m.recorder
}

// Execute mocks base method
func (_m *MockExecutor) Execute(q Query) (doc.Iterator, error) {
	ret := _m.ctrl.Call(_m, "Execute", q)
	ret0, _ := ret[0].(doc.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (_mr *MockExecutorMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Execute", reflect.TypeOf((*MockExecutor)(nil).Execute), arg0)
}

// Close mocks base method
func (_m *MockExecutor) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockExecutorMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockExecutor)(nil).Close))
}

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return _m.recorder
}

// Searcher mocks base method
func (_m *MockQuery) Searcher(rs index.Readers) (Searcher, error) {
	ret := _m.ctrl.Call(_m, "Searcher", rs)
	ret0, _ := ret[0].(Searcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Searcher indicates an expected call of Searcher
func (_mr *MockQueryMockRecorder) Searcher(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Searcher", reflect.TypeOf((*MockQuery)(nil).Searcher), arg0)
}

// MockSearcher is a mock of Searcher interface
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return _m.recorder
}

// Next mocks base method
func (_m *MockSearcher) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (_mr *MockSearcherMockRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Next", reflect.TypeOf((*MockSearcher)(nil).Next))
}

// Current mocks base method
func (_m *MockSearcher) Current() postings.List {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(postings.List)
	return ret0
}

// Current indicates an expected call of Current
func (_mr *MockSearcherMockRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Current", reflect.TypeOf((*MockSearcher)(nil).Current))
}

// Err mocks base method
func (_m *MockSearcher) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (_mr *MockSearcherMockRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Err", reflect.TypeOf((*MockSearcher)(nil).Err))
}

// NumReaders mocks base method
func (_m *MockSearcher) NumReaders() int {
	ret := _m.ctrl.Call(_m, "NumReaders")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumReaders indicates an expected call of NumReaders
func (_mr *MockSearcherMockRecorder) NumReaders() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NumReaders", reflect.TypeOf((*MockSearcher)(nil).NumReaders))
}