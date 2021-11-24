package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"blog/conf"
)

var db *gorm.DB

func Init(){
	var err error
	db,err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseConf.User,
		conf.DatabaseConf.Password,
		conf.DatabaseConf.Host,
		conf.DatabaseConf.Name)),&gorm.Config{})
	if err != nil {
		log.Fatalf("connect mysql fail, error: %v \n",err)
	}
	db.AutoMigrate(&User{})
}