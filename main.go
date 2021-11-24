package main

//import "github.com/gin-gonic/gin"
import (
	"blog/model"
	"blog/router"
)

func init(){
	model.Init()
}

func main() {
	r := router.InitRouter()
	r.Run()
}