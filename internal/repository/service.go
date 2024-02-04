package repository

import (
	"github.com/mephistolie/chefbook-backend-profile/internal/config"
)

type Repositories struct {
	Auth         *Auth
	User         *User
	Subscription *Subscription
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
	subscriptionService, err := NewSubscription(*cfg.SubscriptionService.Addr)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Auth:         authService,
		User:         userService,
		Subscription: subscriptionService,
	}, nil
}

func (s *Repositories) Stop() error {
	_ = s.Auth.Conn.Close()
	_ = s.User.Conn.Close()
	_ = s.Subscription.Conn.Close()
	return nil
}
