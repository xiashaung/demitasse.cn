package lib

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"log"
)

type db struct {
	Host string
	Port int
	User string
	Password string
	Name string
}

type server struct {
	Host string
	Port int
	Model string
	TimeFormat string
}

type redis struct {
	Host string
	Port int
	password string
}

var (
	cfg *ini.File
	Server server
	Db db
	Redis redis
)

func init(){
	var err error
	var config_file = GetCurrentPath()+"/etc/config.ini"
	cfg,err = ini.Load(config_file)
	if err != nil {
		logrus.Info(err)
	}
	loadServer()
	loadDb()
	loadRedis()
}

func loadServer(){
	sec := GetSection("server")
	Server.Host = sec.Key("host").MustString("0.0.0.0")
	Server.Port = sec.Key("port").MustInt(8901)
	Server.Model = sec.Key("model").MustString("debug")
	Server.TimeFormat = sec.Key("time_format").MustString("2006-01-02 15:04:05")
}

func loadDb(){
	sec := GetSection("database")
	Db.Host = sec.Key("host").MustString("")
	Db.User = sec.Key("user").MustString("")
	Db.Password = sec.Key("password").MustString("")
	Db.Name = sec.Key("name").MustString("")
	Db.Port = sec.Key("port").MustInt(3306)
}

func loadRedis(){
	sec := GetSection("redis")
	Redis.Host = sec.Key("host").MustString("127.0.0.1")
	Redis.Port = sec.Key("port").MustInt(6379)
	Redis.password = sec.Key("password").MustString("")
}

func loadApp(){
	_, err := cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
}

func GetConfig(section string,key string, def string)string{
	sec := GetSection(section)
	return sec.Key(key).MustString(def)
}

func GetSection(name string) *ini.Section {
	r,_ := cfg.GetSection(name)
	return r
}
