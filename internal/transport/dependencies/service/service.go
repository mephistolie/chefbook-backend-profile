package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
)

type Service interface {
	GetProfilesMinInfo(ctx context.Context, profileIds []string) (map[string]entity.ProfileMinInfo, error)
	GetProfile(ctx context.Context, profileId *uuid.UUID, profileNickname *string, requesterId uuid.UUID) (entity.Profile, error)
}
