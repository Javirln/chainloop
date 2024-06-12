// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package testhelpers

import (
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/authz"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/conf/controlplane/config/v1"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data"
	"github.com/chainloop-dev/chainloop/app/controlplane/plugins/sdk/v1"
	"github.com/chainloop-dev/chainloop/internal/robotaccount/cas"
	"github.com/chainloop-dev/chainloop/pkg/blobmanager"
	"github.com/chainloop-dev/chainloop/pkg/credentials"
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

// wireTestData init testing data
func WireTestData(testDatabase *TestDatabase, t *testing.T, logger log.Logger, readerWriter credentials.ReaderWriter, builder *robotaccount.Builder, auth *conf.Auth, availablePlugins sdk.AvailablePlugins, providers backend.Providers) (*TestingUseCases, func(), error) {
	confData := NewConfData(testDatabase, t)
	newConfig := NewDataConfig(confData)
	dataData, cleanup, err := data.NewData(newConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	membershipRepo := data.NewMembershipRepo(dataData, logger)
	organizationRepo := data.NewOrganizationRepo(dataData, logger)
	casBackendRepo := data.NewCASBackendRepo(dataData, logger)
	casBackendUseCase := biz.NewCASBackendUseCase(casBackendRepo, readerWriter, providers, logger)
	integrationRepo := data.NewIntegrationRepo(dataData, logger)
	integrationAttachmentRepo := data.NewIntegrationAttachmentRepo(dataData, logger)
	workflowRepo := data.NewWorkflowRepo(dataData, logger)
	newIntegrationUseCaseOpts := &biz.NewIntegrationUseCaseOpts{
		IRepo:   integrationRepo,
		IaRepo:  integrationAttachmentRepo,
		WfRepo:  workflowRepo,
		CredsRW: readerWriter,
		Logger:  logger,
	}
	integrationUseCase := biz.NewIntegrationUseCase(newIntegrationUseCaseOpts)
	organizationUseCase := biz.NewOrganizationUseCase(organizationRepo, casBackendUseCase, integrationUseCase, membershipRepo, logger)
	membershipUseCase := biz.NewMembershipUseCase(membershipRepo, organizationUseCase, logger)
	workflowContractRepo := data.NewWorkflowContractRepo(dataData, logger)
	workflowContractUseCase := biz.NewWorkflowContractUseCase(workflowContractRepo, logger)
	workflowUseCase := biz.NewWorkflowUsecase(workflowRepo, workflowContractUseCase, logger)
	workflowRunRepo := data.NewWorkflowRunRepo(dataData, logger)
	workflowRunUseCase, err := biz.NewWorkflowRunUseCase(workflowRunRepo, workflowRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	newUserUseCaseParams := &biz.NewUserUseCaseParams{
		UserRepo:            userRepo,
		MembershipUseCase:   membershipUseCase,
		OrganizationUseCase: organizationUseCase,
		Logger:              logger,
	}
	userUseCase := biz.NewUserUseCase(newUserUseCaseParams)
	robotAccountRepo := data.NewRobotAccountRepo(dataData, logger)
	robotAccountUseCase := biz.NewRootAccountUseCase(robotAccountRepo, workflowRepo, auth, logger)
	casMappingRepo := data.NewCASMappingRepo(dataData, casBackendRepo, logger)
	casMappingUseCase := biz.NewCASMappingUseCase(casMappingRepo, membershipRepo, logger)
	orgInvitationRepo := data.NewOrgInvitation(dataData, logger)
	orgInvitationUseCase, err := biz.NewOrgInvitationUseCase(orgInvitationRepo, membershipRepo, userRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	referrerRepo := data.NewReferrerRepo(dataData, workflowRepo, logger)
	referrerSharedIndex := _wireReferrerSharedIndexValue
	referrerUseCase, err := biz.NewReferrerUseCase(referrerRepo, workflowRepo, membershipRepo, referrerSharedIndex, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	apiTokenRepo := data.NewAPITokenRepo(dataData, logger)
	data_Database := confData.Database
	enforcer, err := authz.NewDatabaseEnforcer(data_Database)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	apiTokenUseCase, err := biz.NewAPITokenUseCase(apiTokenRepo, auth, enforcer, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	attestationStateRepo := data.NewAttestationStateRepo(dataData, logger)
	attestationStateUseCase, err := biz.NewAttestationStateUseCase(attestationStateRepo, workflowRunRepo)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	testingRepos := &TestingRepos{
		Membership:       membershipRepo,
		Referrer:         referrerRepo,
		Workflow:         workflowRepo,
		WorkflowRunRepo:  workflowRunRepo,
		AttestationState: attestationStateRepo,
	}
	testingUseCases := &TestingUseCases{
		DB:                     testDatabase,
		Data:                   dataData,
		L:                      logger,
		Membership:             membershipUseCase,
		CASBackend:             casBackendUseCase,
		Integration:            integrationUseCase,
		Organization:           organizationUseCase,
		WorkflowContract:       workflowContractUseCase,
		Workflow:               workflowUseCase,
		WorkflowRun:            workflowRunUseCase,
		User:                   userUseCase,
		RobotAccount:           robotAccountUseCase,
		RegisteredIntegrations: availablePlugins,
		CASMapping:             casMappingUseCase,
		OrgInvitation:          orgInvitationUseCase,
		Referrer:               referrerUseCase,
		APIToken:               apiTokenUseCase,
		Enforcer:               enforcer,
		AttestationState:       attestationStateUseCase,
		Repos:                  testingRepos,
	}
	return testingUseCases, func() {
		cleanup()
	}, nil
}

var (
	_wireReferrerSharedIndexValue = &conf.ReferrerSharedIndex{}
)