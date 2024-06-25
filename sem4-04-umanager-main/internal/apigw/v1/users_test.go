package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/robotomize/gb-golang/homework/03-04-umanager/pkg/api/apiv1"
	"gitlab.com/robotomize/gb-golang/homework/03-04-umanager/pkg/pb"
	"google.golang.org/grpc"
)

// MockUsersClient is a mock of UserServiceClient interface
type MockUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersClientMockRecorder
}

// MockUsersClientMockRecorder is the mock recorder for MockUsersClient
type MockUsersClientMockRecorder struct {
	mock *MockUsersClient
}

// NewMockUsersClient creates a new mock instance
func NewMockUsersClient(ctrl *gomock.Controller) *MockUsersClient {
	mock := &MockUsersClient{ctrl: ctrl}
	mock.recorder = &MockUsersClientMockRecorder{mock}
	return mock
}

func (m *MockUsersClient) EXPECT() *MockUsersClientMockRecorder {
	return m.recorder
}

func (m *MockUsersClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsersClient)(nil).CreateUser), ctx, in)
}

func (m *MockUsersClient) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) DeleteUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUsersClient)(nil).DeleteUser), ctx, in)
}

// Add the missing methods
func (m *MockUsersClient) GetUser(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, in)
	ret0, _ := ret[0].(*pb.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) GetUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsersClient)(nil).GetUser), ctx, in)
}

func (m *MockUsersClient) ListUsers(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.ListUsersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, in)
	ret0, _ := ret[0].(*pb.ListUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) ListUsers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUsersClient)(nil).ListUsers), ctx, in)
}

func (m *MockUsersClient) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) UpdateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsersClient)(nil).UpdateUser), ctx, in)
}

func TestPostUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockUsersClient(ctrl)
	handler := newUsersHandler(mockClient)

	user := apiv1.UserCreate{
		Id:       "1",
		Username: "testuser",
		Password: "password123",
	}

	mockClient.EXPECT().CreateUser(gomock.Any(), &pb.CreateUserRequest{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}).Return(&pb.Empty{}, nil)

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.PostUsers(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestDeleteUsersId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockUsersClient(ctrl)
	handler := newUsersHandler(mockClient)

	userID := "1"

	mockClient.EXPECT().DeleteUser(gomock.Any(), &pb.DeleteUserRequest{Id: userID}).Return(&pb.Empty{}, nil)

	req, _ := http.NewRequest("DELETE", "/users/"+userID, nil)
	rr := httptest.NewRecorder()

	handler.DeleteUsersId(rr, req, userID)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
