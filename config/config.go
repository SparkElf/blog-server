package config

import (
	"gopkg.in/ini.v1"
)

var (
	Username string
	APIToken string
)

func Init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		panic(err)
	}
	Username = cfg.Section("user").Key("username").String()
	APIToken = cfg.Section("github").Key("apiToken").String()
}
