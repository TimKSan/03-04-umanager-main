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

// MockLinksClient is a mock of LinkServiceClient interface
type MockLinksClient struct {
	ctrl     *gomock.Controller
	recorder *MockLinksClientMockRecorder
}

// MockLinksClientMockRecorder is the mock recorder for MockLinksClient
type MockLinksClientMockRecorder struct {
	mock *MockLinksClient
}

// NewMockLinksClient creates a new mock instance
func NewMockLinksClient(ctrl *gomock.Controller) *MockLinksClient {
	mock := &MockLinksClient{ctrl: ctrl}
	mock.recorder = &MockLinksClientMockRecorder{mock}
	return mock
}

func (m *MockLinksClient) EXPECT() *MockLinksClientMockRecorder {
	return m.recorder
}

func (m *MockLinksClient) CreateLink(ctx context.Context, in *pb.CreateLinkRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLink", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) CreateLink(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLink", reflect.TypeOf((*MockLinksClient)(nil).CreateLink), ctx, in)
}

func (m *MockLinksClient) DeleteLink(ctx context.Context, in *pb.DeleteLinkRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLink", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) DeleteLink(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLink", reflect.TypeOf((*MockLinksClient)(nil).DeleteLink), ctx, in)
}

// Add the missing methods
func (m *MockLinksClient) GetLink(ctx context.Context, in *pb.GetLinkRequest, opts ...grpc.CallOption) (*pb.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink", ctx, in)
	ret0, _ := ret[0].(*pb.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) GetLink(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*MockLinksClient)(nil).GetLink), ctx, in)
}

func (m *MockLinksClient) GetLinkByUserID(ctx context.Context, in *pb.GetLinksByUserId, opts ...grpc.CallOption) (*pb.ListLinkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLinkByUserID", ctx, in)
	ret0, _ := ret[0].(*pb.ListLinkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) GetLinkByUserID(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLinkByUserID", reflect.TypeOf((*MockLinksClient)(nil).GetLinkByUserID), ctx, in)
}

func (m *MockLinksClient) ListLinks(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.ListLinkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLinks", ctx, in)
	ret0, _ := ret[0].(*pb.ListLinkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) ListLinks(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinks", reflect.TypeOf((*MockLinksClient)(nil).ListLinks), ctx, in)
}

func (m *MockLinksClient) UpdateLink(ctx context.Context, in *pb.UpdateLinkRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLink", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) UpdateLink(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLink", reflect.TypeOf((*MockLinksClient)(nil).UpdateLink), ctx, in)
}

func TestPostLinks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockLinksClient(ctrl)
	handler := newLinksHandler(mockClient)

	link := apiv1.LinkCreate{
		Id:     "1",
		Title:  "Test Title",
		Url:    "http://example.com",
		Images: []string{"image1", "image2"},
		Tags:   []string{"tag1", "tag2"},
		UserId: "user1",
	}

	mockClient.EXPECT().CreateLink(gomock.Any(), &pb.CreateLinkRequest{
		Id:     link.Id,
		Title:  link.Title,
		Url:    link.Url,
		Images: link.Images,
		Tags:   link.Tags,
		UserId: link.UserId,
	}).Return(&pb.Empty{}, nil)

	body, _ := json.Marshal(link)
	req, _ := http.NewRequest("POST", "/links", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.PostLinks(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestDeleteLinksId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockLinksClient(ctrl)
	handler := newLinksHandler(mockClient)

	linkID := "1"

	mockClient.EXPECT().DeleteLink(gomock.Any(), &pb.DeleteLinkRequest{Id: linkID}).Return(&pb.Empty{}, nil)

	req, _ := http.NewRequest("DELETE", "/links/"+linkID, nil)
	rr := httptest.NewRecorder()

	handler.DeleteLinksId(rr, req, linkID)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
