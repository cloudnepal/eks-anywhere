// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/workflows/interfaces (interfaces: Bootstrapper,ClusterManager,GitOpsManager,Validator,CAPIManager,EksdInstaller,EksdUpgrader,PackageInstaller)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	bootstrapper "github.com/aws/eks-anywhere/pkg/bootstrapper"
	cluster "github.com/aws/eks-anywhere/pkg/cluster"
	constants "github.com/aws/eks-anywhere/pkg/constants"
	providers "github.com/aws/eks-anywhere/pkg/providers"
	types "github.com/aws/eks-anywhere/pkg/types"
	validations "github.com/aws/eks-anywhere/pkg/validations"
	gomock "github.com/golang/mock/gomock"
)

// MockBootstrapper is a mock of Bootstrapper interface.
type MockBootstrapper struct {
	ctrl     *gomock.Controller
	recorder *MockBootstrapperMockRecorder
}

// MockBootstrapperMockRecorder is the mock recorder for MockBootstrapper.
type MockBootstrapperMockRecorder struct {
	mock *MockBootstrapper
}

// NewMockBootstrapper creates a new mock instance.
func NewMockBootstrapper(ctrl *gomock.Controller) *MockBootstrapper {
	mock := &MockBootstrapper{ctrl: ctrl}
	mock.recorder = &MockBootstrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBootstrapper) EXPECT() *MockBootstrapperMockRecorder {
	return m.recorder
}

// CreateBootstrapCluster mocks base method.
func (m *MockBootstrapper) CreateBootstrapCluster(arg0 context.Context, arg1 *cluster.Spec, arg2 ...bootstrapper.BootstrapClusterOption) (*types.Cluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBootstrapCluster", varargs...)
	ret0, _ := ret[0].(*types.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBootstrapCluster indicates an expected call of CreateBootstrapCluster.
func (mr *MockBootstrapperMockRecorder) CreateBootstrapCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBootstrapCluster", reflect.TypeOf((*MockBootstrapper)(nil).CreateBootstrapCluster), varargs...)
}

// DeleteBootstrapCluster mocks base method.
func (m *MockBootstrapper) DeleteBootstrapCluster(arg0 context.Context, arg1 *types.Cluster, arg2 constants.Operation, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBootstrapCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBootstrapCluster indicates an expected call of DeleteBootstrapCluster.
func (mr *MockBootstrapperMockRecorder) DeleteBootstrapCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBootstrapCluster", reflect.TypeOf((*MockBootstrapper)(nil).DeleteBootstrapCluster), arg0, arg1, arg2, arg3)
}

// MockClusterManager is a mock of ClusterManager interface.
type MockClusterManager struct {
	ctrl     *gomock.Controller
	recorder *MockClusterManagerMockRecorder
}

// MockClusterManagerMockRecorder is the mock recorder for MockClusterManager.
type MockClusterManagerMockRecorder struct {
	mock *MockClusterManager
}

// NewMockClusterManager creates a new mock instance.
func NewMockClusterManager(ctrl *gomock.Controller) *MockClusterManager {
	mock := &MockClusterManager{ctrl: ctrl}
	mock.recorder = &MockClusterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterManager) EXPECT() *MockClusterManagerMockRecorder {
	return m.recorder
}

// ApplyBundles mocks base method.
func (m *MockClusterManager) ApplyBundles(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyBundles", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyBundles indicates an expected call of ApplyBundles.
func (mr *MockClusterManagerMockRecorder) ApplyBundles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyBundles", reflect.TypeOf((*MockClusterManager)(nil).ApplyBundles), arg0, arg1, arg2)
}

// CreateAwsIamAuthCaSecret mocks base method.
func (m *MockClusterManager) CreateAwsIamAuthCaSecret(arg0 context.Context, arg1 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAwsIamAuthCaSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAwsIamAuthCaSecret indicates an expected call of CreateAwsIamAuthCaSecret.
func (mr *MockClusterManagerMockRecorder) CreateAwsIamAuthCaSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAwsIamAuthCaSecret", reflect.TypeOf((*MockClusterManager)(nil).CreateAwsIamAuthCaSecret), arg0, arg1)
}

// CreateEKSANamespace mocks base method.
func (m *MockClusterManager) CreateEKSANamespace(arg0 context.Context, arg1 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEKSANamespace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEKSANamespace indicates an expected call of CreateEKSANamespace.
func (mr *MockClusterManagerMockRecorder) CreateEKSANamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEKSANamespace", reflect.TypeOf((*MockClusterManager)(nil).CreateEKSANamespace), arg0, arg1)
}

// CreateEKSAResources mocks base method.
func (m *MockClusterManager) CreateEKSAResources(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.DatacenterConfig, arg4 []providers.MachineConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEKSAResources", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEKSAResources indicates an expected call of CreateEKSAResources.
func (mr *MockClusterManagerMockRecorder) CreateEKSAResources(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEKSAResources", reflect.TypeOf((*MockClusterManager)(nil).CreateEKSAResources), arg0, arg1, arg2, arg3, arg4)
}

// CreateWorkloadCluster mocks base method.
func (m *MockClusterManager) CreateWorkloadCluster(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.Provider) (*types.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkloadCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkloadCluster indicates an expected call of CreateWorkloadCluster.
func (mr *MockClusterManagerMockRecorder) CreateWorkloadCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkloadCluster", reflect.TypeOf((*MockClusterManager)(nil).CreateWorkloadCluster), arg0, arg1, arg2, arg3)
}

// DeleteCluster mocks base method.
func (m *MockClusterManager) DeleteCluster(arg0 context.Context, arg1, arg2 *types.Cluster, arg3 providers.Provider, arg4 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockClusterManagerMockRecorder) DeleteCluster(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterManager)(nil).DeleteCluster), arg0, arg1, arg2, arg3, arg4)
}

