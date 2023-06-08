package app

import (
	userpb "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func setHealthStatus(healthServer *health.Server, status healthpb.HealthCheckResponse_ServingStatus) {
	healthServer.SetServingStatus("", status)
	healthServer.SetServingStatus(userpb.ProfileService_ServiceDesc.ServiceName, status)
}
