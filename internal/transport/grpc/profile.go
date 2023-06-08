package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-profile/internal/transport/grpc/dto"
)

func (s *ProfileServer) GetProfile(_ context.Context, req *api.GetProfileRequest) (*api.GetProfileResponse, error) {
	requesterId, err := uuid.Parse(req.RequesterId)
	if err != nil {
		return nil, fail.GrpcInvalidBody
	}

	var profileId *uuid.UUID = nil
	if id, err := uuid.Parse(req.ProfileId); err == nil {
		profileId = &id
	}
	var profileNickname *string = nil
	if len(req.ProfileNickname) > 0 {
		profileNickname = &req.ProfileNickname
	}

	if profileId == nil && profileNickname == nil {
		profileId = &requesterId
	}

	profile, err := s.service.GetProfile(profileId, profileNickname, requesterId)
	if err != nil {
		return nil, err
	}

	return dto.NewGetProfileResponse(profile), nil
}