// EKSAClusterSpecChanged mocks base method.
func (m *MockClusterManager) EKSAClusterSpecChanged(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EKSAClusterSpecChanged", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EKSAClusterSpecChanged indicates an expected call of EKSAClusterSpecChanged.
func (mr *MockClusterManagerMockRecorder) EKSAClusterSpecChanged(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EKSAClusterSpecChanged", reflect.TypeOf((*MockClusterManager)(nil).EKSAClusterSpecChanged), arg0, arg1, arg2)
}

// GetCurrentClusterSpec mocks base method.
func (m *MockClusterManager) GetCurrentClusterSpec(arg0 context.Context, arg1 *types.Cluster, arg2 string) (*cluster.Spec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentClusterSpec", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cluster.Spec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentClusterSpec indicates an expected call of GetCurrentClusterSpec.
func (mr *MockClusterManagerMockRecorder) GetCurrentClusterSpec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentClusterSpec", reflect.TypeOf((*MockClusterManager)(nil).GetCurrentClusterSpec), arg0, arg1, arg2)
}

// InstallAwsIamAuth mocks base method.
func (m *MockClusterManager) InstallAwsIamAuth(arg0 context.Context, arg1, arg2 *types.Cluster, arg3 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallAwsIamAuth", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallAwsIamAuth indicates an expected call of InstallAwsIamAuth.
func (mr *MockClusterManagerMockRecorder) InstallAwsIamAuth(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallAwsIamAuth", reflect.TypeOf((*MockClusterManager)(nil).InstallAwsIamAuth), arg0, arg1, arg2, arg3)
}

// InstallCAPI mocks base method.
func (m *MockClusterManager) InstallCAPI(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCAPI", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallCAPI indicates an expected call of InstallCAPI.
func (mr *MockClusterManagerMockRecorder) InstallCAPI(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCAPI", reflect.TypeOf((*MockClusterManager)(nil).InstallCAPI), arg0, arg1, arg2, arg3)
}

// InstallCustomComponents mocks base method.
func (m *MockClusterManager) InstallCustomComponents(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCustomComponents", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallCustomComponents indicates an expected call of InstallCustomComponents.
func (mr *MockClusterManagerMockRecorder) InstallCustomComponents(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCustomComponents", reflect.TypeOf((*MockClusterManager)(nil).InstallCustomComponents), arg0, arg1, arg2, arg3)
}

// InstallMachineHealthChecks mocks base method.
func (m *MockClusterManager) InstallMachineHealthChecks(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallMachineHealthChecks", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallMachineHealthChecks indicates an expected call of InstallMachineHealthChecks.
func (mr *MockClusterManagerMockRecorder) InstallMachineHealthChecks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallMachineHealthChecks", reflect.TypeOf((*MockClusterManager)(nil).InstallMachineHealthChecks), arg0, arg1, arg2)
}

// InstallNetworking mocks base method.
func (m *MockClusterManager) InstallNetworking(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallNetworking", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallNetworking indicates an expected call of InstallNetworking.
func (mr *MockClusterManagerMockRecorder) InstallNetworking(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallNetworking", reflect.TypeOf((*MockClusterManager)(nil).InstallNetworking), arg0, arg1, arg2, arg3)
}

// InstallStorageClass mocks base method.
func (m *MockClusterManager) InstallStorageClass(arg0 context.Context, arg1 *types.Cluster, arg2 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallStorageClass", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallStorageClass indicates an expected call of InstallStorageClass.
func (mr *MockClusterManagerMockRecorder) InstallStorageClass(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallStorageClass", reflect.TypeOf((*MockClusterManager)(nil).InstallStorageClass), arg0, arg1, arg2)
}

// MoveCAPI mocks base method.
func (m *MockClusterManager) MoveCAPI(arg0 context.Context, arg1, arg2 *types.Cluster, arg3 string, arg4 *cluster.Spec, arg5 ...types.NodeReadyChecker) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MoveCAPI", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveCAPI indicates an expected call of MoveCAPI.
func (mr *MockClusterManagerMockRecorder) MoveCAPI(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveCAPI", reflect.TypeOf((*MockClusterManager)(nil).MoveCAPI), varargs...)
}

// PauseEKSAControllerReconcile mocks base method.
func (m *MockClusterManager) PauseEKSAControllerReconcile(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseEKSAControllerReconcile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PauseEKSAControllerReconcile indicates an expected call of PauseEKSAControllerReconcile.
func (mr *MockClusterManagerMockRecorder) PauseEKSAControllerReconcile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseEKSAControllerReconcile", reflect.TypeOf((*MockClusterManager)(nil).PauseEKSAControllerReconcile), arg0, arg1, arg2, arg3)
}

// ResumeEKSAControllerReconcile mocks base method.
func (m *MockClusterManager) ResumeEKSAControllerReconcile(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeEKSAControllerReconcile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeEKSAControllerReconcile indicates an expected call of ResumeEKSAControllerReconcile.
func (mr *MockClusterManagerMockRecorder) ResumeEKSAControllerReconcile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeEKSAControllerReconcile", reflect.TypeOf((*MockClusterManager)(nil).ResumeEKSAControllerReconcile), arg0, arg1, arg2, arg3)
}

// RunPostCreateWorkloadCluster mocks base method.
func (m *MockClusterManager) RunPostCreateWorkloadCluster(arg0 context.Context, arg1, arg2 *types.Cluster, arg3 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunPostCreateWorkloadCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunPostCreateWorkloadCluster indicates an expected call of RunPostCreateWorkloadCluster.
func (mr *MockClusterManagerMockRecorder) RunPostCreateWorkloadCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPostCreateWorkloadCluster", reflect.TypeOf((*MockClusterManager)(nil).RunPostCreateWorkloadCluster), arg0, arg1, arg2, arg3)
}

// SaveLogsManagementCluster mocks base method.
func (m *MockClusterManager) SaveLogsManagementCluster(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLogsManagementCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLogsManagementCluster indicates an expected call of SaveLogsManagementCluster.
func (mr *MockClusterManagerMockRecorder) SaveLogsManagementCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLogsManagementCluster", reflect.TypeOf((*MockClusterManager)(nil).SaveLogsManagementCluster), arg0, arg1, arg2)
}

// SaveLogsWorkloadCluster mocks base method.
func (m *MockClusterManager) SaveLogsWorkloadCluster(arg0 context.Context, arg1 providers.Provider, arg2 *cluster.Spec, arg3 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLogsWorkloadCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLogsWorkloadCluster indicates an expected call of SaveLogsWorkloadCluster.
func (mr *MockClusterManagerMockRecorder) SaveLogsWorkloadCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLogsWorkloadCluster", reflect.TypeOf((*MockClusterManager)(nil).SaveLogsWorkloadCluster), arg0, arg1, arg2, arg3)
}

// Upgrade mocks base method.
func (m *MockClusterManager) Upgrade(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec) (*types.ChangeDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.ChangeDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockClusterManagerMockRecorder) Upgrade(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockClusterManager)(nil).Upgrade), arg0, arg1, arg2, arg3)
}

// UpgradeCluster mocks base method.
func (m *MockClusterManager) UpgradeCluster(arg0 context.Context, arg1, arg2 *types.Cluster, arg3 *cluster.Spec, arg4 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeCluster", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeCluster indicates an expected call of UpgradeCluster.
func (mr *MockClusterManagerMockRecorder) UpgradeCluster(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeCluster", reflect.TypeOf((*MockClusterManager)(nil).UpgradeCluster), arg0, arg1, arg2, arg3, arg4)
}

// UpgradeNetworking mocks base method.
func (m *MockClusterManager) UpgradeNetworking(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec, arg4 providers.Provider) (*types.ChangeDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeNetworking", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.ChangeDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeNetworking indicates an expected call of UpgradeNetworking.
func (mr *MockClusterManagerMockRecorder) UpgradeNetworking(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeNetworking", reflect.TypeOf((*MockClusterManager)(nil).UpgradeNetworking), arg0, arg1, arg2, arg3, arg4)
}

// MockGitOpsManager is a mock of GitOpsManager interface.
type MockGitOpsManager struct {
	ctrl     *gomock.Controller
	recorder *MockGitOpsManagerMockRecorder
}

// MockGitOpsManagerMockRecorder is the mock recorder for MockGitOpsManager.
type MockGitOpsManagerMockRecorder struct {
	mock *MockGitOpsManager
}

// NewMockGitOpsManager creates a new mock instance.
func NewMockGitOpsManager(ctrl *gomock.Controller) *MockGitOpsManager {
	mock := &MockGitOpsManager{ctrl: ctrl}
	mock.recorder = &MockGitOpsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitOpsManager) EXPECT() *MockGitOpsManagerMockRecorder {
	return m.recorder
}

// CleanupGitRepo mocks base method.
func (m *MockGitOpsManager) CleanupGitRepo(arg0 context.Context, arg1 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupGitRepo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupGitRepo indicates an expected call of CleanupGitRepo.
func (mr *MockGitOpsManagerMockRecorder) CleanupGitRepo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupGitRepo", reflect.TypeOf((*MockGitOpsManager)(nil).CleanupGitRepo), arg0, arg1)
}

// ForceReconcileGitRepo mocks base method.
func (m *MockGitOpsManager) ForceReconcileGitRepo(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceReconcileGitRepo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceReconcileGitRepo indicates an expected call of ForceReconcileGitRepo.
func (mr *MockGitOpsManagerMockRecorder) ForceReconcileGitRepo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceReconcileGitRepo", reflect.TypeOf((*MockGitOpsManager)(nil).ForceReconcileGitRepo), arg0, arg1, arg2)
}

// Install mocks base method.
func (m *MockGitOpsManager) Install(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockGitOpsManagerMockRecorder) Install(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockGitOpsManager)(nil).Install), arg0, arg1, arg2, arg3)
}

// InstallGitOps mocks base method.
func (m *MockGitOpsManager) InstallGitOps(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.DatacenterConfig, arg4 []providers.MachineConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallGitOps", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallGitOps indicates an expected call of InstallGitOps.
func (mr *MockGitOpsManagerMockRecorder) InstallGitOps(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallGitOps", reflect.TypeOf((*MockGitOpsManager)(nil).InstallGitOps), arg0, arg1, arg2, arg3, arg4)
}

// PauseClusterResourcesReconcile mocks base method.
func (m *MockGitOpsManager) PauseClusterResourcesReconcile(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseClusterResourcesReconcile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PauseClusterResourcesReconcile indicates an expected call of PauseClusterResourcesReconcile.
func (mr *MockGitOpsManagerMockRecorder) PauseClusterResourcesReconcile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseClusterResourcesReconcile", reflect.TypeOf((*MockGitOpsManager)(nil).PauseClusterResourcesReconcile), arg0, arg1, arg2, arg3)
}

// ResumeClusterResourcesReconcile mocks base method.
func (m *MockGitOpsManager) ResumeClusterResourcesReconcile(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeClusterResourcesReconcile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeClusterResourcesReconcile indicates an expected call of ResumeClusterResourcesReconcile.
func (mr *MockGitOpsManagerMockRecorder) ResumeClusterResourcesReconcile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeClusterResourcesReconcile", reflect.TypeOf((*MockGitOpsManager)(nil).ResumeClusterResourcesReconcile), arg0, arg1, arg2, arg3)
}

// UpdateGitEksaSpec mocks base method.
func (m *MockGitOpsManager) UpdateGitEksaSpec(arg0 context.Context, arg1 *cluster.Spec, arg2 providers.DatacenterConfig, arg3 []providers.MachineConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGitEksaSpec", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGitEksaSpec indicates an expected call of UpdateGitEksaSpec.
func (mr *MockGitOpsManagerMockRecorder) UpdateGitEksaSpec(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGitEksaSpec", reflect.TypeOf((*MockGitOpsManager)(nil).UpdateGitEksaSpec), arg0, arg1, arg2, arg3)
}

// Upgrade mocks base method.
func (m *MockGitOpsManager) Upgrade(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec) (*types.ChangeDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.ChangeDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockGitOpsManagerMockRecorder) Upgrade(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockGitOpsManager)(nil).Upgrade), arg0, arg1, arg2, arg3)
}

// Validations mocks base method.
func (m *MockGitOpsManager) Validations(arg0 context.Context, arg1 *cluster.Spec) []validations.Validation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validations", arg0, arg1)
	ret0, _ := ret[0].([]validations.Validation)
	return ret0
}

