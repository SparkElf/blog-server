package config

import (
	"gopkg.in/ini.v1"
)

var (
	//user
	Username string
	//github
	APIToken string
	//server
	AppMode  string
	HttpPort string
	//jwt
	JwtKey     string
	ExpireTime int64
	//db
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func Init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		panic(err)
	}
	//user
	Username = cfg.Section("user").Key("Username").String()
	//github
	APIToken = cfg.Section("github").Key("APIToken").String()
	//server
	AppMode = cfg.Section("server").Key("AppMode").String()
	HttpPort = cfg.Section("server").Key("HttpPort").String()
	//jwt
	JwtKey = cfg.Section("jwt").Key("JwtKey").String()
	ExpireTime = parseTimeStr(cfg.Section("jwt").Key("ExpireTime").String())
	//db
	Db = cfg.Section("database").Key("Db").String()
	DbHost = cfg.Section("database").Key("DbHost").String()
	DbPort = cfg.Section("database").Key("DbPort").String()
	DbUser = cfg.Section("database").Key("DbUser").String()
	DbPassword = cfg.Section("database").Key("DbPassword").String()
	DbName = cfg.Section("database").Key("DbName").String()
}
