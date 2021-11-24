package validate

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/pkg/ctx"
)

func Validate(c *gin.Context,form interface{})error{
	err := c.ShouldBind(form)
	if err != nil {
		ctx.Response(c,http.StatusOK,err.Error(),nil)
		return err
	}
	return nil
}