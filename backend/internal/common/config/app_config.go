package config

import (
	"fmt"
	"strings"
)

func App() appConfig {
	ensureConfigLoaded()
	return app
}

var app appConfig

type appConfig struct {
	ENV     string
	Port    int
	Domains []string
	Secret  string
}

func loadAppConfig() {
	app = appConfig{
		ENV:     getENV("APP_ENV"),
		Port:    getIntENV("APP_PORT"),
		Domains: strings.Split(getENV("APP_DOMAINS"), ";"),
		Secret:  getENV("APP_SECRET"),
	}

	if !(app.ENV == "dev" || app.ENV == "staging" || app.ENV == "prod") {
		panic(fmt.Sprintf("Expected env with key 'APP_ENV' to be 'dev' or 'staging' or 'prod', found '%v'", app.ENV))
	}
}
