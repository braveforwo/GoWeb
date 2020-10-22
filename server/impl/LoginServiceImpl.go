package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type LoginServiceImpl struct {
}

func (lsi LoginServiceImpl) Login(user *domain.User) error {
	db := connector.GetDBConn()
	var searchUser domain.User
	db.Where("phonenumber = ?", user.Phonenumber).First(&searchUser)
	if searchUser.Id > 0 {
		h := md5.New()
		h.Write([]byte(user.Password))
		cipherStr := h.Sum(nil)
		if searchUser.Password == hex.EncodeToString(cipherStr) {
			user.Id = searchUser.Id
			user.UserName = searchUser.UserName
			user.Phonenumber = ""
			user.Password = ""
			return nil
		}
		return errors.New("密码错误！")
	}
	return errors.New("你的手机号尚未注册！")
}
