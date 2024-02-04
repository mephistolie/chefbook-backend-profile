package config

import (
	"github.com/mephistolie/chefbook-backend-common/log"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	AuthService         Service
	UserService         Service
	SubscriptionService Service
}

type Service struct {
	Addr *string
}

func (c Config) Validate() error {
	if *c.Environment != EnvProd {
		*c.Environment = EnvDev
	}
	return nil
}

func (c Config) Print() {
	log.Infof("PROFILE SERVICE CONFIGURATION\n"+
		"Environment: %v\n"+
		"Port: %v\n"+
		"Logs path: %v\n\n"+
		"Auth Service Address: %v\n"+
		"User Service Address: %v\n"+
		"Subscription Service Address: %v\n\n",
		*c.Environment, *c.Port, *c.LogsPath,
		*c.AuthService.Addr, *c.UserService.Addr, *c.SubscriptionService.Addr,
	)
}
