package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"mars/internal/handler/service/client"
	"mars/internal/handler/service/client/client_mock"
)

func Test_HappyPathForRobotMovementEvent(t *testing.T) {
	cli := client_mock.New()
	srv := New(cli)
	firstRobot := Robot{Position: Position{CoorX: 50, CoorY: 49, Orientation: "N"}, Instructions: "F"}
	req := Request{Robots: []Robot{firstRobot}}
	resp := &client.Response{CoorX: 50, CoorY: 50, Orientation: "N"}

	cli.On("CheckIfInstructionIsDangerous", mock.Anything, mock.Anything).Return(false)
	cli.On("CheckConnection", mock.Anything, mock.Anything, mock.Anything).Return(true)
	cli.On("MoveRobotPosition", mock.Anything, mock.Anything).Return(resp)
	response := srv.GetInstructionsToRobots(context.Background(), req)

	assert.Equal(t, len(response.Robots), len(req.Robots))
	assert.Equal(t, response.Robots[0].Position.CoorX, resp.CoorX)
	assert.Equal(t, response.Robots[0].Position.CoorY, resp.CoorY)
	assert.Equal(t, response.Robots[0].Position.Orientation, resp.Orientation)
}

func Test_HappyPathForRobotOrientationEvent(t *testing.T) {
	cli := client_mock.New()
	srv := New(cli)
	firstRobot := Robot{Position: Position{CoorX: 50, CoorY: 49, Orientation: "N"}, Instructions: "L"}
	req := Request{Robots: []Robot{firstRobot}}
	resp := &client.Response{CoorX: 50, CoorY: 49, Orientation: "W"}

	cli.On("CheckIfInstructionIsDangerous", mock.Anything, mock.Anything).Return(false)
	cli.On("MoveRobotOrientation", mock.Anything, mock.Anything).Return(resp)
	cli.On("CheckConnection", mock.Anything, mock.Anything, mock.Anything).Return(true)
	response := srv.GetInstructionsToRobots(context.Background(), req)

	assert.Equal(t, len(response.Robots), len(req.Robots))
	assert.Equal(t, response.Robots[0].Position.CoorX, resp.CoorX)
	assert.Equal(t, response.Robots[0].Position.CoorY, resp.CoorY)
	assert.Equal(t, response.Robots[0].Position.Orientation, resp.Orientation)
}

func Test_ExistsDangerousInstruction(t *testing.T) {
	cli := client_mock.New()
	srv := New(cli)
	firstRobot := Robot{Position: Position{CoorX: 50, CoorY: 49, Orientation: "N"}, Instructions: "F"}
	req := Request{Robots: []Robot{firstRobot}}
	resp := &client.Response{CoorX: 50, CoorY: 49, Orientation: "N"}

	cli.On("CheckIfInstructionIsDangerous", mock.Anything, mock.Anything).Return(true)
	response := srv.GetInstructionsToRobots(context.Background(), req)

	assert.Equal(t, len(response.Robots), len(req.Robots))
	assert.Equal(t, response.Robots[0].Position.CoorX, resp.CoorX)
	assert.Equal(t, response.Robots[0].Position.CoorY, resp.CoorY)
	assert.Equal(t, response.Robots[0].Position.Orientation, resp.Orientation)
}

func Test_RobotIsLost(t *testing.T) {
	cli := client_mock.New()
	srv := New(cli)
	firstRobot := Robot{Position: Position{CoorX: 50, CoorY: 49, Orientation: "N"}, Instructions: "F"}
	req := Request{Robots: []Robot{firstRobot}}
	resp := &client.Response{CoorX: 50, CoorY: 49, Orientation: "N"}

	cli.On("CheckIfInstructionIsDangerous", mock.Anything, mock.Anything).Return(false)
	cli.On("MoveRobotPosition", mock.Anything, mock.Anything).Return(resp)
	cli.On("CheckConnection", mock.Anything, mock.Anything).Return(false)
	response := srv.GetInstructionsToRobots(context.Background(), req)

	assert.Equal(t, len(response.Robots), len(req.Robots))
	assert.Equal(t, response.Robots[0].Position.CoorX, resp.CoorX)
	assert.Equal(t, response.Robots[0].Position.CoorY, resp.CoorY)
	assert.Equal(t, response.Robots[0].Position.Orientation, resp.Orientation)
}
