package main

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"

	"mars/internal/handler"
	"mars/internal/handler/service"
	"mars/internal/handler/service/client"
	"mars/internal/platform/tracing"
)

func main() {
	tracer, closer := tracing.Init("mars")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("main")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	fileName := "file.txt"
	clnt := client.New()
	srvc := service.New(clnt)
	hndlr := handler.New(srvc)
	response, err := hndlr.NewPositionForRobots(ctx, fileName)

	if err != nil {
		logrus.Error("something went wrong.")
		logrus.Error(err)
	} else {
		logrus.Info("new positions of robots: ")
		logrus.Info(*response)
	}

}
