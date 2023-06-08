package main

import (
	"flag"
	"github.com/mephistolie/chefbook-backend-profile/internal/app"
	"github.com/mephistolie/chefbook-backend-profile/internal/config"
	"github.com/peterbourgon/ff/v3"
	"os"
)

func main() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := config.Config{
		Environment: fs.String("environment", "debug", "service environment"),
		Port:        fs.Int("port", 8080, "service port"),
		LogsPath:    fs.String("logs-path", "", "logs file path"),

		AuthService: config.AuthService{
			Addr: fs.String("auth-addr", "", "auth service address"),
		},

		UserService: config.UserService{
			Addr: fs.String("user-addr", "", "user service address"),
		},
	}
	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars()); err != nil {
		panic(err)
	}

	err := cfg.Validate()
	if err != nil {
		panic(err)
	}

	app.Run(&cfg)
}
