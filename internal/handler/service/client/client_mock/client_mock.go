package client_mock

import (
	"context"
	
	"github.com/stretchr/testify/mock"

	"mars/internal/handler/service/client"
)

type ClientMock struct {
	mock.Mock
}

func New() *ClientMock {
	return &ClientMock{}
}

func (cm *ClientMock) MoveRobotPosition(ctx context.Context, robot *client.Request) *client.Response {
	args := cm.Called(ctx, robot)
	if len(cm.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return &client.Response{}
	}
	return args.Get(0).(*client.Response)
}

func (cm *ClientMock) MoveRobotOrientation(ctx context.Context, robot *client.Request) *client.Response {
	args := cm.Called(ctx, robot)
	if len(cm.ExpectedCalls) == 0 {
		return nil
	}
	if args.Get(0) == nil {
		return &client.Response{}
	}
	return args.Get(0).(*client.Response)
}

func (cm *ClientMock) CheckIfInstructionIsDangerous(ctx context.Context, robot *client.Request) bool {
	args := cm.Called(ctx, robot)
	if len(cm.ExpectedCalls) == 0 {
		return false
	}
	if args.Get(0) == nil {
		return true
	}
	return args.Get(0).(bool)
}

func (cm *ClientMock) CheckConnection(ctx context.Context, robot *client.Request, resp *client.Response) bool {
	args := cm.Called(ctx, robot)
	if len(cm.ExpectedCalls) == 0 {
		return false
	}
	if args.Get(0) == nil {
		return true
	}
	return args.Get(0).(bool)
}
