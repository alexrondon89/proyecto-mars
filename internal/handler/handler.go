package handler

import (
	"context"
	"errors"
	"strconv"
	"strings"

	ot "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"

	"mars/internal/handler/service"
	"mars/internal/platform/files"
)

const (
	ErrorInScanner                string = "one error occurred trying to scan file: "
	ErrorInRobotPositionInput     string = "position input for robot is invalid"
	ErrorParsingStringToInt       string = "error parsing string to int: "
	ErrorInRobotInstructionsInput string = "set of instructions must be greater than 0 and smaller or equal to 100"
)

type Handler interface {
	NewPositionForRobots(ctx context.Context, fileName string) (*service.Response, error)
}
type handler struct {
	service service.Service
}

func New(srv service.Service) Handler {
	return &handler{
		service: srv,
	}
}

func (h *handler) NewPositionForRobots(ctx context.Context, fileName string) (*service.Response, error) {
	span, _ := ot.StartSpanFromContext(ctx, "Handler")
	span.SetTag("file name", fileName)
	defer span.Finish()

	file, err := files.OpenFile(fileName)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	scanner := files.ScannerFile(file)
	request := service.Request{}

	for scanner.Scan() {
		orientation := scanner.Text()
		robot, err := h.extractRobotPosition(orientation)
		if err != nil {
			scanner.Scan()
			continue //robot can not be sent to service because of invalid information
		}
		scanner.Scan()
		instructions := scanner.Text()
		err = h.extractRobotInstructions(instructions, robot)
		if err != nil {
			continue //robot can not be sent to service because of invalid information
		}
		request.Robots = append(request.Robots, *robot)
	}

	logrus.Info("list of robots: ")
	logrus.Info(request.Robots)

	if err = scanner.Err(); err != nil {
		return nil, errors.New(ErrorInScanner + err.Error())
	}

	robots := h.service.GetInstructionsToRobots(ctx, request)
	return robots, nil
}

func (h *handler) extractRobotPosition(orientation string) (*service.Robot, error) {
	oriList := strings.Split(orientation, " ")
	if len(oriList) != 3 {
		return nil, errors.New(ErrorInRobotPositionInput)
	}

	cx, err := strconv.Atoi(oriList[0])
	if err != nil {
		return nil, errors.New(ErrorParsingStringToInt + err.Error())
	}
	cy, err := strconv.Atoi(oriList[1])
	if err != nil {
		return nil, errors.New(ErrorParsingStringToInt + err.Error())
	}

	or := oriList[2]

	return &service.Robot{
		Position: service.Position{
			CoorX:       cx,
			CoorY:       cy,
			Orientation: or,
		},
	}, nil
}

func (h *handler) extractRobotInstructions(instructions string, robot *service.Robot) error {
	n := len(instructions)
	if n < 1 || n > 100 {
		return errors.New(ErrorInRobotInstructionsInput)
	}
	robot.Instructions = instructions
	return nil
}
