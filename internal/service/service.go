package service

import (
	"github.com/mephistolie/chefbook-backend-profile/internal/repository"
)

type Service struct {
	repos *repository.Repositories
}

func New(
	repos *repository.Repositories,
) *Service {
	return &Service{repos: repos}
}
