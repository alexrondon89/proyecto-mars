package client

import (
	"context"

	ot "github.com/opentracing/opentracing-go"

	"mars/internal/platform/repository"
)

type Client interface {
	MoveRobotPosition(ctx context.Context, robot *Request) *Response
	MoveRobotOrientation(ctx context.Context, robot *Request) *Response
	CheckIfInstructionIsDangerous(ctx context.Context, robot *Request) bool
	CheckConnection(ctx context.Context, robot *Request, resp *Response) bool
}

type client struct {
	Movements             []repository.Movement
	DangerousInstructions []repository.DangerousInstructions
}

func New() Client {
	return &client{
		Movements:             repository.GetOrientationRepository(),
		DangerousInstructions: repository.GetDangerousInstructionsRepository(),
	}
}

func (c *client) MoveRobotPosition(ctx context.Context, req *Request) *Response {
	span, _ := ot.StartSpanFromContext(ctx, "client")
	span.SetTag("method", "MoveRobotPosition")
	defer span.Finish()

	response := Response{
		CoorX:       req.CoorX,
		CoorY:       req.CoorY,
		Orientation: req.Orientation,
	}

	switch req.Orientation {
	case N:
		response.CoorY = response.CoorY + 1
	case S:
		response.CoorY = response.CoorY - 1
	case E:
		response.CoorX = response.CoorX + 1
	case W:
		response.CoorX = response.CoorX - 1
	}

	return &response
}

func (c *client) MoveRobotOrientation(ctx context.Context, robot *Request) *Response {
	span, _ := ot.StartSpanFromContext(ctx, "client")
	span.SetTag("method", "MoveRobotOrientation")
	defer span.Finish()

	for _, movement := range c.Movements {
		if movement[Turn] == robot.Instruction && movement[Orientation] == robot.Orientation {
			robot.Orientation = movement[Result]
			break
		}
	}
	return &Response{
		CoorX:       robot.CoorX,
		CoorY:       robot.CoorY,
		Orientation: robot.Orientation,
	}
}

func (c *client) CheckIfInstructionIsDangerous(ctx context.Context, robot *Request) bool {
	span, _ := ot.StartSpanFromContext(ctx, "client")
	span.SetTag("method", "CheckIfInstructionIsDangerous")
	defer span.Finish()

	for _, position := range c.DangerousInstructions {
		if position[CoorX] == robot.CoorX &&
			position[CoorY] == robot.CoorY &&
			position[Orientation] == robot.Orientation &&
			position[Instruction] == robot.Instruction {
			return true
		}
	}

	return false
}

func (c *client) CheckConnection(ctx context.Context, req *Request, resp *Response) bool {
	span, _ := ot.StartSpanFromContext(ctx, "client")
	span.SetTag("method", "CheckConnection")
	defer span.Finish()

	if resp.CoorX > 50 || resp.CoorX < 0 || resp.CoorY > 50 || resp.CoorY < 0 {
		instruction := map[string]interface{}{
			CoorX:       req.CoorX,
			CoorY:       req.CoorY,
			Orientation: req.Orientation,
			Instruction: req.Instruction,
		}
		resp.Orientation = resp.Orientation + " LOST"
		c.DangerousInstructions = append(c.DangerousInstructions, instruction)
		return false
	}

	return true
}
