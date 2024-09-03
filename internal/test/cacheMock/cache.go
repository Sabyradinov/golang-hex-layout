package cacheMock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockICache is a mock of ICache interface.
type MockICache struct {
	ctrl     *gomock.Controller
	recorder *MockICacheMockRecorder
}

// MockICacheMockRecorder is the mock recorder for MockICache.
type MockICacheMockRecorder struct {
	mock *MockICache
}

// NewMockICache creates a new mock instance.
func NewMockICache(ctrl *gomock.Controller) *MockICache {
	mock := &MockICache{ctrl: ctrl}
	mock.recorder = &MockICacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICache) EXPECT() *MockICacheMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockICache) Del(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockICacheMockRecorder) Del(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockICache)(nil).Del), ctx, key)
}

// Exists mocks base method.
func (m *MockICache) Exists(ctx context.Context, key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockICacheMockRecorder) Exists(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockICache)(nil).Exists), ctx, key)
}

// Get mocks base method.
func (m *MockICache) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockICacheMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockICache)(nil).Get), ctx, key)
}

// RegisterMetrics mocks base method.
func (m *MockICache) RegisterMetrics() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterMetrics")
}

// RegisterMetrics indicates an expected call of RegisterMetrics.
func (mr *MockICacheMockRecorder) RegisterMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMetrics", reflect.TypeOf((*MockICache)(nil).RegisterMetrics))
}

// Set mocks base method.
func (m *MockICache) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, val, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockICacheMockRecorder) Set(ctx, key, val, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockICache)(nil).Set), ctx, key, val, ttl)
}