// Validations indicates an expected call of Validations.
func (mr *MockGitOpsManagerMockRecorder) Validations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validations", reflect.TypeOf((*MockGitOpsManager)(nil).Validations), arg0, arg1)
}

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// PreflightValidations mocks base method.
func (m *MockValidator) PreflightValidations(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreflightValidations", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreflightValidations indicates an expected call of PreflightValidations.
func (mr *MockValidatorMockRecorder) PreflightValidations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreflightValidations", reflect.TypeOf((*MockValidator)(nil).PreflightValidations), arg0)
}

// MockCAPIManager is a mock of CAPIManager interface.
type MockCAPIManager struct {
	ctrl     *gomock.Controller
	recorder *MockCAPIManagerMockRecorder
}

// MockCAPIManagerMockRecorder is the mock recorder for MockCAPIManager.
type MockCAPIManagerMockRecorder struct {
	mock *MockCAPIManager
}

// NewMockCAPIManager creates a new mock instance.
func NewMockCAPIManager(ctrl *gomock.Controller) *MockCAPIManager {
	mock := &MockCAPIManager{ctrl: ctrl}
	mock.recorder = &MockCAPIManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAPIManager) EXPECT() *MockCAPIManagerMockRecorder {
	return m.recorder
}

// EnsureEtcdProvidersInstallation mocks base method.
func (m *MockCAPIManager) EnsureEtcdProvidersInstallation(arg0 context.Context, arg1 *types.Cluster, arg2 providers.Provider, arg3 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureEtcdProvidersInstallation", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureEtcdProvidersInstallation indicates an expected call of EnsureEtcdProvidersInstallation.
func (mr *MockCAPIManagerMockRecorder) EnsureEtcdProvidersInstallation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureEtcdProvidersInstallation", reflect.TypeOf((*MockCAPIManager)(nil).EnsureEtcdProvidersInstallation), arg0, arg1, arg2, arg3)
}

