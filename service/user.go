package service

import (
	"blog/model"
	"blog/pkg/util"
)

type User struct {
	ID int
	Name string
	Email string
	Password string
	Desc string
	Avatar string
	BlogsNum int
	FollowingNum int
	FollowerNum int
}

func GetList()([]*model.User,error){
	var (
		users []*model.User
	)
	users,err := model.GetUserList()
	if err != nil {
		return nil,err
	}
	return users,nil
}

func (u *User)GetByName()(*model.User,error){
	user,err := model.GetUserByName(u.Name)
	if err != nil {
		return nil,err
	}
	return user,nil
}

func (u *User) Create()error{
	user := map[string]interface{}{
		"name":u.Name,
		"email":u.Email,
		"password":util.Md5(u.Password),
		"desc":u.Desc,
		"avatar":u.Avatar,
	}
	err := model.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}