package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	authV1 "github.com/satanaroom/auth/internal/api/auth_v1"
	noteRepository "github.com/satanaroom/auth/internal/repository/auth"
	noteService "github.com/satanaroom/auth/internal/service/auth"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("loading .env file: %s", err.Error())
	}

	grpcPort := os.Getenv("GRPC_PORT")
	pgPort := os.Getenv("PG_PORT")
	pgDatabase := os.Getenv("PG_DATABASE")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")

	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Fatalf("failed to get listener: %s", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)

	dbConn := fmt.Sprintf(
		"host=localhost port=%s dbname=%s user=%s password=%s sslmode=disable",
		pgPort, pgDatabase, pgUser, pgPassword,
	)
	pgCfg, err := pgxpool.ParseConfig(dbConn)
	if err != nil {
		logger.Fatalf("failed to get db config: %s", err.Error())
	}

	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		logger.Fatalf("failed to get db connection: %s", err.Error())
	}
	defer dbc.Close()

	if err = dbc.Ping(ctx); err != nil {
		logger.Fatalf("ping database: %s", err.Error())
	}

	authRepo := noteRepository.NewRepository(dbc)
	authSrv := noteService.NewService(authRepo)

	desc.RegisterAuthV1Server(s, authV1.NewImplementation(authSrv))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err = s.Serve(list); err != nil {
			logger.Fatalf("failed to serve: %s", err.Error())
		}
	}()
	logger.Infof("service started on port %s", grpcPort)

	<-quit

	logger.Info("service shutting down")
}
