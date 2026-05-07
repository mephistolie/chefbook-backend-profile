package app

import (
	"context"
	"fmt"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	profilepb "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-profile/internal/config"
	"github.com/mephistolie/chefbook-backend-profile/internal/repository"
	"github.com/mephistolie/chefbook-backend-profile/internal/service"
	profile "github.com/mephistolie/chefbook-backend-profile/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"
)

func Run(cfg *config.Config) {
	log.InitWithService("profile", *cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	ctx := context.Background()

	repositories, err := repository.NewRepositories(cfg)
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "app.startup.failed",
			Message:   "service startup failed",
			Component: "app",
		}, err)
		return
	}

	profileService := service.New(repositories)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.Port))
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "app.startup.failed",
			Message:   "service startup failed",
			Component: "app",
		}, err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			log.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	userServer := profile.NewServer(profileService)

	setHealthStatus(healthServer, healthpb.HealthCheckResponse_SERVING)

	profilepb.RegisterProfileServiceServer(grpcServer, userServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.LogError(ctx, log.Event{
				Event:     "grpc.server.failed",
				Message:   "error occurred while running grpc server",
				Component: log.ComponentGRPC,
			}, err)
		} else {
			log.Log(ctx, log.Event{
				Event:     "grpc.server.started",
				Message:   "grpc server started",
				Component: log.ComponentGRPC,
			})
		}
	}()

	wait := shutdown.Graceful(ctx, 5*time.Second, map[string]shutdown.Operation{
		"grpc-server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"services": func(ctx context.Context) error {
			return repositories.Stop()
		},
	})
	<-wait
}
