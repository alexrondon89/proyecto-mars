package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"mars/internal/handler/service"
	"mars/internal/handler/service/service_mock"
)

func Test_fileDoNotExist(t *testing.T) {
	srv := service_mock.New()
	handler := New(srv)
	fileName := "file.txt"
	resp := &service.Response{Robots: []service.Robot{}}

	srv.On("GetInstructionsToRobots", mock.Anything, mock.Anything).Return(resp)
	response, err := handler.NewPositionForRobots(context.Background(), fileName)
	assert.Nil(t, response)
	assert.Equal(t, "one error occurred trying to open file: open file.txt: no such file or directory", err.Error())
}
