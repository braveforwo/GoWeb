package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"errors"
)

type ResgisterServiceImpl struct {
}

func (rsi ResgisterServiceImpl) Register(user *domain.User) error {
	db := connector.GetDBConn()

	var searchUser domain.User
	db.Where("phonenumber = ?", user.Phonenumber).Find(&searchUser)
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
