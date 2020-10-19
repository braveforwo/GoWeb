package impl

import (
	"GoWeb/domain"
	"GoWeb/mysqlConnector"
	"errors"
)

type ResgisterServiceImpl struct {
}

func (rsi ResgisterServiceImpl) Register(user *domain.User) error {
	db := mysqlConnector.GetDBConn()
	var searchUser domain.User
	db.Where("phonenumber = ?", user.Phonenumber).First(&searchUser)
	if searchUser.Id > 0 {
		return errors.New("手机号已被注册")
	}
	tx := db.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
