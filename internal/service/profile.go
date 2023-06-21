package service

import (
	"context"
	"github.com/google/uuid"
	auth "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
	user "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"sync"
	"time"
)

func (s *Service) GetProfile(profileId *uuid.UUID, profileNickname *string, requesterId uuid.UUID) (entity.Profile, error) {
	wg := sync.WaitGroup{}

	profile := entity.Profile{}

	id, nickname := s.extractIdentifiers(profileId, profileNickname, &profile)

	wg.Add(1)
	go s.fillProfileAuthInfo(id, nickname, &profile, &wg)
	if profile.Id == nil {
		wg.Wait()
		if profile.Id == nil {
			return entity.Profile{}, fail.GrpcNotFound
		}
	}

	wg.Add(1)
	go s.fillProfileUserInfo(*profile.Id, &profile, &wg)

	wg.Wait()

	if profile.Id == nil {
		return entity.Profile{}, fail.GrpcNotFound
	}

	if profile.IsBlocked {
		return entity.Profile{Id: profile.Id, IsBlocked: true}, nil
	}

	if *profile.Id != requesterId.String() {
		s.clearConfidentialInfo(&profile)
	}

	return profile, nil
}

func (s *Service) extractIdentifiers(profileId *uuid.UUID, profileNickname *string, profile *entity.Profile) (string, string) {
	id := ""
	nickname := ""
	if profileId != nil {
		id = profileId.String()
		profile.Id = &id
	}
	if profileNickname != nil {
		nickname = *profileNickname
	}

	return id, nickname
}

func (s *Service) fillProfileAuthInfo(id, nickname string, profile *entity.Profile, wg *sync.WaitGroup) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	info, err := s.repos.Auth.GetAuthInfo(ctx, &auth.GetAuthInfoRequest{
		Id:       id,
		Nickname: nickname,
	})
	cancelCtx()

	if err == nil {
		profile.Id = &info.Id
		profile.Nickname = info.Nickname
		profile.IsBlocked = info.IsBlocked
		profile.Email = &info.Email
		profile.Role = info.Role
		registrationTimestamp := info.RegistrationTimestamp.AsTime()
		profile.RegistrationTimestamp = &registrationTimestamp

		if info.OAuth != nil {
			profile.OAuth = &entity.OAuth{
				GoogleId: info.OAuth.GoogleId,
				VkId:     info.OAuth.VkId,
			}
		}
	}

	wg.Done()
}

func (s *Service) fillProfileUserInfo(id string, profile *entity.Profile, wg *sync.WaitGroup) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	info, err := s.repos.User.GetUserInfo(ctx, &user.GetUserInfoRequest{
		UserId: id,
	})
	cancelCtx()

	if err == nil {
		profile.Id = &info.UserId
		profile.FirstName = info.FirstName
		profile.LastName = info.LastName
		profile.Description = info.Description
		profile.AvatarLink = info.Avatar
	}

	wg.Done()
}

func (s *Service) clearConfidentialInfo(profile *entity.Profile) {
	profile.Email = nil
	profile.OAuth = nil
}
