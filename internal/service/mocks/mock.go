// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/olexsmir/noter/internal/domain"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Logout mocks base method.
func (m *MockUsers) Logout(userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUsersMockRecorder) Logout(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUsers)(nil).Logout), userID)
}

// RefreshTokens mocks base method.
func (m *MockUsers) RefreshTokens(refreshToken string) (domain.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", refreshToken)
	ret0, _ := ret[0].(domain.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockUsersMockRecorder) RefreshTokens(refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockUsers)(nil).RefreshTokens), refreshToken)
}

// SignIn mocks base method.
func (m *MockUsers) SignIn(input domain.UserSignIn) (domain.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", input)
	ret0, _ := ret[0].(domain.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUsersMockRecorder) SignIn(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUsers)(nil).SignIn), input)
}

// SignUp mocks base method.
func (m *MockUsers) SignUp(user domain.UserSignUp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUsersMockRecorder) SignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUsers)(nil).SignUp), user)
}

// MockNotes is a mock of Notes interface.
type MockNotes struct {
	ctrl     *gomock.Controller
	recorder *MockNotesMockRecorder
}

// MockNotesMockRecorder is the mock recorder for MockNotes.
type MockNotesMockRecorder struct {
	mock *MockNotes
}

// NewMockNotes creates a new mock instance.
func NewMockNotes(ctrl *gomock.Controller) *MockNotes {
	mock := &MockNotes{ctrl: ctrl}
	mock.recorder = &MockNotesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotes) EXPECT() *MockNotesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNotes) Create(input domain.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockNotesMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNotes)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockNotes) Delete(id, authorID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, authorID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNotesMockRecorder) Delete(id, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNotes)(nil).Delete), id, authorID)
}

// DeleteAll mocks base method.
func (m *MockNotes) DeleteAll(notebookID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", notebookID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockNotesMockRecorder) DeleteAll(notebookID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockNotes)(nil).DeleteAll), notebookID, userID)
}

// GetAll mocks base method.
func (m *MockNotes) GetAll(authorID, notebookID, page int) ([]domain.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", authorID, notebookID, page)
	ret0, _ := ret[0].([]domain.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNotesMockRecorder) GetAll(authorID, notebookID, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNotes)(nil).GetAll), authorID, notebookID, page)
}

// GetByID mocks base method.
func (m *MockNotes) GetByID(id int) (domain.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(domain.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockNotesMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockNotes)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockNotes) Update(id, authorID, notebookID int, inp domain.UpdateNoteInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, authorID, notebookID, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNotesMockRecorder) Update(id, authorID, notebookID, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNotes)(nil).Update), id, authorID, notebookID, inp)
}

// MockNotebooks is a mock of Notebooks interface.
type MockNotebooks struct {
	ctrl     *gomock.Controller
	recorder *MockNotebooksMockRecorder
}

// MockNotebooksMockRecorder is the mock recorder for MockNotebooks.
type MockNotebooksMockRecorder struct {
	mock *MockNotebooks
}

// NewMockNotebooks creates a new mock instance.
func NewMockNotebooks(ctrl *gomock.Controller) *MockNotebooks {
	mock := &MockNotebooks{ctrl: ctrl}
	mock.recorder = &MockNotebooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotebooks) EXPECT() *MockNotebooksMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNotebooks) Create(input domain.Notebook) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockNotebooksMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNotebooks)(nil).Create), input)
}

// Delete mocks base method.
func (m *MockNotebooks) Delete(id, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNotebooksMockRecorder) Delete(id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNotebooks)(nil).Delete), id, userID)
}

// GetAll mocks base method.
func (m *MockNotebooks) GetAll(userID, page int) ([]domain.Notebook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userID, page)
	ret0, _ := ret[0].([]domain.Notebook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNotebooksMockRecorder) GetAll(userID, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNotebooks)(nil).GetAll), userID, page)
}

// GetById mocks base method.
func (m *MockNotebooks) GetById(id, userID int) (domain.Notebook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id, userID)
	ret0, _ := ret[0].(domain.Notebook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockNotebooksMockRecorder) GetById(id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockNotebooks)(nil).GetById), id, userID)
}

// Update mocks base method.
func (m *MockNotebooks) Update(id, userID int, inp domain.UpdateNotebookInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, userID, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNotebooksMockRecorder) Update(id, userID, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNotebooks)(nil).Update), id, userID, inp)
}
