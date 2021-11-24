package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;notNull;comment:用户名"`
	Email string `json:"email" gorm:"comment:邮箱"`
	Password string `json:"password" gorm:"comment:密码"`
	Desc string `json:"desc" gorm:"简介"`
	Avatar string `json:"avatar" gorm:"头像"`
	BlogsNum int `json:"blogs_num" gorm:"微博数"`
	FollowingNum int `json:"follow_num" gorm:"关注数"`
	FollowerNum int `json:"follower_num" gorm:"粉丝数"`
}

func GetUserList()([]*User,error){
	var users []*User
	err := db.Find(&users).Error
	if err != nil {
		return nil,err
	}
	return users,err
}

func GetUserByName(name string)(*User,error){
	var user *User
	err := db.Where("name = ?",name).First(&user).Error
	if err != nil {
		return nil,err
	}
	return user,nil
}

func AddUser(data map[string]interface{})error{
	user := User{
		Name: data["name"].(string),
		Email :data["email"].(string),
		Password :data["password"].(string),
		Desc :data["desc"].(string),
		Avatar :data["avatar"].(string),
	}
	err := db.Create(&user).Error
	if err != nil{
		return  err
	}
	return nil
}