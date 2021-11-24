package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File

type database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseConf = &database{}

type redis struct {
	Host string
	Password string
}

var RedisConf = &redis{}

func init(){
	var err error
	cfg,err = ini.Load("./app.ini")
	if err != nil {
		log.Fatalf("conf setup fail,error :%v",err)
	}

	mapTo("database",DatabaseConf)
	mapTo("redis",RedisConf)
}

func mapTo(name string,v interface{}){
	err := cfg.Section(name).MapTo(v)
	if err != nil {
		log.Fatalf("fail to map %s,err :%v",name,err)
	}
}