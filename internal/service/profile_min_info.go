package service

import (
	"context"
	auth "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
	user "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"sync"
	"time"
)

func (s *Service) GetProfilesMinInfo(profileIds []string) (map[string]entity.ProfileMinInfo, error) {
	wg := sync.WaitGroup{}

	fallbackNames := make(map[string]string)

	wg.Add(1)
	go func() {
		fallbackNames = s.getProfilesAuthNames(profileIds)
		wg.Done()
	}()

	infos, hasAllNames := s.getProfilesUserInfo(profileIds)

	if !hasAllNames {
		wg.Wait()
		for id, info := range infos {
			if info.VisibleName == nil {
				if fallbackName, ok := fallbackNames[id]; ok {
					info.VisibleName = &fallbackName
					infos[id] = info
				}
			}
		}
	}

	return infos, nil
}

func (s *Service) getProfilesUserInfo(profileIds []string) (map[string]entity.ProfileMinInfo, bool) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := s.repos.User.GetUsersMinInfo(ctx, &user.GetUsersMinInfoRequest{UserIds: profileIds})
	cancelCtx()

	requestProfilesCount := len(profileIds)
	responseProfileCount := len(res.Infos)

	hasAllNames := responseProfileCount >= requestProfilesCount

	infos := make(map[string]entity.ProfileMinInfo)
	if err == nil {
		for id, dto := range res.Infos {
			info := entity.ProfileMinInfo{
				Id:          id,
				VisibleName: dto.FullName,
				Avatar:      dto.Avatar,
			}
			if info.VisibleName == nil {
				hasAllNames = false
			}
			infos[id] = info
		}
	} else {
		log.Warnf("unable to get profiles minimal info")
	}

	return infos, hasAllNames
}

func (s *Service) getProfilesAuthNames(profileIds []string) map[string]string {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := s.repos.Auth.GetVisibleNames(ctx, &auth.GetVisibleNamesRequest{UserIds: profileIds})
	cancelCtx()

	names := make(map[string]string)

	if err == nil {
		for id, name := range res.UserVisibleNames {
			names[id] = name
		}
	} else {
		log.Warnf("unable to get profiles auth names")
	}

	return names
}
