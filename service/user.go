package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"test-todolist/pkg/ctl"
	"test-todolist/pkg/utils"
	"test-todolist/respository/db/dao"
	"test-todolist/respository/model"
	"test-todolist/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		u = &model.UserList{
			UserName: req.UserName,
			Password: req.Password,
		}
		if err = u.SetPassword(req.Password); err != nil {
			fmt.Println(err, "SetPassword")
			return
		}

		if err = userDao.CreateUser(u); err != nil {
			fmt.Println(err, "CreateUser")
			return
		}
		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}

func (s *UserSrv) Login(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByName(req.UserName)
	if err == gorm.ErrRecordNotFound {
		err = errors.New("用户不存在")
		return
	}

	if !user.CheckPassword(req.Password) {
		err = errors.New("密码错误")
		return
	}

	token, err := utils.GenerateUserToken(user.ID, req.UserName, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	u := &types.UserResp{
		ID:       user.ID,
		Username: req.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}

	uResp := &types.TokenData{
		User:  u,
		Token: token,
	}

	return ctl.RespSuccessWithData(uResp), nil
}
