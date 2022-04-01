package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_robotMovePositionToNortSuccesfully(t *testing.T) {
	cli := New()
	request := &Request{CoorX: 2, CoorY: 3, Orientation: "N", Instruction: "F"}
	response := cli.MoveRobotPosition(context.Background(), request)
	assert.Equal(t, 4, response.CoorY)
	assert.Equal(t, request.CoorX, response.CoorX)
}

func Test_robotMovePositionToSouthSuccesfully(t *testing.T) {
	cli := New()
	request := &Request{CoorX: 2, CoorY: 3, Orientation: "S", Instruction: "F"}
	response := cli.MoveRobotPosition(context.Background(), request)
	assert.Equal(t, 2, response.CoorY)
	assert.Equal(t, request.CoorX, response.CoorX)
}

func Test_robotMovePositionToEastSuccesfully(t *testing.T) {
	cli := New()
	request := &Request{CoorX: 2, CoorY: 3, Orientation: "E", Instruction: "F"}
	response := cli.MoveRobotPosition(context.Background(), request)
	assert.Equal(t, 3, response.CoorX)
	assert.Equal(t, request.CoorY, response.CoorY)
}

func Test_robotMovePositionToWestSuccesfully(t *testing.T) {
	cli := New()
	request := &Request{CoorX: 2, CoorY: 3, Orientation: "W", Instruction: "F"}
	response := cli.MoveRobotPosition(context.Background(), request)
	assert.Equal(t, 1, response.CoorX)
	assert.Equal(t, request.CoorY, response.CoorY)
}

func Test_robotMoveOrientationSucessfully(t *testing.T) {
	cli := New()
	request := &Request{CoorX: 0, CoorY: 0, Orientation: "W", Instruction: "L"}
	response := cli.MoveRobotOrientation(context.Background(), request)
	assert.Equal(t, "S", response.Orientation)
}

func Test_robotDontMoveOrientationBecauseOfInvalidInstruction(t *testing.T) {
	cli := New()
	request := &Request{CoorX: 0, CoorY: 0, Orientation: "W", Instruction: "Q"}
	response := cli.MoveRobotOrientation(context.Background(), request)
	assert.Equal(t, "W", response.Orientation)
}

func Test_robotLostConnection(t *testing.T) {
	cli := New()
	req := &Request{CoorX: 50, CoorY: 0, Orientation: "S", Instruction: "F"}
	resp := &Response{CoorX: 50, CoorY: -1, Orientation: "S"}
	orientationUpdated := resp.Orientation + " LOST"

	response := cli.CheckConnection(context.Background(), req, resp)
	assert.Equal(t, response, false)
	assert.Equal(t, orientationUpdated, resp.Orientation)
}

func Test_robotCheckConnection(t *testing.T) {
	cli := New()
	req := &Request{CoorX: 50, CoorY: 0, Orientation: "S", Instruction: "F"}
	resp := &Response{CoorX: 50, CoorY: 1, Orientation: "S"}
	orientationUpdated := resp.Orientation

	response := cli.CheckConnection(context.Background(), req, resp)
	assert.Equal(t, response, true)
	assert.Equal(t, orientationUpdated, resp.Orientation)
}
