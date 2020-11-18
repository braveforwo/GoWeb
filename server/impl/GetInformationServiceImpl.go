package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
)

type GetInfomationServiceImpl struct {
}

func (uisi GetInfomationServiceImpl) GetInfomationByUserId(userid int) (domain.Userinfo, error) {
	var userinfo domain.Userinfo = domain.Userinfo{}
	db := connector.GetDBConn()
	if err := db.Where("userid = ?", userid).Find(&userinfo).Error; err != nil {
		return userinfo, err
	}
	return userinfo, nil
}
