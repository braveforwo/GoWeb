package impl

import (
	"GoWeb/connector"
	"GoWeb/domain"
)

type UploadInfomationServiceImpl struct {
}

func (uisi UploadInfomationServiceImpl) UploadAvatar(user *domain.User, url string) error {
	db := connector.GetDBConn()
	var userinfo domain.Userinfo = domain.Userinfo{}
	db.Where("userid = ?", user.Id).Find(&userinfo)
	tx := db.Begin()
	if userinfo.Id > 0 {
		if err := tx.Model(&userinfo).Update("avatar", url).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
	userinfo.UserId = user.Id
	userinfo.Avatar = url
	if err := tx.Create(&userinfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (uisi UploadInfomationServiceImpl) UpdateInformation(userinfo domain.Userinfo) error {
	db := connector.GetDBConn()
	var userinfo1 domain.Userinfo = domain.Userinfo{}
	db.Where("userid = ?", userinfo.UserId).Find(&userinfo1)
	tx := db.Begin()
	if userinfo1.Id > 0 {
		if err := tx.Model(&userinfo1).Updates(userinfo).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
	if err := tx.Create(&userinfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
