package service

import (
	"context"
	"fmt"
	auth "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-profile/internal/entity"
	user "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"strings"
	"sync"
	"time"
)

func (s *Service) GetProfilesMinInfo(profileIds []string) (map[string]entity.ProfileMinInfo, error) {
	wg := sync.WaitGroup{}

	var nicknames map[string]string
	var users map[string]*user.UserMinInfo

	wg.Add(1)
	go func() {
		nicknames = s.getProfilesAuthNames(profileIds)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		users = s.getProfilesUserInfo(profileIds)
		wg.Done()
	}()

	infos := make(map[string]entity.ProfileMinInfo)
	for _, id := range profileIds {
		fallbackName := id
		separatorIndex := strings.Index(id, "-")
		if separatorIndex >= 0 {
			fallbackName = fmt.Sprintf("#%s", fallbackName[0:separatorIndex])
		}
		infos[id] = entity.ProfileMinInfo{
			Id:          id,
			VisibleName: &fallbackName,
		}
	}

	wg.Wait()

	for id, profile := range users {
		info := infos[id]
		if profile.FullName != nil {
			info.VisibleName = profile.FullName
		}
		info.Avatar = profile.Avatar
		infos[id] = info
	}
	for id, nickname := range nicknames {
		info := infos[id]
		info.VisibleName = &nickname
		infos[id] = info
	}

	return infos, nil
}

func (s *Service) getProfilesUserInfo(profileIds []string) map[string]*user.UserMinInfo {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := s.repos.User.GetUsersMinInfo(ctx, &user.GetUsersMinInfoRequest{UserIds: profileIds})
	cancelCtx()

	if err != nil {
		log.Warnf("unable to get profiles minimal info")
		return make(map[string]*user.UserMinInfo)
	}

	return res.Infos
}

func (s *Service) getProfilesAuthNames(profileIds []string) map[string]string {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := s.repos.Auth.GetVisibleNames(ctx, &auth.GetVisibleNamesRequest{UserIds: profileIds})
	cancelCtx()

	if err != nil {
		log.Warnf("unable to get profiles auth names")
		return make(map[string]string)
	}

	return res.UserVisibleNames
}
