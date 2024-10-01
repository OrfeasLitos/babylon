// Code generated by MockGen. DO NOT EDIT.
// Source: x/btcstaking/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	types "github.com/babylonlabs-io/babylon/types"
	types0 "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	types1 "github.com/babylonlabs-io/babylon/x/btclightclient/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBTCLightClientKeeper is a mock of BTCLightClientKeeper interface.
type MockBTCLightClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCLightClientKeeperMockRecorder
}

// MockBTCLightClientKeeperMockRecorder is the mock recorder for MockBTCLightClientKeeper.
type MockBTCLightClientKeeperMockRecorder struct {
	mock *MockBTCLightClientKeeper
}

// NewMockBTCLightClientKeeper creates a new mock instance.
func NewMockBTCLightClientKeeper(ctrl *gomock.Controller) *MockBTCLightClientKeeper {
	mock := &MockBTCLightClientKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCLightClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCLightClientKeeper) EXPECT() *MockBTCLightClientKeeperMockRecorder {
	return m.recorder
}

// GetBaseBTCHeader mocks base method.
func (m *MockBTCLightClientKeeper) GetBaseBTCHeader(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseBTCHeader", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetBaseBTCHeader indicates an expected call of GetBaseBTCHeader.
func (mr *MockBTCLightClientKeeperMockRecorder) GetBaseBTCHeader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseBTCHeader", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetBaseBTCHeader), ctx)
}

// GetHeaderByHash mocks base method.
func (m *MockBTCLightClientKeeper) GetHeaderByHash(ctx context.Context, hash *types.BTCHeaderHashBytes) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash.
func (mr *MockBTCLightClientKeeperMockRecorder) GetHeaderByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetHeaderByHash), ctx, hash)
}

// GetTipInfo mocks base method.
func (m *MockBTCLightClientKeeper) GetTipInfo(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTipInfo", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetTipInfo indicates an expected call of GetTipInfo.
func (mr *MockBTCLightClientKeeperMockRecorder) GetTipInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTipInfo", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetTipInfo), ctx)
}

// MockBtcCheckpointKeeper is a mock of BtcCheckpointKeeper interface.
type MockBtcCheckpointKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBtcCheckpointKeeperMockRecorder
}

// MockBtcCheckpointKeeperMockRecorder is the mock recorder for MockBtcCheckpointKeeper.
type MockBtcCheckpointKeeperMockRecorder struct {
	mock *MockBtcCheckpointKeeper
}

// NewMockBtcCheckpointKeeper creates a new mock instance.
func NewMockBtcCheckpointKeeper(ctrl *gomock.Controller) *MockBtcCheckpointKeeper {
	mock := &MockBtcCheckpointKeeper{ctrl: ctrl}
	mock.recorder = &MockBtcCheckpointKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcCheckpointKeeper) EXPECT() *MockBtcCheckpointKeeperMockRecorder {
	return m.recorder
}

// GetParams mocks base method.
func (m *MockBtcCheckpointKeeper) GetParams(ctx context.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetParams), ctx)
}

// MockFinalityKeeper is a mock of FinalityKeeper interface.
type MockFinalityKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFinalityKeeperMockRecorder
}

// MockFinalityKeeperMockRecorder is the mock recorder for MockFinalityKeeper.
type MockFinalityKeeperMockRecorder struct {
	mock *MockFinalityKeeper
}

// NewMockFinalityKeeper creates a new mock instance.
func NewMockFinalityKeeper(ctrl *gomock.Controller) *MockFinalityKeeper {
	mock := &MockFinalityKeeper{ctrl: ctrl}
	mock.recorder = &MockFinalityKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinalityKeeper) EXPECT() *MockFinalityKeeperMockRecorder {
	return m.recorder
}

// HasTimestampedPubRand mocks base method.
func (m *MockFinalityKeeper) HasTimestampedPubRand(ctx context.Context, fpBtcPK *types.BIP340PubKey, height uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTimestampedPubRand", ctx, fpBtcPK, height)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTimestampedPubRand indicates an expected call of HasTimestampedPubRand.
func (mr *MockFinalityKeeperMockRecorder) HasTimestampedPubRand(ctx, fpBtcPK, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTimestampedPubRand", reflect.TypeOf((*MockFinalityKeeper)(nil).HasTimestampedPubRand), ctx, fpBtcPK, height)
}

// MockBtcStakingHooks is a mock of BtcStakingHooks interface.
type MockBtcStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockBtcStakingHooksMockRecorder
}

// MockBtcStakingHooksMockRecorder is the mock recorder for MockBtcStakingHooks.
type MockBtcStakingHooksMockRecorder struct {
	mock *MockBtcStakingHooks
}

// NewMockBtcStakingHooks creates a new mock instance.
func NewMockBtcStakingHooks(ctrl *gomock.Controller) *MockBtcStakingHooks {
	mock := &MockBtcStakingHooks{ctrl: ctrl}
	mock.recorder = &MockBtcStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcStakingHooks) EXPECT() *MockBtcStakingHooksMockRecorder {
	return m.recorder
}

// AfterFinalityProviderActivated mocks base method.
func (m *MockBtcStakingHooks) AfterFinalityProviderActivated(ctx context.Context, fpPk *types.BIP340PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFinalityProviderActivated", ctx, fpPk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterFinalityProviderActivated indicates an expected call of AfterFinalityProviderActivated.
func (mr *MockBtcStakingHooksMockRecorder) AfterFinalityProviderActivated(ctx, fpPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFinalityProviderActivated", reflect.TypeOf((*MockBtcStakingHooks)(nil).AfterFinalityProviderActivated), ctx, fpPk)
}