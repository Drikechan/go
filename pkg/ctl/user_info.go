package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FormContext(ctx)

	if !ok {
		return nil, errors.New("查询用户失败")
	}
	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FormContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}
