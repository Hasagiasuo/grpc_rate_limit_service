package rlgrpc

import (
	"context"
	"fmt"
	rlservicepb "rlservice/gen/v1"
	"rlservice/pkg/logger"
)

// ! service usecase methods
type ServiceProvider interface {
	UpdateService(token string, cost int) error
	RemoveService(token string) error
}

// todo: script usecase methods
type ScriptProvider interface {
	ServeTimeRule(ctx context.Context, req *rlservicepb.CheckRequest) (*rlservicepb.CheckResponse, int, error)
}

type Service struct {
	log             *logger.Logger
	serviceProvider ServiceProvider
	scriptProvider  ScriptProvider
	rlservicepb.UnimplementedRateLimitServiceServer
}

func NewService(scriptProvider ScriptProvider, serviceProvider ServiceProvider, log *logger.Logger) *Service {
	return &Service{
		log:             log,
		scriptProvider:  scriptProvider,
		serviceProvider: serviceProvider,
	}
}

// Todo: update this function
func (s *Service) Check(ctx context.Context, req *rlservicepb.CheckRequest) (*rlservicepb.CheckResponse, error) {
	const op = "rlgrpc.Check"
	resp, cost, err := s.scriptProvider.ServeTimeRule(ctx, req)
	if err != nil {
		s.log.Error(op, fmt.Sprintf("cannot serve time rule: %v", err))
		return nil, err
	}
	if err := s.serviceProvider.UpdateService(req.GetServiceToken(), cost); err != nil {
		s.log.Error(op, fmt.Sprintf("cannot update service: %v", err))
		return nil, err
	}
	return  resp, nil
}
