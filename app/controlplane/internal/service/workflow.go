//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"

	pb "github.com/chainloop-dev/chainloop/app/controlplane/api/controlplane/v1"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkflowService struct {
	pb.UnimplementedWorkflowServiceServer
	*service

	useCase *biz.WorkflowUseCase
}

func NewWorkflowService(uc *biz.WorkflowUseCase, opts ...NewOpt) *WorkflowService {
	return &WorkflowService{
		service: newService(opts...),
		useCase: uc,
	}
}

func (s *WorkflowService) Create(ctx context.Context, req *pb.WorkflowServiceCreateRequest) (*pb.WorkflowServiceCreateResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	createOpts := &biz.WorkflowCreateOpts{
		OrgID:       currentOrg.ID,
		Name:        req.GetName(),
		Project:     req.GetProject(),
		Team:        req.GetTeam(),
		ContractID:  req.GetSchemaId(),
		Description: req.GetDescription(),
	}

	p, err := s.useCase.Create(ctx, createOpts)
	if err != nil {
		return nil, handleUseCaseErr("workflow", err, s.log)
	}

	workflow := bizWorkflowToPb(p)
	if err := workflow.ValidateAll(); err != nil {
		return nil, err
	}

	return &pb.WorkflowServiceCreateResponse{Result: workflow}, nil
}

func (s *WorkflowService) Update(ctx context.Context, req *pb.WorkflowServiceUpdateRequest) (*pb.WorkflowServiceUpdateResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	updateOpts := &biz.WorkflowUpdateOpts{
		Name:        req.Name,
		Project:     req.Project,
		Team:        req.Team,
		Public:      req.Public,
		Description: req.Description,
	}

	p, err := s.useCase.Update(ctx, currentOrg.ID, req.Id, updateOpts)
	if err != nil {
		return nil, handleUseCaseErr("workflow", err, s.log)
	}

	workflow := bizWorkflowToPb(p)
	if err := workflow.ValidateAll(); err != nil {
		return nil, err
	}

	return &pb.WorkflowServiceUpdateResponse{Result: workflow}, nil
}

func (s *WorkflowService) List(ctx context.Context, _ *pb.WorkflowServiceListRequest) (*pb.WorkflowServiceListResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	workflows, err := s.useCase.List(ctx, currentOrg.ID)
	if err != nil {
		return nil, handleUseCaseErr("workflow", err, s.log)
	}

	result := make([]*pb.WorkflowItem, 0, len(workflows))
	for _, p := range workflows {
		result = append(result, bizWorkflowToPb(p))
	}

	return &pb.WorkflowServiceListResponse{Result: result}, nil
}

func (s *WorkflowService) Delete(ctx context.Context, req *pb.WorkflowServiceDeleteRequest) (*pb.WorkflowServiceDeleteResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.useCase.Delete(ctx, currentOrg.ID, req.Id); err != nil {
		return nil, handleUseCaseErr("workflow", err, s.log)
	}

	return &pb.WorkflowServiceDeleteResponse{}, nil
}

func (s *WorkflowService) View(ctx context.Context, req *pb.WorkflowServiceViewRequest) (*pb.WorkflowServiceViewResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	wf, err := s.useCase.FindByIDInOrg(ctx, currentOrg.ID, req.Id)
	if err != nil {
		return nil, handleUseCaseErr("workflow", err, s.log)
	}

	return &pb.WorkflowServiceViewResponse{Result: bizWorkflowToPb(wf)}, nil
}

func bizWorkflowToPb(wf *biz.Workflow) *pb.WorkflowItem {
	item := &pb.WorkflowItem{
		Id: wf.ID.String(), Name: wf.Name, CreatedAt: timestamppb.New(*wf.CreatedAt),
		Project: wf.Project, Team: wf.Team, RunsCount: int32(wf.RunsCounter), Public: wf.Public,
		Description: wf.Description, ContractRevisionLatest: int32(wf.ContractRevisionLatest),
	}

	if wf.ContractID != uuid.Nil {
		item.ContractId = wf.ContractID.String()
	}

	if wf.LastRun != nil {
		item.LastRun = bizWorkFlowRunToPb(wf.LastRun)
	}

	return item
}
