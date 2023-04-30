// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/dicomio/reader.go

// Package mock_dicomio is a generated GoMock package.
package mock_dicomio

import (
	binary "encoding/binary"
	gomock "github.com/golang/mock/gomock"
	charset "github.com/jamesshenjian/dicom/pkg/charset"
	reflect "reflect"
)

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReader) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReaderMockRecorder) Read(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReader)(nil).Read), p)
}

// ReadUInt8 mocks base method
func (m *MockReader) ReadUInt8() (uint8, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUInt8")
	ret0, _ := ret[0].(uint8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUInt8 indicates an expected call of ReadUInt8
func (mr *MockReaderMockRecorder) ReadUInt8() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUInt8", reflect.TypeOf((*MockReader)(nil).ReadUInt8))
}

// ReadUInt16 mocks base method
func (m *MockReader) ReadUInt16() (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUInt16")
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUInt16 indicates an expected call of ReadUInt16
func (mr *MockReaderMockRecorder) ReadUInt16() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUInt16", reflect.TypeOf((*MockReader)(nil).ReadUInt16))
}

// ReadUInt32 mocks base method
func (m *MockReader) ReadUInt32() (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUInt32")
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUInt32 indicates an expected call of ReadUInt32
func (mr *MockReaderMockRecorder) ReadUInt32() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUInt32", reflect.TypeOf((*MockReader)(nil).ReadUInt32))
}

// ReadInt16 mocks base method
func (m *MockReader) ReadInt16() (int16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInt16")
	ret0, _ := ret[0].(int16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInt16 indicates an expected call of ReadInt16
func (mr *MockReaderMockRecorder) ReadInt16() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInt16", reflect.TypeOf((*MockReader)(nil).ReadInt16))
}

// ReadInt32 mocks base method
func (m *MockReader) ReadInt32() (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInt32")
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInt32 indicates an expected call of ReadInt32
func (mr *MockReaderMockRecorder) ReadInt32() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInt32", reflect.TypeOf((*MockReader)(nil).ReadInt32))
}

// ReadFloat32 mocks base method
func (m *MockReader) ReadFloat32() (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFloat32")
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFloat32 indicates an expected call of ReadFloat32
func (mr *MockReaderMockRecorder) ReadFloat32() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFloat32", reflect.TypeOf((*MockReader)(nil).ReadFloat32))
}

// ReadFloat64 mocks base method
func (m *MockReader) ReadFloat64() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFloat64")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFloat64 indicates an expected call of ReadFloat64
func (mr *MockReaderMockRecorder) ReadFloat64() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFloat64", reflect.TypeOf((*MockReader)(nil).ReadFloat64))
}

// ReadString mocks base method
func (m *MockReader) ReadString(n uint32) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadString", n)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadString indicates an expected call of ReadString
func (mr *MockReaderMockRecorder) ReadString(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadString", reflect.TypeOf((*MockReader)(nil).ReadString), n)
}

// Skip mocks base method
func (m *MockReader) Skip(n int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Skip", n)
	ret0, _ := ret[0].(error)
	return ret0
}

// Skip indicates an expected call of Skip
func (mr *MockReaderMockRecorder) Skip(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockReader)(nil).Skip), n)
}

// Peek mocks base method
func (m *MockReader) Peek(n int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek", n)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peek indicates an expected call of Peek
func (mr *MockReaderMockRecorder) Peek(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockReader)(nil).Peek), n)
}

// PushLimit mocks base method
func (m *MockReader) PushLimit(n int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushLimit", n)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushLimit indicates an expected call of PushLimit
func (mr *MockReaderMockRecorder) PushLimit(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushLimit", reflect.TypeOf((*MockReader)(nil).PushLimit), n)
}

// PopLimit mocks base method
func (m *MockReader) PopLimit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PopLimit")
}

// PopLimit indicates an expected call of PopLimit
func (mr *MockReaderMockRecorder) PopLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopLimit", reflect.TypeOf((*MockReader)(nil).PopLimit))
}

// IsLimitExhausted mocks base method
func (m *MockReader) IsLimitExhausted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLimitExhausted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLimitExhausted indicates an expected call of IsLimitExhausted
func (mr *MockReaderMockRecorder) IsLimitExhausted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLimitExhausted", reflect.TypeOf((*MockReader)(nil).IsLimitExhausted))
}

// BytesLeftUntilLimit mocks base method
func (m *MockReader) BytesLeftUntilLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BytesLeftUntilLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// BytesLeftUntilLimit indicates an expected call of BytesLeftUntilLimit
func (mr *MockReaderMockRecorder) BytesLeftUntilLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BytesLeftUntilLimit", reflect.TypeOf((*MockReader)(nil).BytesLeftUntilLimit))
}

// SetTransferSyntax mocks base method
func (m *MockReader) SetTransferSyntax(bo binary.ByteOrder, implicit bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransferSyntax", bo, implicit)
}

// SetTransferSyntax indicates an expected call of SetTransferSyntax
func (mr *MockReaderMockRecorder) SetTransferSyntax(bo, implicit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransferSyntax", reflect.TypeOf((*MockReader)(nil).SetTransferSyntax), bo, implicit)
}

// IsImplicit mocks base method
func (m *MockReader) IsImplicit() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsImplicit")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsImplicit indicates an expected call of IsImplicit
func (mr *MockReaderMockRecorder) IsImplicit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsImplicit", reflect.TypeOf((*MockReader)(nil).IsImplicit))
}

// SetCodingSystem mocks base method
func (m *MockReader) SetCodingSystem(cs charset.CodingSystem) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCodingSystem", cs)
}

// SetCodingSystem indicates an expected call of SetCodingSystem
func (mr *MockReaderMockRecorder) SetCodingSystem(cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCodingSystem", reflect.TypeOf((*MockReader)(nil).SetCodingSystem), cs)
}

// ByteOrder indicates an expected call of GetByteOrder
func (mr *MockReaderMockRecorder) ByteOrder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByteOrder", reflect.TypeOf((*MockReader)(nil).ByteOrder))
}

// ByteOrder mocks base method
func (m *MockReader) ByteOrder() binary.ByteOrder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByteOrder")
	ret0, _ := ret[0].(binary.ByteOrder)
	return ret0
}
