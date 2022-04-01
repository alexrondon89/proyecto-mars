package service

import (
	"context"
	
	ot "github.com/opentracing/opentracing-go"

	"mars/internal/handler/service/client"
)

const ForwardEvent string = "F"

type Service interface {
	GetInstructionsToRobots(ctx context.Context, request Request) *Response
	ReadInstructionsToEachRobot(ctx context.Context, robot Robot) *Robot
}

type service struct {
	client client.Client
}

func New(client client.Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetInstructionsToRobots(ctx context.Context, request Request) *Response {
	span, _ := ot.StartSpanFromContext(ctx, "service")
	span.SetTag("method", "GetInstructionsToRobots")
	span.SetTag("robots", request.Robots)
	defer span.Finish()

	var robotList []Robot
	for _, robot := range request.Robots {
		robotUpdated := s.ReadInstructionsToEachRobot(ctx, robot)
		robotList = append(robotList, *robotUpdated)
	}

	return &Response{
		Robots: robotList,
	}
}

func (s *service) ReadInstructionsToEachRobot(ctx context.Context, robot Robot) *Robot {
	span, _ := ot.StartSpanFromContext(ctx, "service")
	span.SetTag("method", "ReadInstructionsToEachRobot")
	span.SetTag("robot position", robot.Position)
	span.SetTag("robot instructions ", robot.Instructions)

	defer span.Finish()

	for _, instruction := range robot.Instructions {
		ins := string(instruction)
		var robotResponse *client.Response

		request := s.MapperClientRequest(robot, ins)
		if s.client.CheckIfInstructionIsDangerous(ctx, request) {
			continue
		}

		if ins == ForwardEvent {
			robotResponse = s.client.MoveRobotPosition(ctx, request)
		} else {
			robotResponse = s.client.MoveRobotOrientation(ctx, request)
		}

		if !s.client.CheckConnection(ctx, request, robotResponse) {
			s.MapperClientResponse(robotResponse, &robot)
			break
		}

		s.MapperClientResponse(robotResponse, &robot)
	}

	return &robot
}

func (s *service) MapperClientRequest(robot Robot, instruction string) *client.Request {
	return &client.Request{
		CoorX:       robot.Position.CoorX,
		CoorY:       robot.Position.CoorY,
		Orientation: robot.Position.Orientation,
		Instruction: instruction,
	}
}

func (s *service) MapperClientResponse(resp *client.Response, robot *Robot) {
	robot.Position.CoorX = resp.CoorX
	robot.Position.CoorY = resp.CoorY
	robot.Position.Orientation = resp.Orientation
}
