package repository

import (
	"github.com/mephistolie/chefbook-backend-profile/internal/config"
)

type Repositories struct {
	Auth *Auth
	User *User
}

func NewRepositories(cfg *config.Config) (*Repositories, error) {
	authService, err := NewAuth(*cfg.AuthService.Addr)
	if err != nil {
		return nil, err
	}
	userService, err := NewUser(*cfg.UserService.Addr)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Auth: authService,
		User: userService,
	}, nil
}

func (s *Repositories) Stop() error {
	_ = s.Auth.Conn.Close()
	_ = s.User.Conn.Close()
	return nil
}
