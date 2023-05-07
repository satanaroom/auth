package app

import (
	"context"
	"fmt"
	"net"

	"github.com/satanaroom/auth/internal/closer"
	"github.com/satanaroom/auth/internal/config"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, fmt.Errorf("init deps: %w", err)
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	if err := a.runGRPCServer(); err != nil {
		return fmt.Errorf("run GRPC server: %w", err)
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return fmt.Errorf("init: %w", err)
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)

	desc.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Port())
	if err != nil {
		logger.Fatalf("failed to get listener: %s", err.Error())
	}

	if err = a.grpcServer.Serve(list); err != nil {
		logger.Fatalf("failed to serve: %s", err.Error())
	}

	return nil
}
