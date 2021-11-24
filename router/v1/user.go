package v1

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"blog/pkg/ctx"
	"blog/pkg/util"
	"blog/pkg/validate"
	"blog/service"
)

func Login(c *gin.Context){
	var (
		loginForm validate.Login
		err       error
	)
	if err = validate.Validate(c,&loginForm);err != nil{
		return
	}
	userService := service.User{
		Name :     loginForm.Name,
		Password : loginForm.Password,
	}
	user,err := userService.GetByName()
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			ctx.Response(c,http.StatusBadRequest,ctx.UserNotExist,nil)
			return
		}
		ctx.Response(c,http.StatusInternalServerError,ctx.ServerError,nil)
		return
	}
	if user.Password != util.Md5(userService.Password){
		ctx.Response(c,http.StatusBadRequest,ctx.PasswordWrong,nil)
		return
	}
	session := sessions.Default(c)
	session.Set("userInfo",user)
	session.Save()
	ctx.Response(c,http.StatusOK,ctx.OK,nil)
	return
}

func Register(c *gin.Context){
	var (
		registerForm validate.Register
		err error
	)
	if err = validate.Validate(c,&registerForm);err != nil{
		return
	}
	userService := service.User{
		Name :registerForm.Name,
		Email :registerForm.Email,
		Password :registerForm.Password,
		Desc :registerForm.Desc,
		Avatar :registerForm.Avatar,
	}
	user,err := userService.GetByName()
	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound){
		ctx.Response(c,http.StatusInternalServerError,ctx.ServerError,nil)
		return
	}
	if user != nil{
		ctx.Response(c,http.StatusBadRequest,ctx.UserExist,nil)
		return
	}
	err = userService.Create()
	if err != nil {
		ctx.Response(c,http.StatusInternalServerError,ctx.ServerError,nil)
		return
	}
	ctx.Response(c,http.StatusOK,ctx.OK,nil)
}