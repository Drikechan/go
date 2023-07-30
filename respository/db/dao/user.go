package dao

import (
	"context"
	"gorm.io/gorm"
	"test-todolist/respository/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		// 创建一个空的context
		ctx = context.Background()
	}
	return &UserDao{NewClient(ctx)}
}

func (dao *UserDao) FindUserByName(userName string) (user *model.UserList, err error) {
	err = dao.DB.Model(&model.UserList{}).Where("user_name=?", userName).First(&user).Error
	return
}

func (dao *UserDao) CreateUser(user *model.UserList) (err error) {
	err = dao.DB.Model(&model.UserList{}).Create(user).Error
	return
}

func (dao *UserDao) FindUserByUserId(Id uint) (user *model.UserList, err error) {
	err = dao.DB.Model(&model.UserList{}).Where("id=?", Id).First(&user).Error
	return
}
