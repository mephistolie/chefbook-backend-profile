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

	infos := make(map[string]entity.ProfileMinInfo)
	fallbackNames := make(map[string]string)

	wg.Add(1)
	go s.fillProfilesAuthNames(profileIds, &fallbackNames, &wg)

	hasAllNames := s.fillProfilesUserInfo(profileIds, &infos)

	if !hasAllNames {
		log.Debug("unable to get user name for several users; using nicknames...")
		wg.Wait()
		for id, info := range infos {
			if info.VisibleName == nil {
				if fallbackName, ok := fallbackNames[id]; ok {
					info.VisibleName = &fallbackName
				}
			}
		}
	}

	return infos, nil
}

func (s *Service) fillProfilesUserInfo(profileIds []string, infos *map[string]entity.ProfileMinInfo) bool {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := s.repos.User.GetUsersMinInfo(ctx, &user.GetUsersMinInfoRequest{UserIds: profileIds})
	cancelCtx()

	requestProfilesCount := len(profileIds)
	responseProfileCount := len(res.Infos)

	hasAllNames := responseProfileCount >= requestProfilesCount

	if err == nil {
		for id, dto := range res.Infos {
			info := entity.ProfileMinInfo{
				Id: id,
			}
			if len(dto.FullName) > 0 {
				info.VisibleName = &dto.FullName
			} else {
				hasAllNames = false
			}
			if len(dto.Avatar) > 0 {
				info.Avatar = &dto.Avatar
			}

			(*infos)[id] = info
		}
	}

	return hasAllNames
}

func (s *Service) fillProfilesAuthNames(profileIds []string, names *map[string]string, wg *sync.WaitGroup) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := s.repos.Auth.GetVisibleNames(ctx, &auth.GetVisibleNamesRequest{UserIds: profileIds})
	cancelCtx()

	if err == nil {
		for id, name := range res.UserVisibleNames {
			(*names)[id] = name
		}
	}

	wg.Done()
}
