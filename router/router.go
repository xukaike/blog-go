package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
	"blog/conf"
	v1 "blog/router/v1"
)

func InitRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	store,err := redis.NewStore(10,"tcp",conf.RedisConf.Host,conf.RedisConf.Password,[]byte("secret"))
	if err != nil {
		log.Fatalf("init redis fail,error :%v",err)
	}
	r.Use(sessions.Sessions("token", store))
	user := r.Group("/user")
	{
		user.POST("/register",v1.Register)
		user.POST("/login",v1.Login)
	}

	return r
}