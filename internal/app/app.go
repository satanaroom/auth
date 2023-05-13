package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/satanaroom/auth/internal/closer"
	"github.com/satanaroom/auth/internal/config"
	"github.com/satanaroom/auth/internal/interceptor"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
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

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := a.runGRPCServer(); err != nil {
			logger.Fatalf("run GRPC server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runHTTPServer(); err != nil {
			logger.Fatalf("run HTTP server: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
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
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
		//grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := desc.RegisterAuthV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Port(), opts); err != nil {
		return fmt.Errorf("register auth v1 handler from endpoint: %w", err)
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Port(),
		Handler: mux,
	}

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Infof("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Port())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Port())
	if err != nil {
		return fmt.Errorf("failed to get listener: %s", err.Error())
	}

	if err = a.grpcServer.Serve(list); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Port())

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}
