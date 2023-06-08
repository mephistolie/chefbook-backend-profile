package service

import (
	"context"
	"github.com/google/uuid"
	auth "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
	"github.com/mephistolie/chefbook-backend-profile/internal/repository"
	user "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"strings"
	"sync"
	"time"
)

type Service struct {
	repos *repository.Repositories
}

func New(
	repos *repository.Repositories,
) *Service {
	return &Service{repos: repos}
}

func (s *Service) GetProfile(profileId *uuid.UUID, profileNickname *string, requesterId uuid.UUID) (entity.Profile, error) {
	wg := sync.WaitGroup{}

	profile := entity.Profile{}

	id, nickname := s.extractIdentifiers(profileId, profileNickname, &profile)

	wg.Add(1)
	go s.fillAuthInfo(id, nickname, &profile, &wg)
	if profile.Id == nil {
		wg.Wait()
		if profile.Id == nil {
			return entity.Profile{}, fail.GrpcNotFound
		}
	}

	wg.Add(1)
	go s.fillUserInfo(*profile.Id, &profile, &wg)

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

func (s *Service) fillAuthInfo(id, nickname string, profile *entity.Profile, wg *sync.WaitGroup) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	info, err := s.repos.Auth.GetAuthInfo(ctx, &auth.GetAuthInfoRequest{
		Id:       id,
		Nickname: nickname,
	})
	cancelCtx()

	if err == nil {
		profile.Id = &info.Id

		if len(info.Email) > 0 {
			profile.Email = &info.Email
		}
		if len(info.Nickname) > 0 {
			profile.Nickname = &info.Nickname
		}
		if len(info.Role) > 0 {
			role := strings.ToLower(info.Role)
			profile.Role = &role
		}
		profile.IsBlocked = info.IsBlocked
		registrationTimestamp := info.RegistrationTimestamp.AsTime()
		profile.RegistrationTimestamp = &registrationTimestamp

		if info.OAuth != nil {
			oauth := entity.OAuth{}
			if len(info.OAuth.GoogleId) > 0 {
				oauth.GoogleId = &info.OAuth.GoogleId
			}
			if info.OAuth.VkId > 0 {
				oauth.VkId = &info.OAuth.VkId
			}
			profile.OAuth = &oauth
		}
	}

	wg.Done()
}

func (s *Service) fillUserInfo(id string, profile *entity.Profile, wg *sync.WaitGroup) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	info, err := s.repos.User.GetUserInfo(ctx, &user.GetUserInfoRequest{
		UserId: id,
	})
	cancelCtx()

	if err == nil {
		profile.Id = &info.UserId

		if len(info.FirstName) > 0 {
			profile.FirstName = &info.FirstName
		}
		if len(info.LastName) > 0 {
			profile.LastName = &info.LastName
		}
		if len(info.Description) > 0 {
			profile.Description = &info.Description
		}
		if len(info.Avatar) > 0 {
			profile.AvatarLink = &info.Avatar
		}
	}

	wg.Done()
}

func (s *Service) clearConfidentialInfo(profile *entity.Profile) {
	profile.Email = nil
	profile.OAuth = nil
}