// Upgrade mocks base method.
func (m *MockCAPIManager) Upgrade(arg0 context.Context, arg1 *types.Cluster, arg2 providers.Provider, arg3, arg4 *cluster.Spec) (*types.ChangeDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.ChangeDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockCAPIManagerMockRecorder) Upgrade(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockCAPIManager)(nil).Upgrade), arg0, arg1, arg2, arg3, arg4)
}

// MockEksdInstaller is a mock of EksdInstaller interface.
type MockEksdInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockEksdInstallerMockRecorder
}

// MockEksdInstallerMockRecorder is the mock recorder for MockEksdInstaller.
type MockEksdInstallerMockRecorder struct {
	mock *MockEksdInstaller
}

// NewMockEksdInstaller creates a new mock instance.
func NewMockEksdInstaller(ctrl *gomock.Controller) *MockEksdInstaller {
	mock := &MockEksdInstaller{ctrl: ctrl}
	mock.recorder = &MockEksdInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEksdInstaller) EXPECT() *MockEksdInstallerMockRecorder {
	return m.recorder
}

// InstallEksdCRDs mocks base method.
func (m *MockEksdInstaller) InstallEksdCRDs(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallEksdCRDs", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallEksdCRDs indicates an expected call of InstallEksdCRDs.
func (mr *MockEksdInstallerMockRecorder) InstallEksdCRDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallEksdCRDs", reflect.TypeOf((*MockEksdInstaller)(nil).InstallEksdCRDs), arg0, arg1, arg2)
}

