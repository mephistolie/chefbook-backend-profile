package grpc

import (
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-profile/internal/transport/dependencies/service"
)

type ProfileServer struct {
	api.UnsafeProfileServiceServer
	service service.Service
}

func NewServer(service service.Service) *ProfileServer {
	return &ProfileServer{service: service}
}
