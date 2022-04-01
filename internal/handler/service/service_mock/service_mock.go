package service_mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"mars/internal/handler/service"
)

type ServiceMock struct {
	mock.Mock
}

func New() *ServiceMock {
	return &ServiceMock{}
}

func (sm *ServiceMock) GetInstructionsToRobots(ctx context.Context, request service.Request) *service.Response {
	args := sm.Called(ctx, request)
	if len(sm.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return &service.Response{}
	}
	return args.Get(0).(*service.Response)
}

func (sm *ServiceMock) ReadInstructionsToEachRobot(ctx context.Context, request service.Robot) *service.Robot {
	args := sm.Called(ctx, request)
	if len(sm.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return &service.Robot{}
	}
	return args.Get(0).(*service.Robot)
}
