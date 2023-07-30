package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserList struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Password       string
	PasswordDigest string
}

const PasswordCost = 12

func (user *UserList) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *UserList) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
