package server

import "GoWeb/domain"

type GetInfomationService interface {
	GetInfomationByUserId(userid int) domain.Userinfo
}
