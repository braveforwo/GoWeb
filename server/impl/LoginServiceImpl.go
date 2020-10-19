package impl

import (
	"GoWeb/domain"
	"GoWeb/mysqlConnector"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type LoginServiceImpl struct {
}

func (lsi LoginServiceImpl) Login(user *domain.User) error {
	db := mysqlConnector.GetDBConn()
	var searchUser domain.User
	db.Where("phonenumber = ?", user.Phonenumber).First(&searchUser)
	if searchUser.Id > 0 {
		h := md5.New()
		h.Write([]byte(user.Password))
		cipherStr := h.Sum(nil)
		if searchUser.Password == hex.EncodeToString(cipherStr) {
			return nil
		}
		return errors.New("密码错误！")
	}
	return errors.New("你的手机号尚未注册！")
}
