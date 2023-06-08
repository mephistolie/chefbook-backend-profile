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
	email := ""
	if profile.Email != nil {
		email = *profile.Email
	}
	nickname := ""
	if profile.Nickname != nil {
		nickname = *profile.Nickname
	}
	role := ""
	if profile.Role != nil {
		role = *profile.Role
	}
	var registrationTimestamp *timestamppb.Timestamp = nil
	if profile.RegistrationTimestamp != nil {
		registrationTimestamp = timestamppb.New(*profile.RegistrationTimestamp)
	}
	var oauthPtr *api.OAuthConnections = nil
	if profile.OAuth != nil {
		oauth := api.OAuthConnections{}
		if profile.OAuth.GoogleId != nil {
			oauth.GoogleId = *profile.OAuth.GoogleId
		}
		if profile.OAuth.VkId != nil {
			oauth.VkId = *profile.OAuth.VkId
		}
		oauthPtr = &oauth
	}

	firstName := ""
	if profile.FirstName != nil {
		firstName = *profile.FirstName
	}
	lastName := ""
	if profile.LastName != nil {
		lastName = *profile.LastName
	}
	description := ""
	if profile.Description != nil {
		description = *profile.Description
	}
	avatar := ""
	if profile.AvatarLink != nil {
		avatar = *profile.AvatarLink
	}

	return &api.GetProfileResponse{
		Id:                    id,
		Email:                 email,
		Nickname:              nickname,
		Role:                  role,
		RegistrationTimestamp: registrationTimestamp,
		IsBlocked:             profile.IsBlocked,
		OAuth:                 oauthPtr,

		FirstName:   firstName,
		LastName:    lastName,
		Description: description,
		Avatar:      avatar,
	}
}
