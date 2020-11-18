package server

import "GoWeb/domain"

type UploadInfomationService interface {
	UploadAvatar(user *domain.User, url string) error
	UpdateInformation(user domain.User) error
}
