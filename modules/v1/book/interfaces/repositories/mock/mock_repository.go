package mock_repository

import (
	domain "bookLib/modules/v1/book/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryPresenter is a mock of RepositoryPresenter interface.
type MockRepositoryPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryPresenterMockRecorder
}

// MockRepositoryPresenterMockRecorder is the mock recorder for MockRepositoryPresenter.
type MockRepositoryPresenterMockRecorder struct {
	mock *MockRepositoryPresenter
}

// NewMockRepositoryPresenter creates a new mock instance.
func NewMockRepositoryPresenter(ctrl *gomock.Controller) *MockRepositoryPresenter {
	mock := &MockRepositoryPresenter{ctrl: ctrl}
	mock.recorder = &MockRepositoryPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryPresenter) EXPECT() *MockRepositoryPresenterMockRecorder {
	return m.recorder
}

// AllBooks mocks base method.
func (m *MockRepositoryPresenter) AllBooks() (domain.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllBooks")
	ret0, _ := ret[0].(domain.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllBooks indicates an expected call of AllBooks.
func (mr *MockRepositoryPresenterMockRecorder) AllBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllBooks", reflect.TypeOf((*MockRepositoryPresenter)(nil).AllBooks))
}

// CreateBook mocks base method.
func (m *MockRepositoryPresenter) CreateBook(book domain.Book) (domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", book)
	ret0, _ := ret[0].(domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockRepositoryPresenterMockRecorder) CreateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockRepositoryPresenter)(nil).CreateBook), book)
}

// DeleteBook mocks base method.
func (m *MockRepositoryPresenter) DeleteBook(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockRepositoryPresenterMockRecorder) DeleteBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockRepositoryPresenter)(nil).DeleteBook), id)
}

// GetBookByID mocks base method.
func (m *MockRepositoryPresenter) GetBookByID(id string) (domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", id)
	ret0, _ := ret[0].(domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockRepositoryPresenterMockRecorder) GetBookByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockRepositoryPresenter)(nil).GetBookByID), id)
}

// UpdateBook mocks base method.
func (m *MockRepositoryPresenter) UpdateBook(id string, book domain.Book) (domain.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", id, book)
	ret0, _ := ret[0].(domain.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockRepositoryPresenterMockRecorder) UpdateBook(id, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockRepositoryPresenter)(nil).UpdateBook), id, book)
}
