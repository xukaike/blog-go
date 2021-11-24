package ctx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context,status uint,message string,data interface{}){
	res := make(map[string]interface{})
	res["status"] = status
	res["message"] = message
	res["data"] = data
	c.JSON(http.StatusOK,res)
}