// InstallEksdManifest mocks base method.
func (m *MockEksdInstaller) InstallEksdManifest(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallEksdManifest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallEksdManifest indicates an expected call of InstallEksdManifest.
func (mr *MockEksdInstallerMockRecorder) InstallEksdManifest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallEksdManifest", reflect.TypeOf((*MockEksdInstaller)(nil).InstallEksdManifest), arg0, arg1, arg2)
}

// MockEksdUpgrader is a mock of EksdUpgrader interface.
type MockEksdUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockEksdUpgraderMockRecorder
}

// MockEksdUpgraderMockRecorder is the mock recorder for MockEksdUpgrader.
type MockEksdUpgraderMockRecorder struct {
	mock *MockEksdUpgrader
}

// NewMockEksdUpgrader creates a new mock instance.
func NewMockEksdUpgrader(ctrl *gomock.Controller) *MockEksdUpgrader {
	mock := &MockEksdUpgrader{ctrl: ctrl}
	mock.recorder = &MockEksdUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEksdUpgrader) EXPECT() *MockEksdUpgraderMockRecorder {
	return m.recorder
}

// Upgrade mocks base method.
func (m *MockEksdUpgrader) Upgrade(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec) (*types.ChangeDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.ChangeDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockEksdUpgraderMockRecorder) Upgrade(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockEksdUpgrader)(nil).Upgrade), arg0, arg1, arg2, arg3)
}

// MockPackageInstaller is a mock of PackageInstaller interface.
type MockPackageInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockPackageInstallerMockRecorder
}

// MockPackageInstallerMockRecorder is the mock recorder for MockPackageInstaller.
type MockPackageInstallerMockRecorder struct {
	mock *MockPackageInstaller
}

// NewMockPackageInstaller creates a new mock instance.
func NewMockPackageInstaller(ctrl *gomock.Controller) *MockPackageInstaller {
	mock := &MockPackageInstaller{ctrl: ctrl}
	mock.recorder = &MockPackageInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageInstaller) EXPECT() *MockPackageInstallerMockRecorder {
	return m.recorder
}

// InstallCuratedPackages mocks base method.
func (m *MockPackageInstaller) InstallCuratedPackages(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCuratedPackages", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallCuratedPackages indicates an expected call of InstallCuratedPackages.
func (mr *MockPackageInstallerMockRecorder) InstallCuratedPackages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCuratedPackages", reflect.TypeOf((*MockPackageInstaller)(nil).InstallCuratedPackages), arg0)
}
