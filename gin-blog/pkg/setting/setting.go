package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var Cfg *ini.File
var RunMode string
var HTTPPort int
var ReadTimeout time.Duration
var WriteTimeout time.Duration
var PageSize int
var JwtSecret string

func init() {
	configFile := "conf/app.ini"
	var err error
	Cfg, err = ini.Load(configFile)
	if err != nil {
		log.Fatalf("Fail to parse %s: %v", configFile, err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
