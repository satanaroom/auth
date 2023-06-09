package app

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/satanaroom/auth/internal/closer"
	"github.com/satanaroom/auth/internal/config"
	"github.com/satanaroom/auth/internal/interceptor"
	"github.com/satanaroom/auth/internal/metric"
	accessV1 "github.com/satanaroom/auth/pkg/access_v1"
	authV1 "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"
	userV1 "github.com/satanaroom/auth/pkg/user_v1"
	_ "github.com/satanaroom/auth/statik"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider  *serviceProvider
	grpcServer       *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
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

	wg.Add(4)
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

	go func() {
		defer wg.Done()
		if err := a.runSwaggerServer(); err != nil {
			logger.Fatalf("run swagger server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runPrometheusServer(); err != nil {
			logger.Fatalf("run prometheus server: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		metric.Init,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
		a.initSwaggerServer,
		a.initPrometheusServer,
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
		grpc.Creds(a.serviceProvider.TLSCredentials(ctx)),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.ValidateInterceptor,
				interceptor.ErrorCodesInterceptor,
				interceptor.NewRateLimiterInterceptor(a.serviceProvider.RateLimiter(ctx)).Unary,
				interceptor.MetricsInterceptor,
				interceptor.NewCircuitBreakerInterceptor(a.serviceProvider.CircuitBreaker(ctx)).Unary,
			),
		),
	)

	reflection.Register(a.grpcServer)

	userV1.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))
	authV1.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))
	accessV1.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.AccessImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := userV1.RegisterUserV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Host(), opts); err != nil {
		return fmt.Errorf("register user v1 handler from endpoint: %w", err)
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Host(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return fmt.Errorf("failed to create statik file system: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/access.swagger.json/", serveSwaggerFile("/access.swagger.json"))
	mux.HandleFunc("/user.swagger.json/", serveSwaggerFile("/user.swagger.json"))
	mux.HandleFunc("/auth.swagger.json/", serveSwaggerFile("/auth.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.serviceProvider.SwaggerConfig().Host(),
		Handler: mux,
	}

	return nil
}

func (a *App) initPrometheusServer(_ context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:    a.serviceProvider.PrometheusConfig().Host(),
		Handler: mux,
	}

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Infof("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Host())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Host())
	if err != nil {
		return fmt.Errorf("failed to get listener: %s", err.Error())
	}

	if err = a.grpcServer.Serve(list); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Host())

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	logger.Infof("Swagger server is running on %s", a.serviceProvider.SwaggerConfig().Host())

	if err := a.swaggerServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	logger.Infof("Prometheus server is running on %s", a.serviceProvider.PrometheusConfig().Host())

	if err := a.prometheusServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
