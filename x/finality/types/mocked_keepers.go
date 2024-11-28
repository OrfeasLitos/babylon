// Code generated by MockGen. DO NOT EDIT.
// Source: x/finality/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	types "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	types0 "github.com/babylonlabs-io/babylon/x/epoching/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBTCStakingKeeper is a mock of BTCStakingKeeper interface.
type MockBTCStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCStakingKeeperMockRecorder
}

// MockBTCStakingKeeperMockRecorder is the mock recorder for MockBTCStakingKeeper.
type MockBTCStakingKeeperMockRecorder struct {
	mock *MockBTCStakingKeeper
}

// NewMockBTCStakingKeeper creates a new mock instance.
func NewMockBTCStakingKeeper(ctrl *gomock.Controller) *MockBTCStakingKeeper {
	mock := &MockBTCStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCStakingKeeper) EXPECT() *MockBTCStakingKeeperMockRecorder {
	return m.recorder
}

// ClearPowerDistUpdateEvents mocks base method.
func (m *MockBTCStakingKeeper) ClearPowerDistUpdateEvents(ctx context.Context, btcHeight uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearPowerDistUpdateEvents", ctx, btcHeight)
}

// ClearPowerDistUpdateEvents indicates an expected call of ClearPowerDistUpdateEvents.
func (mr *MockBTCStakingKeeperMockRecorder) ClearPowerDistUpdateEvents(ctx, btcHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearPowerDistUpdateEvents", reflect.TypeOf((*MockBTCStakingKeeper)(nil).ClearPowerDistUpdateEvents), ctx, btcHeight)
}

// GetAllPowerDistUpdateEvents mocks base method.
func (m *MockBTCStakingKeeper) GetAllPowerDistUpdateEvents(ctx context.Context, lastBTCTipHeight, btcTipHeight uint32) []*types.EventPowerDistUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPowerDistUpdateEvents", ctx, lastBTCTipHeight, btcTipHeight)
	ret0, _ := ret[0].([]*types.EventPowerDistUpdate)
	return ret0
}

// GetAllPowerDistUpdateEvents indicates an expected call of GetAllPowerDistUpdateEvents.
func (mr *MockBTCStakingKeeperMockRecorder) GetAllPowerDistUpdateEvents(ctx, lastBTCTipHeight, btcTipHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPowerDistUpdateEvents", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetAllPowerDistUpdateEvents), ctx, lastBTCTipHeight, btcTipHeight)
}

// GetBTCDelegation mocks base method.
func (m *MockBTCStakingKeeper) GetBTCDelegation(ctx context.Context, stakingTxHashStr string) (*types.BTCDelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBTCDelegation", ctx, stakingTxHashStr)
	ret0, _ := ret[0].(*types.BTCDelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBTCDelegation indicates an expected call of GetBTCDelegation.
func (mr *MockBTCStakingKeeperMockRecorder) GetBTCDelegation(ctx, stakingTxHashStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBTCDelegation", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetBTCDelegation), ctx, stakingTxHashStr)
}

// GetBTCHeightAtBabylonHeight mocks base method.
func (m *MockBTCStakingKeeper) GetBTCHeightAtBabylonHeight(ctx context.Context, babylonHeight uint64) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBTCHeightAtBabylonHeight", ctx, babylonHeight)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetBTCHeightAtBabylonHeight indicates an expected call of GetBTCHeightAtBabylonHeight.
func (mr *MockBTCStakingKeeperMockRecorder) GetBTCHeightAtBabylonHeight(ctx, babylonHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBTCHeightAtBabylonHeight", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetBTCHeightAtBabylonHeight), ctx, babylonHeight)
}

// GetCurrentBTCHeight mocks base method.
func (m *MockBTCStakingKeeper) GetCurrentBTCHeight(ctx context.Context) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBTCHeight", ctx)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetCurrentBTCHeight indicates an expected call of GetCurrentBTCHeight.
func (mr *MockBTCStakingKeeperMockRecorder) GetCurrentBTCHeight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBTCHeight", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetCurrentBTCHeight), ctx)
}

// GetFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) GetFinalityProvider(ctx context.Context, fpBTCPK []byte) (*types.FinalityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(*types.FinalityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFinalityProvider indicates an expected call of GetFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) GetFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetFinalityProvider), ctx, fpBTCPK)
}

// GetParams mocks base method.
func (m *MockBTCStakingKeeper) GetParams(ctx context.Context) types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockBTCStakingKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockBTCStakingKeeper)(nil).GetParams), ctx)
}

// HasFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) HasFinalityProvider(ctx context.Context, fpBTCPK []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFinalityProvider indicates an expected call of HasFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) HasFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).HasFinalityProvider), ctx, fpBTCPK)
}

// JailFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) JailFinalityProvider(ctx context.Context, fpBTCPK []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JailFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// JailFinalityProvider indicates an expected call of JailFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) JailFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JailFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).JailFinalityProvider), ctx, fpBTCPK)
}

// SlashFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) SlashFinalityProvider(ctx context.Context, fpBTCPK []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// SlashFinalityProvider indicates an expected call of SlashFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) SlashFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).SlashFinalityProvider), ctx, fpBTCPK)
}

// UnjailFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) UnjailFinalityProvider(ctx context.Context, fpBTCPK []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnjailFinalityProvider", ctx, fpBTCPK)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnjailFinalityProvider indicates an expected call of UnjailFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) UnjailFinalityProvider(ctx, fpBTCPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnjailFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).UnjailFinalityProvider), ctx, fpBTCPK)
}

// UpdateFinalityProvider mocks base method.
func (m *MockBTCStakingKeeper) UpdateFinalityProvider(ctx context.Context, fp *types.FinalityProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFinalityProvider", ctx, fp)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFinalityProvider indicates an expected call of UpdateFinalityProvider.
func (mr *MockBTCStakingKeeperMockRecorder) UpdateFinalityProvider(ctx, fp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFinalityProvider", reflect.TypeOf((*MockBTCStakingKeeper)(nil).UpdateFinalityProvider), ctx, fp)
}

// MockCheckpointingKeeper is a mock of CheckpointingKeeper interface.
type MockCheckpointingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointingKeeperMockRecorder
}

// MockCheckpointingKeeperMockRecorder is the mock recorder for MockCheckpointingKeeper.
type MockCheckpointingKeeperMockRecorder struct {
	mock *MockCheckpointingKeeper
}

// NewMockCheckpointingKeeper creates a new mock instance.
func NewMockCheckpointingKeeper(ctrl *gomock.Controller) *MockCheckpointingKeeper {
	mock := &MockCheckpointingKeeper{ctrl: ctrl}
	mock.recorder = &MockCheckpointingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointingKeeper) EXPECT() *MockCheckpointingKeeperMockRecorder {
	return m.recorder
}

// GetEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetEpoch(ctx context.Context) *types0.Epoch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpoch", ctx)
	ret0, _ := ret[0].(*types0.Epoch)
	return ret0
}

// GetEpoch indicates an expected call of GetEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetEpoch), ctx)
}

// GetLastFinalizedEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetLastFinalizedEpoch(ctx context.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastFinalizedEpoch", ctx)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLastFinalizedEpoch indicates an expected call of GetLastFinalizedEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetLastFinalizedEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastFinalizedEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetLastFinalizedEpoch), ctx)
}

// MockIncentiveKeeper is a mock of IncentiveKeeper interface.
type MockIncentiveKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockIncentiveKeeperMockRecorder
}

// MockIncentiveKeeperMockRecorder is the mock recorder for MockIncentiveKeeper.
type MockIncentiveKeeperMockRecorder struct {
	mock *MockIncentiveKeeper
}

// NewMockIncentiveKeeper creates a new mock instance.
func NewMockIncentiveKeeper(ctrl *gomock.Controller) *MockIncentiveKeeper {
	mock := &MockIncentiveKeeper{ctrl: ctrl}
	mock.recorder = &MockIncentiveKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIncentiveKeeper) EXPECT() *MockIncentiveKeeperMockRecorder {
	return m.recorder
}

// IndexRefundableMsg mocks base method.
func (m *MockIncentiveKeeper) IndexRefundableMsg(ctx context.Context, msg types1.Msg) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IndexRefundableMsg", ctx, msg)
}

// IndexRefundableMsg indicates an expected call of IndexRefundableMsg.
func (mr *MockIncentiveKeeperMockRecorder) IndexRefundableMsg(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexRefundableMsg", reflect.TypeOf((*MockIncentiveKeeper)(nil).IndexRefundableMsg), ctx, msg)
}

// RewardBTCStaking mocks base method.
func (m *MockIncentiveKeeper) RewardBTCStaking(ctx context.Context, height uint64, filteredDc *VotingPowerDistCache) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RewardBTCStaking", ctx, height, filteredDc)
}

// RewardBTCStaking indicates an expected call of RewardBTCStaking.
func (mr *MockIncentiveKeeperMockRecorder) RewardBTCStaking(ctx, height, filteredDc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewardBTCStaking", reflect.TypeOf((*MockIncentiveKeeper)(nil).RewardBTCStaking), ctx, height, filteredDc)
}
