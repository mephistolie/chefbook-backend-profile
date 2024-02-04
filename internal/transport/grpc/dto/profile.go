package dto

import (
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)
import api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"

func NewGetProfileResponse(profile entity.Profile) *api.GetProfileResponse {
	id := ""
	if profile.Id != nil {
		id = *profile.Id
	}
	var registrationTimestamp *timestamppb.Timestamp = nil
	if profile.RegistrationTimestamp != nil {
		registrationTimestamp = timestamppb.New(*profile.RegistrationTimestamp)
	}
	var oAuth *api.OAuthConnections = nil
	if profile.OAuth != nil {
		oAuth = &api.OAuthConnections{
			GoogleId: profile.OAuth.GoogleId,
			VkId:     profile.OAuth.VkId,
		}
	}

	return &api.GetProfileResponse{
		Id:                    id,
		Email:                 profile.Email,
		Nickname:              profile.Nickname,
		Role:                  profile.Role,
		RegistrationTimestamp: registrationTimestamp,
		IsBlocked:             profile.IsBlocked,
		OAuth:                 oAuth,

		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		Description: profile.Description,
		Avatar:      profile.AvatarLink,

		SubscriptionPlan: profile.SubscriptionPlan,
	}
}
