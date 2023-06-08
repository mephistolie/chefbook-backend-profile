package service

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
)

type Service interface {
	GetProfile(profileId *uuid.UUID, profileNickname *string, requesterId uuid.UUID) (entity.Profile, error)
}
