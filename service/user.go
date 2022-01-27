package service

import (
	"github.com/jinzhu/gorm"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    e.GetMsg(e.InvalidParams),
		}
	}
	user.UserName = service.UserName
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    err.Error(),
		}
	}
	err := model.DB.Create(&user).Error
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    "用户创建成功",
	}
}
func (service *UserService) Login() serializer.Response {

	var user model.User
	err := model.DB.Where("user_name=?", service.UserName).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: e.InvalidParams,
				Msg:    "用户不存在，请先登录",
			}
		}
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    "密码错误",
		}
	}
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登陆成功",
	}
}
