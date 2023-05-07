package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/satanaroom/auth/internal/app"
	"github.com/satanaroom/auth/pkg/logger"
)

func main() {
	ctx := context.Background()

	authApp, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatalf("failed to initialize app: %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	logger.Info("service starting up")

	if err = authApp.Run(); err != nil {
		logger.Fatalf("failed to run app: %s", err.Error())
	}

	<-quit

	logger.Info("service shutting down")
}
